package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jdlgj/core"
	"net/http"
)

//WxSession 微信登陆接口返回session
type WxSession struct {
	SessionKey string `json:"session_key"`
	ExpireIn   int    `json:"expires_in"`
	OpenID     string `json:"openid"`
}

//WxLogin 微信用户授权
func WxLogin(jscode string) (session WxSession, err error) {
	client := &http.Client{}
	appID := core.GetEnv("APP_ID", "")
	appSecret := core.GetEnv("APP_SECRET", "")
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appID, appSecret, jscode)

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, _ := client.Do(reqest)

	body, err := ioutil.ReadAll(response.Body)

	jsonStr := string(body)
	//解析json
	if err := json.Unmarshal(body, &session); err != nil {
		session.SessionKey = jsonStr
		return session, err
	}

	return session, err
}
