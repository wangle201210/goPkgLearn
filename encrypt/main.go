package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	appId = "xrJpl4e7"
	message = "appId%spreSign%stimestamp%d"
	secret  = "****"
	shanyanValidateUrl = "https://api.253.com/open/flashsdk/mobile-validate-web"
	mobile = "18223593333"
)

type myFun struct {

}
type Res struct {
	AppId string
	Sign string
	timestamp int64
}

type validateReq struct {
	appId string
	token string
	mobile string
	sign string
}

type validateResp struct {
	code string
	message string
	chargeStatus string
	sign string
	data validateData
}

type validateData struct {
	isVerify string
	tradeNo string
}

func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	//	hex.EncodeToString(h.Sum(nil)) //转16进制
	return base64.StdEncoding.EncodeToString([]byte(sha))
}

func shanyanValidate(token,mobile,sign string) (r bool) {
	data := validateReq{
		appId: appId,
		token: token,
		mobile: mobile,
		sign: sign,
	}
	res := new(validateResp)
	post := Post(shanyanValidateUrl, data, "application/x-www-form-urlencoded")
	if err := json.Unmarshal(post, res); err != nil {
		fmt.Printf("err(%+v)",err)
		return
	}
	fmt.Printf("data(%+v)",res)
	return true
}

func Post(url string, data interface{}, contentType string) []byte {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return result
}

func (f myFun) getSign(preSign string)  {
	now := time.Now().Unix()
	preSign = "T25lS2V5TG9naW5NYW5hZ2VyX2NhbGxiYWNrXzE2MDY0NTgyNzEyNDZ8MS4wLjF8S1BEM0xCMTlEMFNWMDMwNUFCSUhKSjkwOUJQRzFPTEd8ZGZkNTNiN2UxNWYzMDQxZDg4MjBlNTg5ZmEwYjQxNzQ="
	mes := fmt.Sprintf(message,appId,preSign,now)
	sign := ComputeHmacSha256(mes, secret)
	r := Res{
		AppId: appId,
		Sign: sign,
		timestamp: now,
	}
	fmt.Printf("%+v",r)
}

func (f myFun) validate(token,mobile,sign string) bool {
	return shanyanValidate(token, mobile, sign)
}

func main() {
	f := myFun{}
	f.getSign("****")
	f.validate("****","****","****")
}
