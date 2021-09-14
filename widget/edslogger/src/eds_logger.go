package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

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
	fmt.Println(hiddenParams)
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

	workLog("2021-09-10")

}
