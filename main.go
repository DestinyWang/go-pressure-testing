package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

type strArray []string

func (arr *strArray) String() string {
	return fmt.Sprint(*arr)
}

func (arr *strArray) Set(s string) error {
	*arr = append(*arr, s)
	return nil
}

func main() {
	//var err error
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	// 命令行
	var (
		conCurr  uint64   // 并发数
		totalNum uint64   // 请求数
		debug    bool     // debug 模式
		reqUrl   string   // 压测 URL
		path     string   // curl 文件路径, http 接口压测, 自定义参数设置
		verify   string   //
		headers  strArray // 自定义 header 信息
		body     string   // http request body
	)
	flag.Uint64Var(&conCurr, "c", 1, "并发数")
	flag.Uint64Var(&totalNum, "n", 1, "请求数")
	flag.BoolVar(&debug, "d", false, "调试模式")
	flag.StringVar(&reqUrl, "u", "", "压测地址")
	flag.StringVar(&path, "p", "", "文件路径")
	flag.StringVar(&verify, "v", "", "验证方法: http-statusCode/json")
	flag.Var(&headers, "H", "自定义 Header 信息, 实例: 'Content-Type: application/json'")
	flag.StringVar(&body, "D", "", "HTTP POST request body")
	flag.Parse()
	logrus.WithFields(logrus.Fields{
		"c": conCurr,
		"n": totalNum,
		"d": debug,
		"u": reqUrl,
		"p": path,
		"v": verify,
		"H": headers,
		"D": body,
	}).Infof("flags")
	if conCurr == 0 || totalNum == 0 || (reqUrl == "" && path == "") {
		fmt.Printf("参数不合法\n")
		flag.Usage()
	}
}
