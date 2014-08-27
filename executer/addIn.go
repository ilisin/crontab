package executer

import "os/exec"

type AddIn struct {
}

func (c AddIn)Exec(cmd string,params ...string) (chan bool,error) {
	retChan := make(chan bool)
	command := exec.Command(cmd,params...)
	go func(){
		command.Run()
		retChan <- true
	}()
	return retChan,nil
}
