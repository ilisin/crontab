package crontab

import (
	"imooly.com/crontab/provider"
	"imooly.com/crontab/job"
)

var s_engine Engine

type Engine struct {
	Provider provider.Provider
	InStream chan job.BasicJob
}

func (e Engine)Execute(){
	for{
		job := <- e.InStream
		job.StartDo()
	}
}

func (e *Engine)Start(){
	e.InStream = e.Provider.OutStream
	e.Provider.Start()
	go e.Execute()
}

func (e *Engine)PushAJob(j job.BasicJob){
	e.Provider.PushAJob(j)
}
