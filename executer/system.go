package executer

import "errors"

var sysJobRegistery map[string]func() error

type System struct {
}

func (s System)Exec(cmd string,params ...string) (chan bool,error) {
	retChan := make(chan bool)
	fn,ok := sysJobRegistery[cmd]
	if !ok {
		return retChan,errors.New("函数不存在")
	}
	go func() {
		fn()
		retChan <- true
	}()
	return retChan,nil
}

func RegisterSystemJob(sysId string,fn func() error) error{
	if sysJobRegistery == nil {
		sysJobRegistery = make(map[string]func()error)
	}

	if _,ok := sysJobRegistery[sysId];ok {
		return errors.New("重复注册")
	}
	sysJobRegistery[sysId] = fn
	return nil
}
