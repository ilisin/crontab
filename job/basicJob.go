package job

import (
	"time"
	"github.com/ilisin/crontab/executer"
	"github.com/ilisin/crontab/reference/coop"
	"errors"
	"github.com/astaxie/beego"
)


const (
	JOB_SYSTEM = 0
	JOB_CALLBACK = 1
	JOB_ADDIN = 2
)

type JOB_TYPE int

type BasicJob struct {
	Id string
	Name string
	ExeTime time.Time
	Type JOB_TYPE
	Live int
	Interval time.Duration
	Command string
	Params string
}

// private functions
func (bj BasicJob)getExecuter() executer.Executer{
	switch bj.Type{
	case JOB_SYSTEM:
		return executer.System{}
	case JOB_CALLBACK:
		return executer.Callback{}
	case JOB_ADDIN:
		return executer.AddIn{}
	}
	return nil
}

func (bj *BasicJob)ReDo(dt time.Duration) error{
	if bj.Live == 0 {
		return nil
	}
	if err := bj.Do();err != nil {
		return err
	}
	bj.Live--
	ch := coop.After(dt,func(){
			bj.ReDo(dt)
		})
	<- ch
	return nil
}

// public functions
func (bj BasicJob)Do() error {
	etr := bj.getExecuter()
	if etr == nil {
		panic("执行器错误")
	}
	beego.Debug("执行任务 : ID:",bj.Id," Name:",bj.Name)
	_,err := etr.Exec(bj.Command,bj.Params)
	if err != nil {
		return err
	}
	return nil
}

func (bj BasicJob)DoAt(t time.Time) error{
	if t.Before(time.Now()){
		return errors.New("任务已经过期")
	}
	ch := coop.At(t,func(){
			bj.Do()
		})
	<- ch
	return nil
}

func (bj BasicJob)ReDoAt(t time.Time,dt time.Duration) error {
	tt := t
	for {
		if tt.After(time.Now()){
			break;
		}
		tt = tt.Add(dt)
	}
	ch := coop.At(t,func(){
			bj.ReDo(dt)
		})
	<- ch
	return nil
}

func (bj BasicJob)StartDo() error{
	var err error
	if bj.Interval <= 0 {
		err = bj.DoAt(bj.ExeTime)
	}else{
		err = bj.ReDoAt(bj.ExeTime,bj.Interval)
	}
	return err
}

func (bj BasicJob)Compare(j Jobber) bool{
	return bj.Id == j.(BasicJob).Id
}

func (bj BasicJob)IsRepeat() bool{
	return bj.Interval > 0
}
