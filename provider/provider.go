package provider

import (
	"github.com/ilisin/crontab/job"
	"path"
	"os"
	"time"
	"github.com/ilisin/crontab/executer"
	"github.com/astaxie/beego"
)

type Provider struct {
	proxy FileProxy
	OutStream chan job.BasicJob
	date time.Time
}



func (p *Provider)SetProxy(fp FileProxy){
	p.proxy = fp
}

func (p Provider)PushAJob(j job.BasicJob){
	etsd := j.ExeTime.Format("20060102")
	tsd := p.date.Format("20060102")
	if etsd == tsd {
		p.OutStream <- j
	}
	p.proxy.WriteAJob(j)
}

func (p Provider)UpdateAJob(j job.BasicJob) error{
	etsd := j.ExeTime.Format("20060102")
	tsd := p.date.Format("20060102")
	if etsd == tsd {
		return p.proxy.UpdateAJob(j)
	}
	return nil
}

// update job
func (p Provider)getUpdaterJob() error{
	dt := time.Now()
	p.date = dt
	p.proxy.date = dt
	jbs,err := p.proxy.ReadJobs()
	if err != nil {
		beego.Error(err)
		return err
	}
	for _,j := range jbs{
		p.OutStream <- j
	}
	return nil
}

func (p Provider)Start() error{
	executer.RegisterSystemJob("JOB_SYS_1001",p.getUpdaterJob)

	//1分钟后， 每24小时更新一次
	job := job.BasicJob{
		Id:"000000100001",
		Name:"定期查找任务",
		ExeTime:time.Now().Add(1 * time.Minute),
		Type:job.JOB_SYSTEM,
		Live:-1,
		Interval:24 * time.Hour,
		Command: "JOB_SYS_1001",
		Params:""}
//	job := job.BasicJob{
//		Id:"000000100001",
//		Name:"定期查找任务",
//		ExeTime:time.Now().Add(2 * time.Second),
//		Type:job.JOB_SYSTEM,
//		Live:-1,
//		Interval:5 * time.Second,
//		Command: "JOB_SYS_1001",
//		Params:""}
	p.OutStream <- job
	return nil
}


func init(){
	//check directory
	curPath,_ := os.Getwd()
	curPath = path.Join(curPath,"data")
	os.Mkdir(curPath,os.ModePerm)
}
