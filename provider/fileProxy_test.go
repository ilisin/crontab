package provider

import (
	"testing"
	"time"
	"os"
	"github.com/ilisin/crontab/job"
	"io"
	"github.com/astaxie/beego"
)

func TestWriteJobs(t *testing.T){
	bjs := []job.BasicJob{
		job.BasicJob{
			Id:"TEST_JOB_001",
			Name:"测试任务1",
			ExeTime:time.Now().Add(2 * time.Hour),
			Type:job.JOB_ADDIN,
			Live:1,
			Interval:0,
			Command: "yum install update",
			Params:""},
		job.BasicJob{
			Id:"TEST_JOB_002",
			Name:"测试任务2",
			ExeTime:time.Now().Add(-36 * time.Hour),
			Type:job.JOB_CALLBACK,
			Live:5,
			Interval: 24*time.Hour,
			Command: "http://192.168.10.49/callback",
			Params:""},
		job.BasicJob{
			Id:"TEST_JOB_003",
			Name:"测试任务3",
			ExeTime:time.Now().Add(4 * time.Hour),
			Type:job.JOB_SYSTEM,
			Live:1,
			Interval:0,
			Command: "JOB_SYS_0001",
			Params:""},
		job.BasicJob{
			Id:"TEST_JOB_004",
			Name:"测试任务4",
			ExeTime:time.Now().Add(72 * time.Hour),
			Type:job.JOB_SYSTEM,
			Live:1,
			Interval:0,
			Command: "JOB_SYS_0001",
			Params:""},
	}
	var proxy StorageProxyer
	proxy = &FileProxy{
		date : time.Now(),
	}
	err := proxy.WriteJobs(bjs)
	if err != nil {
		t.Fatal(err)
	}
	aBj := job.BasicJob{
		Id:"TEST_JOB_231",
		Name:"测试任务5",
		ExeTime:time.Now().Add(2 * time.Hour),
		Type:job.JOB_ADDIN,
		Live:1,
		Interval:0,
		Command: "yum install update",
		Params:""}
	err = proxy.WriteAJob(aBj)
	if err != nil {
		t.Fatal(err)
	}
	beego.Trace("写入job完成")
}

func copyFile(src, des string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		return err
	}
	defer desFile.Close()

	_,err = io.Copy(desFile, srcFile)
	return err
}

func TestRemove(t *testing.T){
	//copyFile("a.txt","b.txt")
	fileProxy := &FileProxy{
		date : time.Now(),
	}
	err := fileProxy.DeleteAJobById("TEST_JOB_002",time.Now().Add(-36 * time.Hour),true)
	if err != nil {
		t.Fatal(err)
	}
	aBj := job.BasicJob{
		Id:"TEST_JOB_001",
		ExeTime:time.Now().Add(2 * time.Hour),
		Type:job.JOB_ADDIN,
		Live:1,
		Interval:0,
		Command: "yum install update",
		Params:""}
	err = fileProxy.DeleteAJob(aBj)
	if err != nil {
		t.Fatal(err)
	}
	beego.Trace("删除job完成")
}
