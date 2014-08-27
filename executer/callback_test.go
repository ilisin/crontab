package executer

import (
	"testing"
	"imooly.com/utility"
)

func TestCallback(t *testing.T) {
	callback := &Callback{}
	rc,err := callback.Exec("http://192.168.10.49:8080/v1/user")
	if err != nil {
		t.Error("回调错误")
	}
	<-rc
	utility.Trace("回调完成")
}
