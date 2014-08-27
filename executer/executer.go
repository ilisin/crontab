package executer

import (
	"github.com/astaxie/beego"
)

type Executer interface {
	Exec(cmd string,params ...string) (chan bool,error)
}


func init(){
	// register the system job
	RegisterSystemJob("JOB_SYS_0001",Demo)
}


func Demo() error{
	beego.Trace("DEMO注册函数")
	return nil
}
