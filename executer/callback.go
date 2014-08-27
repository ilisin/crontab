package executer

import (
	"time"
	"imooly.com/crontab/reference/httpclient"
	"net/http"
	"regexp"
	"errors"
)

var (
	url_regex = `(http|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`
)

type Callback struct {
}

//params 参数无效，http get 参数直接放入cmd
func (c Callback)Exec(cmd string,params ...string) (chan bool,error) {
	retChan := make(chan bool)
	reg := regexp.MustCompile(url_regex)
	match := reg.MatchString(cmd)
	if !match {
		return retChan,errors.New("http(https)回调url不匹配")
	}
	transport := &httpclient.Transport{
		ConnectTimeout:			2*time.Second,
		RequestTimeout:			10*time.Second,
		ResponseHeaderTimeout:	5*time.Second,
	}
	defer transport.Close()
	go func() {
		client := &http.Client{
			Transport : transport}
		req, _ := http.NewRequest("GET",cmd,nil)
		//ignore
		resp, err := client.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		retChan <- true
	}()
	return retChan,nil
}
