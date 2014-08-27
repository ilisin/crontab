package main

import (
	"github.com/ilisin/crontab"
	"time"
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
	crontab.Run()

	time.Sleep(1 * time.Hour)
}

