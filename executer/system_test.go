package executer

import (
	"testing"
	"github.com/astaxie/beego"
)

func TestSystem(t *testing.T){
	system := &System{}
	system.Exec("JOB_SYS_0001")
	beego.Trace("测试完成")
}

