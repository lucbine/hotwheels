/*
@Time : 2020/7/13 9:49 下午
@Author : lucbine
@File : weixin.go
*/
package msg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type JSON struct {
	Access_token string `json:"access_token"`
}

type MESSAGES struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Text    struct {
		//Subject string `json:"subject"`
		Content string `json:"content"`
	} `json:"text"`
	Safe int `json:"safe"`
}

type WeiXinMsg struct {
}

var Qywx WeiXinMsg

func (en WeiXinMsg) Send(msg string) error {

}

func (en WeiXinMsg) GetAccessToken(corpid, corpsecret string) string {
	gettoken_url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpid + "&corpsecret=" + corpsecret
	//print(gettoken_url)
	client := &http.Client{}
	req, _ := client.Get(gettoken_url)
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	//fmt.Printf("\n%q",string(body))
	var json_str JSON
	json.Unmarshal([]byte(body), &json_str)
	//fmt.Printf("\n%q",json_str.Access_token)
	return json_str.Access_token
}

func (en WeiXinMsg) Send(msg string) error {

}

func (en WeiXinMsg) Send(msg string) error {

}
