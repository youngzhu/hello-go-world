package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const URL_HOME = "http://eds.newtouch.cn"
const URL_LOGIN = "http://eds.newtouch.cn/eds3/DefaultLogin.aspx?lan=zh-cn"
const HOST = "eds.newtouch.cn"
const USET_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36"
const ACCEPT_LANGUAGE = "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7"
const ACCEPT_ENCODING = "gzip, deflate"

// 稍后隐藏
const COOKIE = "ASP.NET_SessionId=4khtnz55xiyhbmncrzmzyzzc; ActionSelect=010601; Hm_lvt_416c770ac83a9d996d7b3793f8c4994d=1569767826; Hm_lpvt_416c770ac83a9d996d7b3793f8c4994d=1569767826; PersonId=12234"

var postProperties = make(map[string]string)
var getProperties = make(map[string]string)

func init() {
	// post
	postProperties["Host"] = HOST
	postProperties["Content-Length"] = "6955"
	postProperties["Cache-Control"] = "max-age=0"
	postProperties["Origin"] = URL_HOME
	postProperties["Upgrade-Insecure-Requests"] = "1"
	postProperties["Content-Type"] = "application/x-www-form-urlencoded"
	postProperties["User-Agent"] = "Mozilla/5.0"
	postProperties["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"
	postProperties["Accept-Encoding"] = ACCEPT_ENCODING
	postProperties["Accept-Language"] = ACCEPT_LANGUAGE
	postProperties["Cookie"] = COOKIE
	postProperties["connection"] = "Keep-Alive"
	postProperties["accept"] = "*/*"
	postProperties["user-agent"] = "Mozilla/5.0"

	// get
	getProperties["Upgrade-Insecure-Requests"] = "1"
	getProperties["User-Agent"] = USET_AGENT
	getProperties["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	getProperties["Accept-Encoding"] = ACCEPT_ENCODING
	getProperties["Accept-Language"] = ACCEPT_LANGUAGE
	getProperties["Cookie"] = COOKIE
	getProperties["Host"] = HOST
}

// http请求，返回 string
func doRequest(url string, method string, body io.Reader) string {
	log.Println(url)
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Referer", url)

	var headerProps map[string]string
	if method == http.MethodPost {
		headerProps = postProperties
	} else {
		headerProps = getProperties
	}
	for key, value := range headerProps {
		request.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// log.Println(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		// io.Copy(os.Stdout, resp.Body)
		log.Panicln(resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(respBody)
}

func login() error {
	// 检验网站是否正常
	resp, err := http.Head(URL_HOME) // 只请求网站的 http header信息
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// 登录
	// data := `{"UserId":"12234", "UserPsd":"young122345"}`
	data := "UserId=12234&UserPsd=young12234"
	// params := url.Values{
	// 	"UserId":  {"12234"},
	// 	"UserPsd": {"young12234"},
	// }
	params := url.Values{}
	params.Add("UserId", "12234")
	params.Add("UserPsd", "young12234")
	var request *http.Request
	request, err = http.NewRequest(http.MethodPost, URL_LOGIN, strings.NewReader(data))
	if err != nil {
		return err
	}
	request.Header.Set("Referer", URL_LOGIN)
	for key, value := range postProperties {
		request.Header.Set(key, value)
	}

	resp, err = http.DefaultClient.Do(request)

	if err != nil {
		return nil
	}

	log.Println(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		panic(resp.StatusCode)
	}

	// 查看一下
	// io.Copy(os.Stdout, resp.Body)

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

	respHtml := doRequest(url, http.MethodGet, nil)
	fmt.Println(respHtml)

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
