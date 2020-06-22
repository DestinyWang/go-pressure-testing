package golink

import (
	"context"
	"crypto/tls"
	"github.com/DestinyWang/go-pressure-testing/model"
	"github.com/DestinyWang/go-pressure-testing/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

// 发送 HTTP 请求
// <param chanId>
// <param ch>
// <param chanId>
// <param wg>
// <param request>
func Http(ctx context.Context, chanId uint64, ch chan *model.RequestResult, totalNum uint64, wg sync.WaitGroup, req *model.Request) {
	defer wg.Done()
	for n := uint64(0); n < totalNum; n++ {
		// 每个协程请求 n 次
		var (
			//startTime = time.Now()
			isSucceed = true
			errCode   = http.StatusOK
			err       error
		)
		var client = http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		var httpReq *http.Request
		if httpReq, err = http.NewRequest(req.Method, req.Url, req.GetBody()); err != nil {
			logrus.Errorf("[%s] NewRequest fail: req=[%s], err=[%+v]", util.RunFuncName(), req, err)
			return
		}
		if _, ok := req.Headers["Content-Type"]; !ok {
			if req.Headers == nil {
				req.Headers = make(map[string]string)
			}
			req.Headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
		}
		for k, v := range req.Headers {
			httpReq.Header.Set(k, v)
		}
		var resp *http.Response
		if resp, err = client.Do(httpReq); err != nil {
			logrus.Errorf("[%s] http fail: err=[%+v]", util.RunFuncName(), err)
			isSucceed = false
			errCode = resp.StatusCode
		}
		
		var result = &model.RequestResult{
			ChanId:    chanId,
			Time:      0,
			IsSucceed: isSucceed,
			ErrCode:   errCode,
		}
		result.SetId(chanId, n)
		ch <- result
	}
}
