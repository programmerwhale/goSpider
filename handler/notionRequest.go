package handler

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func PostHttpsSkip(url string, bytesData []byte) ([]byte, error) {

	Config:=ViperHelper()

	// 创建各类对象
	var client *http.Client
	var request *http.Request
	var resp *http.Response
	var body []byte
	var err error

	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}

	request, err = http.NewRequest("POST", url, bytes.NewReader(bytesData))

	if err != nil {
		log.Println("GetHttpSkip Request Error:", err)
		return nil, nil
	}

	request.Header.Add("Notion-Version", "2022-02-22")
	request.Header.Set("Content-Type", "application/json")

	// 加入 token
	token:=fmt.Sprintf("%v",Config.Get("notion.token"))
	request.Header.Add("Authorization", token)
	resp, err = client.Do(request)
	if err != nil {
		log.Println("GetHttpSkip Response Error:", err)
		return nil, nil
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	defer client.CloseIdleConnections()

	fmt.Println(url)
	return body, err
}