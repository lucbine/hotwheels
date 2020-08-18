package httpcall

import (
	"bytes"
	"errors"
	"fmt"
	"hotwheels/agent/internal/util"
	"net"
	"net/http"
	"net/url"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var (
	TransportDefault = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   200 * time.Millisecond,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConnsPerHost:   1000,
		IdleConnTimeout:       30 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ResponseHeaderTimeout: 800 * time.Millisecond,
	}
)

//复用 http.Transport，默认自带连接池
func NewHttpClient(timeout time.Duration) *http.Client {
	return &http.Client{Transport: TransportDefault, Timeout: timeout}
}

type Req struct {
	Url     string
	Params  map[string]interface{}
	Header  map[string]string
	TimeOut time.Duration
}

type Response struct {
	Result     interface{}
	HttpStatus int
}

/*
	GET 请求
	url : 请求地址
	params ： 请求参数
	setHeader ： 设置头部
	result ： 返回结果
*/

func Post(httpReq Req, result interface{}) (err error) {
	//构建请求
	jsonStr, _ := jsoniter.Marshal(httpReq.Params)
	method := "POST"
	req, err := http.NewRequest(method, httpReq.Url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	//如果设置了头部
	if httpReq.Header != nil && len(httpReq.Header) > 0 {
		for key, val := range httpReq.Header {
			req.Header.Set(key, val)
		}
	}
	//默认一秒超时
	if httpReq.TimeOut.Milliseconds() <= 0 {
		httpReq.TimeOut = time.Second
	}
	//生成http.client
	client := NewHttpClient(httpReq.TimeOut)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	//状态码
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("http status code is %d", resp.StatusCode))
	}
	//关闭响应
	defer resp.Body.Close()
	err = jsoniter.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

/*
   GET 请求
   url : 请求地址
   params ： 请求参数
   setHeader ： 设置头部
   result ： 返回结果
*/
func Get(httpReq Req, result interface{}) (err error) {
	//构建请求
	req, err := http.NewRequest("GET", httpReq.Url, nil)
	if err != nil {
		return err
	}
	if len(httpReq.Params) > 0 {
		q := req.URL.Query()
		for key, val := range httpReq.Params {
			valStr, _ := util.ConvertString(val)
			q.Add(key, valStr)
		}
		req.URL.RawQuery = q.Encode()
	}
	//如果设置了头部
	if httpReq.Header != nil && len(httpReq.Header) > 0 {
		for key, val := range httpReq.Header {
			req.Header.Set(key, val)
		}
	}
	//默认一秒超时
	//if httpReq.TimeOut.Milliseconds() <= 0 {
	//	httpReq.TimeOut = time.Second
	//}
	//生成http.client
	client := NewHttpClient(httpReq.TimeOut)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	//状态码
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("http status code is %d", resp.StatusCode))
	}
	//关闭响应
	defer resp.Body.Close()
	err = jsoniter.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

/*
	GET 请求
	url : 请求地址
	params ： 请求参数
	setHeader ： 设置头部
	result ： 返回结果
*/

func PostForm(httpReq Req, result interface{}) (err error) {
	postArgs := url.Values{}
	for key, val := range httpReq.Params {
		value, _ := util.ConvertString(val)
		postArgs.Set(key, value)
	}
	//构建请求
	req, err := http.NewRequest("POST", httpReq.Url, bytes.NewBuffer([]byte(postArgs.Encode())))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//如果设置了头部
	if httpReq.Header != nil && len(httpReq.Header) > 0 {
		for key, val := range httpReq.Header {
			req.Header.Set(key, val)
		}
	}
	//默认一秒超时
	if httpReq.TimeOut.Milliseconds() <= 0 {
		httpReq.TimeOut = time.Second
	}
	//生成http.client
	client := NewHttpClient(httpReq.TimeOut)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	//状态码
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("http status code is %d", resp.StatusCode))
	}
	//关闭响应
	defer resp.Body.Close()
	err = jsoniter.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}
	return nil
}
