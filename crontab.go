package crontab

import (
	"github.com/ilisin/crontab/provider"
	"github.com/ilisin/crontab/job"
)


func Run(){
	s_engine = Engine{}
	s_engine.Provider = provider.Provider{}
	s_engine.Provider.OutStream = make(chan job.BasicJob,128)
	s_engine.Start()
}
