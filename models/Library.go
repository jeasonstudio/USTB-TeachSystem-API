package models

import (
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"github.com/levigross/grequests"
)

type userInfo struct {
	Name         string `json:"name"`
	MissDate     string `json:"missDate"`
	MaxBBook     string `json:"maxBBook"`
	ReaderType   string `json:"readerType"`
	IllegalTimes string `json:"illegalTimes"`
	UserID       string `json:"userID"`
	StartDate    string `json:"startDate"`
	MaxABook     string `json:"maxABook"`
	BorrowLevel  string `json:"borrowLevel"`
	OweMoney     string `json:"oweMoney"`
	ID           string `json:"id"`
	GetTime      string `json:"getTime"`
	AllBSum      string `json:"allBSum"`
	From         string `json:"from"`
	IDCardNum    string `json:"idCardNum"`
	Gender       string `json:"gender"`
}

func LibLoginInfo(username string, password string) map[string]interface{} {

	getLibReaderLoginURL := beego.AppConfig.String("LIB_READER_LOGIN")
	getCodeURL := beego.AppConfig.String("GET_CODE")
	getLibBaseURL := beego.AppConfig.String("LIB_LOGIN")
	getLibInfoURL := beego.AppConfig.String("LIB_INFO")

	s := grequests.NewSession(nil)
	s.Get(getLibReaderLoginURL, nil)
	//获取验证码
	s.Get(getCodeURL, nil)
	option := &grequests.RequestOptions{
		Params: map[string]string{"number": username,
			"passwd": password, "captcha": "1",
			"select": "cert_no", "returnUrl": ""},
	}
	s.Post(getLibBaseURL, option)
	resp, _ := s.Get(getLibInfoURL, nil)
	beego.Alert(resp.String())
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	var thisUser userInfo

	doc.Find("TD").Each(func(i int, a *goquery.Selection) {
		beego.Alert(a.Text())
	})

	// beego.Alert(resp.String())

	a, _ := json.Marshal(thisUser)

	// beego.Alert(string(a))
	var final map[string]interface{}
	Terr := json.Unmarshal(a, &final)
	if nil != Terr {
		beego.Alert(Terr)
	}

	return final
}

// func getInfo(se *Session) {

// }
