package main

import (
	"encoding/json"
	"fmt"
	captcha "github.com/mojocn/base64Captcha"
	"log"
	"net/http"
)

var store = captcha.DefaultMemStore

func NewDriver() *captcha.DriverString {
	driver := new(captcha.DriverString)
	driver.Height = 44
	driver.Width = 120
	driver.NoiseCount = 5
	driver.ShowLineOptions = captcha.OptionShowSineLine | captcha.OptionShowSlimeLine | captcha.OptionShowHollowLine
	driver.Length = 6
	driver.Source = "1234567890qwertyuipkjhgfdsazxcvbnm"
	driver.Fonts = []string{"wqy-microhei.ttc"}
	return driver
}

// 生成图形验证码
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	var driver = NewDriver().ConvertFonts()
	c := captcha.NewCaptcha(driver, store)
	_, content, answer := c.Driver.GenerateIdQuestionAnswer()
	id := "captcha:med"
	item, _ := c.Driver.DrawCaptcha(content)
	c.Store.Set(id, answer)
	item.WriteTo(w)
}

// 验证
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {

	id := "captcha:med"
	code := r.FormValue("code")
	body := map[string]interface{}{"code": 1000, "msg": "failed"}
	if store.Verify(id, code, true) {
		body = map[string]interface{}{"code": 1001, "msg": "ok"}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

func main() {
	http.HandleFunc("/captcha", generateCaptchaHandler)
	http.HandleFunc("/captcha/verify", captchaVerifyHandle)

	fmt.Println("Server is at :8089")
	if err := http.ListenAndServe(":8089", nil); err != nil {
		log.Fatal(err)
	}
}