package http

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const URL_HOME = "http://eds.newtouch.cn"
const HOST = "eds.newtouch.cn"
const USET_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36"
const ACCEPT_LANGUAGE = "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7"
const ACCEPT_ENCODING = "gzip, deflate"

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
	postProperties["connection"] = "Keep-Alive"
	postProperties["accept"] = "*/*"
	postProperties["user-agent"] = "Mozilla/5.0"

	// get
	getProperties["Upgrade-Insecure-Requests"] = "1"
	getProperties["User-Agent"] = USET_AGENT
	getProperties["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	getProperties["Accept-Encoding"] = ACCEPT_ENCODING
	getProperties["Accept-Language"] = ACCEPT_LANGUAGE
	getProperties["Host"] = HOST
}

// http请求，返回 string
func DoRequest(url, method, cookie string, body io.Reader) string {
	// log.Println(url)
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

	request.Header.Set("Cookie", cookie)

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 查看一下
	// io.Copy(os.Stdout, resp.Body)

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
