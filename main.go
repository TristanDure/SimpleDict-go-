package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	//"io"
	"io/ioutil"
	"log"
	"net/http"
	//"strings"
)

type Dict struct {
	TransType string `json:"trans_type"`
	Source string `json:"source"`
	UserID string `json:"user_id"`
}

type Dictresponse struct {
	Rc int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En string `json:"en"`
		} `json:"prons"`
		Explanations []string `json:"explanations"`
		Synonym []string `json:"synonym"`
		Antonym []string `json:"antonym"`
		WqxExample [][]string `json:"wqx_example"`
		Entry string `json:"entry"`
		Type string `json:"type"`
		Related []interface{} `json:"related"`
		Source string `json:"source"`
	} `json:"dictionary"`
}
func query(word string) {
	client := &http.Client{}
	//var data = strings.NewReader(`{"trans_type":"en2zh","source":"good"}`)
	request:=Dict{TransType: "en2zh",Source:word}
	buf,err:=json.Marshal(request)
	if err!=nil{
		log.Fatal(err)
	}
	var data=bytes.NewReader(buf)

	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode!=200{
		log.Fatal("Bad StatusCode",resp.StatusCode,"body",string(bodyText))
	}
	//fmt.Printf("%s\n", bodyText)
	var dictresponse Dictresponse
	err=json.Unmarshal(bodyText,&dictresponse)
	if err!=nil{
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", dictresponse)
	fmt.Println(word,"UK:",dictresponse.Dictionary.Prons.EnUs,"US:",dictresponse.Dictionary.Prons.EnUs)
	for _, item:=range dictresponse.Dictionary.Explanations{
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr,`Usage: simpleDict WORD example: simpleDict hello`)
		os.Exit(1)
	}
	word:=os.Args[1]
	query(word)	
}