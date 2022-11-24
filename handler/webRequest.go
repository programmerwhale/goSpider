package handler

import (
	"fmt"
	"log"
	"net/http"
)

func WebRequest(requestType string,weburl string,paramStr string ) *http.Response {
	Config:=ViperHelper()
	client := &http.Client{}                                                               //初始化客户端
	req, err := http.NewRequest(requestType, weburl+paramStr, nil) //建立连接
	if err != nil {
		log.Fatal(err)
	}

	headers:=Config.Get("douban")
	for k,v:=range headers.(map[string]interface{}){
		req.Header.Set(k,fmt.Sprintf("%v",v))
	}

	resp, err := client.Do(req) //拿到返回的内容
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
