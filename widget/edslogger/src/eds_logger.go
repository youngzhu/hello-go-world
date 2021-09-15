package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	myhttp "http"
	"secret"

	"github.com/PuerkitoBio/goquery"
)

var secretInfo *secret.Secret
var err error
var cookie string

func init() {
	secretInfo, err = secret.RetrieveSecret()
	if err != nil {
		log.Fatal(err)
	}

	cookie = secretInfo.Cookie
}

func login() error {
	// 检验网站是否正常
	resp, err := http.Head(myhttp.URL_HOME) // 只请求网站的 http header信息
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	loginUrl := "http://eds.newtouch.cn/eds3/DefaultLogin.aspx?lan=zh-cn"
	// 登录
	// data := `{"UserId":"###", "UserPsd":"***"}`
	// data := "UserId=###&UserPsd=***"
	// params := url.Values{
	// 	"UserId":  {"###"},
	// 	"UserPsd": {"***"},
	// }
	params := url.Values{}
	params.Add("UserId", secretInfo.UserId)
	params.Add("UserPsd", secretInfo.UserPsd)
	// var request *http.Request
	// request, err = http.NewRequest(http.MethodPost, URL_LOGIN, strings.NewReader(data))
	// request, err = http.NewRequest(http.MethodPost, loginUrl, strings.NewReader(params.Encode()))

	myhttp.DoRequest(loginUrl, http.MethodPost, cookie, strings.NewReader(params.Encode()))

	return nil
}

func workLog(logDate string) {
	url := "http://eds.newtouch.cn/eds3/worklog.aspx?tabid=0&LogDate=" + logDate

	// 先通过get获取一些隐藏参数，用作后台校验
	hiddenParams := getHiddenParams(url)
	// fmt.Println(hiddenParams)

	doWorkLog(url, logDate, "AM", hiddenParams)
	doWorkLog(url, logDate, "PM", hiddenParams)
}

func doWorkLog(workLogUrl, logDate, timeFlag string, hiddenParams map[string]string) {

	startTime := "10:00"
	endTime := "12:00"
	if "PM" == timeFlag {
		startTime = "13:00"
		endTime = "18:00"
	}

	logParams := url.Values{}
	logParams.Set("__EVENTTARGET", "hplbWorkType")
	logParams.Set("__EVENTARGUMENT", "")
	logParams.Set("__LASTFOCUS", "")
	logParams.Set("__VIEWSTATEGENERATOR", "3A8BE513")
	logParams.Set("txtDate", logDate)
	logParams.Set("txtStartTime", startTime)
	logParams.Set("txtEndTime", endTime)
	logParams.Set("ddlProjectList", "10868")
	logParams.Set("hplbWorkType", "0106")
	logParams.Set("hplbAction", "010601")
	logParams.Set("TextBox1", "")
	logParams.Set("txtMemo", "编码与测试")
	logParams.Set("btnSave", "+%E7%A1%AE+%E5%AE%9A+")
	logParams.Set("txtnodate", logDate)
	logParams.Set("txtnoStartTime", startTime)
	logParams.Set("txtnoEndTime", endTime)
	logParams.Set("TextBox6", "")
	logParams.Set("txtnoMemo", "")
	logParams.Set("txtCRMDate", logDate)
	logParams.Set("txtCRMStartTime", startTime)
	logParams.Set("txtCRMEndTime", endTime)
	logParams.Set("TextBox5", "")
	logParams.Set("txtCRMMemo", "")

	for key, value := range hiddenParams {
		logParams.Set(key, value)
	}

	myhttp.DoRequest(workLogUrl, http.MethodPost, cookie, strings.NewReader(logParams.Encode()))

	log.Println("日志操作成功")
}

func workWeeklyLog(logDate string) {
	logUrl := "http://eds.newtouch.cn/eds36web/WorkWeekly/WorkWeeklyInfo.aspx"

	// 先通过get获取一些隐藏参数，用作后台校验
	hiddenParams := getHiddenParams(logUrl)

	logParams := url.Values{}
	logParams.Set("hidCurrRole", "")
	logParams.Set("hidWeeklyState", "")
	logParams.Set("WeekReportDate", logDate)
	logParams.Set("txtWorkContent", "编码与测试")
	logParams.Set("txtStudyContent", "算法与数据结构")
	logParams.Set("txtSummary", "算法的性能不能不考虑，也不能过度优化")
	logParams.Set("txtPlanWork", "对常量的重构")
	logParams.Set("txtPlanStudy", "Oracle")
	logParams.Set("btnSubmit", "%E6%8F%90%E4%BA%A4")

	for key, value := range hiddenParams {
		logParams.Set(key, value)
	}

	myhttp.DoRequest(logUrl, http.MethodPost, cookie, strings.NewReader(logParams.Encode()))

	log.Println("周报填写成功")
}

func getHiddenParams(url string) map[string]string {
	result := make(map[string]string)

	respHtml := myhttp.DoRequest(url, http.MethodGet, cookie, nil)
	// fmt.Println(respHtml)

	keys := []string{"__EVENTVALIDATION", "__VIEWSTATE"}

	for _, k := range keys {
		result[k] = getValueFromHtml(respHtml, k)
	}

	return result
}

func getValueFromHtml(html, key string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		log.Fatalln(err)
	}

	var value = ""
	doc.Find("input").Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("id")
		if id == key {
			value, _ = s.Attr("value")
			// fmt.Println("i", i, "选中的文本", value)
			return
		}

	})

	return value
}

func main() {

	err := login()
	if err != nil {
		// 正常返回还不行，需要有错误发送邮件通知
		// return
		log.Fatalln("网站服务错误", err)
		// os.Exit(1)
	}
	log.Println("登陆成功")

	// workLog("2021-09-18")

	time.Sleep(5 * time.Second)

	workWeeklyLog("2021-09-13")

}
