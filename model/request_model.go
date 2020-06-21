package model

import (
	"github.com/DestinyWang/go-pressure-testing/util"
	"github.com/antlabs/pcurl"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	Url             string
	Form            string
	Method          string
	Headers         map[string]string
	Body            string
	Verify          string
	VerifyHttp      VerifyHttp
	VerifyWebSocket VerifyWebSocket
	Timeout         time.Duration
	Debug           bool
}

func (r *Request) GetBody() (body io.Reader) {
	return strings.NewReader(r.Body)
}

type VerifyHttp func(request *Request, response *http.Response) (code int, succ bool)
type VerifyWebSocket func(request *Request, seq string, msg []byte) (code int, succ bool)

func NewRequest(url string, verify string, timeout time.Duration, debug bool, path string, reqHeaders []string, reqBody string) (request *Request, err error) {
	var (
		method  string
		headers = make(map[string]string)
		body    string
	)
	if path != "" {
		// 有 path 的情况下优先从磁盘读取
		var curl *pcurl.Curl
		if curl, err = ParseCurlFromFile(path); err != nil {
			logrus.Errorf("[%s] ParseCurlFromFile fail: err=[%+v]", util.RunFuncName(), err)
			return nil, err
		}
		method = curl.Method
		if method == "" {
			method = "GET"
		}
		appendHeaderMap(curl.Header, headers)
		body = curl.Data
	} else {
		if reqBody != "" {
			method = "POST"
			body = reqBody
			headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
			appendHeaderMap(reqHeaders, headers)
		}
	}
	var (
		verifyHttp VerifyHttp
		//ok         bool
	)
	// http
	if verify == "" {
		verify = "statusCode"
	}
	//var key = fmt.Sprintf("%s.%s", "http", verify)
	//verifyHttp, ok =
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	return &Request{
		Url:             url,
		Form:            "",
		Method:          strings.ToUpper(method),
		Headers:         headers,
		Body:            body,
		Verify:          verify,
		VerifyHttp:      verifyHttp,
		VerifyWebSocket: nil,
		Timeout:         timeout,
		Debug:           debug,
	}, nil
}

func appendHeaderMap(headers []string, headerMap map[string]string) {
	for _, header := range headers {
		var index = strings.Index(header, ":")
		if index <= 0 {
			continue
		}
		headerMap[header[:index]] = header[index:]
	}
}
