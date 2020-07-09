package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)



func main() {
	/*
	基本方式
	请求http api
	*/
	call := func (host string,path string,method string) {
		// 因为是GET请求方式，所以body is nil
		req,err := http.NewRequest(method,host + path,nil)
		if err != nil {
			panic(err)
		}

		// 发送请求
		cli := http.DefaultClient
		rsp,err := cli.Do(req)
		if err != nil {
			panic(err)
		}

		// 提取结果
		defer rsp.Body.Close()
		buf,err := ioutil.ReadAll(rsp.Body)

		// 打印结果
		fmt.Println(buf)
	}

	call("http://127.0.0.1:8013","/","GET")
}
