package executer

import (
	"testing"
	"github.com/astaxie/beego"
)

func TestCallback(t *testing.T) {
	callback := &Callback{}
	rc,err := callback.Exec("http://192.168.10.49:8080/v1/user")
	if err != nil {
		t.Error("execute command error")
	}
	<-rc
	beego.Trace("callback success")
}
