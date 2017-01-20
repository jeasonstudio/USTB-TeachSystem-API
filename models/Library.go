package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/astaxie/beego"
)

func getCode(username string) string {
	return "s"
}
func LibLogin(username string, password string, code string) map[string]interface{} {
	code = getCode(username)
	beego.Alert(code)

	getLibBaseURL := beego.AppConfig.String("LIB_LOGIN")

	v := url.Values{"number": {username}, "passwd": {password}, "captcha": {code}, "select": {"cert_no"}, "returnUrl": {""}}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, getLibBaseURL, body)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	beego.Alert(string(data))

	var a = make(map[string]interface{})

	return a
}
