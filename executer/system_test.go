package executer

import (
	"testing"
	"imooly.com/utility"
)

func TestSystem(t *testing.T){
	system := &System{}
	system.Exec("JOB_SYS_0001")
	utility.Trace("测试完成")
}

