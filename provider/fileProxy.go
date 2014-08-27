package provider

import (
	"os"
	"sync"
	"time"
	"path"
	"io/ioutil"
	"encoding/json"
	"imooly.com/crontab/job"
)

type FileProxy struct {
	lock sync.RWMutex
	date time.Time
}

//** private functions
func (fp FileProxy)getSingleJobsFile(t time.Time) string{
	curPath,_ := os.Getwd()
	curPath = path.Join(curPath,"data")
	return path.Join(curPath,t.Format("20060102.json"))
}

func (fp FileProxy)getRepeatJobsFile() string{
	return "repeat.json"
}

func (fp FileProxy)readJobs(filename string) (jbs []job.BasicJob,err error){
	fp.lock.RLock()
	defer fp.lock.RUnlock()
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &jbs)
	if err != nil {
		return nil,err
	}
	return jbs, nil
}

func (fp FileProxy)writeJobs(filename string,jbs []job.BasicJob) error {
	fp.lock.Lock()
	defer fp.lock.Unlock()
	content,err := json.Marshal(jbs)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename,content,os.ModePerm)
}

func (fp FileProxy)convertDate(j job.BasicJob) (rt time.Time,err error){
	etsd := j.ExeTime.Format("20060102")
	tsd := fp.date.Format("20060102")
	if tsd == etsd {
		return fp.date,nil
	}
	if fp.date.Before(j.ExeTime) {
		return j.ExeTime,nil
	}
	tt := j.ExeTime
	for {
		if fp.date.Before(tt) {
			return tt,nil
		}
		tt = tt.Add(24*time.Hour)
	}
}

func (fp FileProxy)jobFile(j *job.BasicJob) string{
	curPath,_ := os.Getwd()
	curPath = path.Join(curPath,"data")
	if j.IsRepeat() {
		t,_ := fp.convertDate(*j)
		j.ExeTime = t
		return path.Join(curPath,t.Format("repeat.json"))
	} else {
		return path.Join(curPath,j.ExeTime.Format("20060102.json"))
	}
}

func (fp FileProxy)ReadJobs() (jbs []job.BasicJob,err error){
	sjs,err := fp.readJobs(fp.getSingleJobsFile(fp.date))
	if err != nil {
		return nil,err
	}
	rjs,err := fp.readJobs(fp.getRepeatJobsFile())
	if err != nil {
		return nil,err
	}
	jbs = sjs[:]
	jbs = append(jbs,rjs...)
	return jbs,nil
}

func (fp FileProxy)WriteJobs(jbs []job.BasicJob) error {
	jobMap := make(map[string][]job.BasicJob)
	for _,j := range jbs {
		sfn := fp.jobFile(&j)
		if _,ok := jobMap[sfn];ok {
			jobMap[sfn] = append(jobMap[sfn],j)
		} else {
			tsl := make([]job.BasicJob,0)
			tsl = append(tsl,j)
			jobMap[sfn] = tsl
		}
	}

	for key,value := range jobMap{
		err := fp.writeJobs(key,value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fp FileProxy)WriteAJob(j job.BasicJob) error{
	filename := fp.jobFile(&j)
	fp.lock.Lock()
	fp.lock.Unlock()
	jbs,err := fp.readJobs(filename)
	if err != nil {
		return err
	}
	jbs = append(jbs,j)
	return fp.writeJobs(filename,jbs)
}

func (fp FileProxy)UpdateAJob(j job.BasicJob) error{
	filename := fp.jobFile(&j)
	jbs,err := fp.readJobs(filename)
	if err != nil {
		return err
	}
	index := 0
	for i,jb := range jbs {
		if j.Compare(jb) {
			index = i
			break
		}
	}
	jt := jbs[0:index]
	jt = append(jt,j)
	//jt = append(jt,jbs,index+1)
	jt = append(jt,jbs[index+1:]...)
	return fp.writeJobs(filename,jt)
}

func (fp FileProxy)DeleteAJob(j job.BasicJob) error{
	filename := fp.jobFile(&j)
	jbs,err := fp.readJobs(filename)
	if err != nil {
		return err
	}
	//jt := make([]job.Jobber,0,len(jbs))
	index := 0
	for i,jb := range jbs {
		if j.Compare(jb) {
			index = i
			break
		}
	}
	jt := jbs[0:index]
	jt = append(jt,jbs[index+1:]...)
	return fp.writeJobs(filename,jt)
}

func (fp FileProxy)DeleteAJobById(jId string,t time.Time,repeat bool) error{
	jb := job.BasicJob{}
	jb.Id = jId
	jb.ExeTime = t
	if repeat {
		jb.Interval = 1
	} else {
		jb.Interval = 0
	}
	return fp.DeleteAJob(jb)
}
