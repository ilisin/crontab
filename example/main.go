package main

import (
	"imooly.com/crontab"
	"time"
	"imooly.com/utility"
)

var (
	logConfig string = `{
	"filename":"logs.log",
	"maxlines":10000,
	"maxsize":10000000,
	"daily":true,
	"maxdays":15,
	"rotate":true
	}`

)

func main(){
	utility.SetLogger("file",logConfig)
	crontab.Run()

	time.Sleep(1 * time.Hour)
}

