package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 将rawData转换成map[string]string
func RawData2Map(rawData []byte) (res map[string]string) {
	res = make(map[string]string)
	str := string(rawData)
	if IsBlank(str) || !strings.Contains(str, "&") {
		return res
	}

	keyValueArr := strings.Split(str, "&")
	for _, kv := range keyValueArr {
		if IsBlank(kv) || !strings.Contains(kv, "=") {
			continue
		}

		kvProp := strings.Split(kv, "=")
		res[kvProp[0]] = kvProp[1]
	}
	return res
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {
	log.Println("url", url)
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送GET请求并传参
func GetWithParam(URL string, param url.Values) string {
	u, err := url.ParseRequestURI(URL)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = param.Encode() // URL encode
	fmt.Println(u.String())

	resp, err := http.Get(u.String()) // 发送请求

	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
	}

	return string(b)
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) (string, error) {
	log.Printf("{url:%s,param:%s,contentType:%s}", url, data, contentType)
	// 超时时间：5秒
	client := &http.Client{Timeout: 1 * time.Second}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		log.Printf("post request json marshal error %s", err.Error())
		return "", err
	}
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Printf("post request post error %s", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("post request read reponse error %s", err.Error())
		return "", err
	}
	res := string(result)
	log.Printf("res{%s}", res)
	return res, nil
}
