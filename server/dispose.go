package server

import (
	"context"
	"github.com/DestinyWang/go-pressure-testing/model"
	"github.com/DestinyWang/go-pressure-testing/server/golink"
	"sync"
)

// <param concurrency>: 并发数
// <param totalNum>:
func Dispose(ctx context.Context, concurrency, totalNum uint64, req *model.Request) {
	// 设置数据缓存
	var ch = make(chan *model.RequestResult, 1000)
	var (
		wg   sync.WaitGroup // 发送数据完成
		wgRe sync.WaitGroup // 数据处理完成
	)
	wgRe.Add(1)
	
	for i := uint64(0); i < concurrency; i++ {
		wg.Add(1)
		switch req.Form {
		case model.FormTypeHttp:
		// http 请求
		golink.DoHttp(ctx, i, ch, totalNum, wg, req)
		default:
			// 类型不支持
			wg.Done()
		}
	}
}
