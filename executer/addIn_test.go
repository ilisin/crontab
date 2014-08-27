package executer

import (
	"testing"
	"github.com/astaxie/beego"
)

func TestAddIn(t *testing.T) {
	addIn := &AddIn{}
	rc,err := addIn.Exec("/usr/bin/sh")
	if err != nil {
		t.Error("can't execute the command")
	}
	<-rc
	beego.Trace("excute ok")
}
