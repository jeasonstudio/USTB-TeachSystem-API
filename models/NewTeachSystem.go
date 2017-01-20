package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	simplejson "github.com/bitly/go-simplejson"
)

type cxScore struct {
	Semestre   string `json:"semestre"`
	CXType     string `json:"cxType"`
	Name       string `json:"name"`
	Score      string `json:"score"`
	InsertTime string `json:"insertTime"`
}

type classScore struct {
	GPA     string      `json:"gpa"`
	AvScore string      `json:"avScore"`
	Body    []bodyScore `json:"body"`
}

type bodyScore struct {
	Semestre   string `json:"semestre"`
	ClassNum   string `json:"classNum"`
	ClassName  string `json:"className"`
	ClassType  string `json:"classType"`
	LearnHour  string `json:"learnHour"`
	StuScore   string `json:"stuScore"`
	FirstScore string `json:"firstScore"`
	FinalScore string `json:"finalScore"`
	Flag       string `json:"flag"`
}

type examTime struct {
	ClassNum     string `json:"classNum"`
	ClassName    string `json:"className"`
	ExamTime     string `json:"examTime"`
	ExamLocation string `json:"examLocation"`
	Info         string `json:"info"`
}

type cetScore struct {
	LangLevel   string `json:"langLevel"`
	Card        string `json:"card"`
	ListenScore string `json:"listenScore"`
	ReadScore   string `json:"readScore"`
	WriteScore  string `json:"writeScore"`
	OtherScore  string `json:"otherScore"`
	AllScore    string `json:"allScore"`
	Date        string `json:"date"`
	Info        string `json:"info"`
}

func login(userName string, password string) string {
	LoginURL := beego.AppConfig.String("SYSTEM_LOGIN")

	v := url.Values{"j_username": {userName + ",undergraduate"}, "j_password": {password}}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, LoginURL, body)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	// data, _ := ioutil.ReadAll(resp.Body)

	res := fmt.Sprintf("%s", resp.Request.URL)
	tagCookies := strings.Split(strings.Split(res, ";")[1], "=")[1]

	return tagCookies
}

func GetCourseFromLogin(userName string, password string, semestre string) map[string]interface{} {
	cookie := login(userName, password)

	return getTrueCourse(cookie, userName, semestre)
}

func getTrueCourse(thatCookie string, userName string, semestre string) map[string]interface{} {
	getTrueCourseURL := beego.AppConfig.String("COURSE_TABLE")

	v := url.Values{"listXnxq": {semestre}, "uid": {userName}}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, getTrueCourseURL, body)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "JSESSIONID="+thatCookie)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	json, err := simplejson.NewJson(data)
	var nodes = make(map[string]interface{})

	nodes, _ = json.Map()

	return nodes
}

func GetCXScoreFromLogin(userName string, password string) []map[string]interface{} {
	cookie := login(userName, password)

	return getTrueCXScore(cookie, userName)
}

func getTrueCXScore(thatCookie string, userName string) []map[string]interface{} {
	getTrueCourseURL := beego.AppConfig.String("INNVOATION_SCORE")

	v := url.Values{"uid": {userName}}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, getTrueCourseURL, body)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "JSESSIONID="+thatCookie)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	str := strings.NewReader(string(data))

	doc, _ := goquery.NewDocumentFromReader(str)

	var finalCXScore []cxScore

	doc.Find(".gridtable tbody tr").Each(func(i int, a *goquery.Selection) {
		var thisCXScore cxScore
		a.Find("td").Each(func(j int, m *goquery.Selection) {
			switch j {
			case 0:
				thisCXScore.Semestre = m.Text()
				break
			case 1:
				thisCXScore.CXType = m.Text()
				break
			case 2:
				thisCXScore.Name = m.Text()
				break
			case 3:
				thisCXScore.Score = m.Text()
				break
			case 4:
				thisCXScore.InsertTime = m.Text()
				break
			}
		})
		finalCXScore = append(finalCXScore, thisCXScore)
	})
	a, _ := json.Marshal(finalCXScore)

	// beego.Alert(string(a))
	var final []map[string]interface{}
	Terr := json.Unmarshal(a, &final)
	if nil != Terr {
		beego.Alert(Terr)
	}
	return final
}

func GetClassScoreFromLogin(userName string, password string) map[string]interface{} {
	cookie := login(userName, password)

	return getTrueClassScore(cookie, userName)
}

func getTrueClassScore(thatCookie string, userName string) map[string]interface{} {
	getTrueClassScoreURL := beego.AppConfig.String("ALL_COURSE_SCORE")

	v := url.Values{"uid": {userName}}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, getTrueClassScoreURL, body)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "JSESSIONID="+thatCookie)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	str := strings.NewReader(string(data))

	doc, _ := goquery.NewDocumentFromReader(str)

	var finalClassScore classScore
	var bodyClassScore []bodyScore

	doc.Find(".gridtable tbody tr").Each(func(i int, a *goquery.Selection) {
		var thisClassScore bodyScore
		a.Find("td").Each(func(j int, m *goquery.Selection) {
			switch j {
			case 0:
				thisClassScore.Semestre = m.Text()
				break
			case 1:
				thisClassScore.ClassNum = m.Text()
				break
			case 2:
				thisClassScore.ClassName = m.Text()
				break
			case 3:
				thisClassScore.ClassType = m.Text()
				break
			case 4:
				thisClassScore.LearnHour = m.Text()
				break
			case 5:
				thisClassScore.StuScore = m.Text()
				break
			case 6:
				thisClassScore.FirstScore = m.Text()
				break
			case 7:
				thisClassScore.FinalScore = m.Text()
				break
			case 8:
				thisClassScore.Flag = m.Text()
				break
			}
		})
		bodyClassScore = append(bodyClassScore, thisClassScore)
	})

	doc.Find("h5").Each(func(i int, a *goquery.Selection) {
		switch i {
		case 0:
			finalClassScore.GPA = strings.Split(a.Text(), ":")[1]
		case 1:
			finalClassScore.AvScore = strings.Split(a.Text(), ":")[1]
		}
	})
	// beego.Alert(doc.Find("h5").Text())

	finalClassScore.Body = bodyClassScore
	a, _ := json.Marshal(finalClassScore)

	// beego.Alert(string(a))
	var final map[string]interface{}
	Terr := json.Unmarshal(a, &final)
	if nil != Terr {
		beego.Alert(Terr)
	}
	return final
}

func GetExamTimeFromLogin(userName string, password string, semestre string) []map[string]interface{} {
	cookie := login(userName, password)

	return getTrueExamTimeScore(cookie, userName, semestre)
}

func getTrueExamTimeScore(thatCookie string, userName string, semestre string) []map[string]interface{} {
	getTrueExamTimeScoreURL := beego.AppConfig.String("EXAM_TIMELOCATION")

	v := url.Values{"uid": {userName}, "winName": {"examListPanel"}, "listXnxq": {semestre}}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, getTrueExamTimeScoreURL, body)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "JSESSIONID="+thatCookie)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	str := strings.NewReader(string(data))

	doc, _ := goquery.NewDocumentFromReader(str)

	var bodyExamTime []examTime

	doc.Find(".gridtable tbody tr").Each(func(i int, a *goquery.Selection) {
		var thisExamTime examTime
		a.Find("td").Each(func(j int, m *goquery.Selection) {
			switch j {
			case 0:
				thisExamTime.ClassNum = m.Text()
				break
			case 1:
				thisExamTime.ClassName = m.Text()
				break
			case 2:
				thisExamTime.ExamTime = m.Text()
				break
			case 3:
				thisExamTime.ExamLocation = m.Text()
				break
			case 4:
				thisExamTime.Info = m.Text()
				break
			}
		})
		bodyExamTime = append(bodyExamTime, thisExamTime)
	})

	a, _ := json.Marshal(bodyExamTime)

	// beego.Alert(string(a))
	var final []map[string]interface{}
	Terr := json.Unmarshal(a, &final)
	if nil != Terr {
		beego.Alert(Terr)
	}
	return final
}

// 四六级成绩
func GetCETScoreFromLogin(userName string, password string) []map[string]interface{} {
	cookie := login(userName, password)

	return getTrueCETScoreScore(cookie, userName)
}

func getTrueCETScoreScore(thatCookie string, userName string) []map[string]interface{} {
	getTrueCETScoreURL := beego.AppConfig.String("CET_SCORE")

	v := url.Values{"uid": {userName}}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, getTrueCETScoreURL, body)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "JSESSIONID="+thatCookie)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	str := strings.NewReader(string(data))

	doc, _ := goquery.NewDocumentFromReader(str)

	var bodyCetScore []cetScore

	doc.Find(".gridtable tbody tr").Each(func(i int, a *goquery.Selection) {
		var thisCetScore cetScore
		a.Find("td").Each(func(j int, m *goquery.Selection) {
			switch j {
			case 0:
				thisCetScore.LangLevel = m.Text()
				break
			case 1:
				thisCetScore.Card = m.Text()
				break
			case 2:
				thisCetScore.ListenScore = m.Text()
				break
			case 3:
				thisCetScore.ReadScore = m.Text()
				break
			case 4:
				thisCetScore.WriteScore = m.Text()
				break
			case 5:
				thisCetScore.OtherScore = m.Text()
				break
			case 6:
				thisCetScore.AllScore = m.Text()
				break
			case 7:
				thisCetScore.Date = m.Text()
				break
			case 8:
				thisCetScore.Info = m.Text()
				break
			}
		})
		bodyCetScore = append(bodyCetScore, thisCetScore)
	})

	a, _ := json.Marshal(bodyCetScore)

	// beego.Alert(string(a))
	var final []map[string]interface{}
	Terr := json.Unmarshal(a, &final)
	if nil != Terr {
		beego.Alert(Terr)
	}
	return final
}
