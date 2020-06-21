package model

import (
	"errors"
	"github.com/DestinyWang/go-pressure-testing/util"
	"github.com/antlabs/pcurl"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// curl 'https://www.baidu.com/s?ie=UTF-8&wd=go%20trimspace' \
//  -H 'Connection: keep-alive' \
//  -H 'Pragma: no-cache' \
//  -H 'Cache-Control: no-cache' \
//  -H 'Upgrade-Insecure-Requests: 1' \
//  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36' \
//  -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
//  -H 'Sec-Fetch-Site: same-origin' \
//  -H 'Sec-Fetch-Mode: navigate' \
//  -H 'Sec-Fetch-User: ?1' \
//  -H 'Sec-Fetch-Dest: document' \
//  -H 'Referer: https://www.baidu.com/' \
//  -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8' \
//  -H 'Cookie: BAIDUID=DF258E1BF033954823A4CB23154E3C97:FG=1; BIDUPSID=DF258E1BF033954823A4CB23154E3C97; PSTM=1536568770; MCITY=-%3A; BD_UPN=123253; BDUSS=mp1YmE2dUdRdFFLWDlQTGI5NU1XTlRqMlhKflZjOS1Xc0xxY24tVjU2VVNHLTVlRVFBQUFBJCQAAAAAAAAAAAEAAACCPEgft9~FrbXE1LK058nZxOoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABKOxl4SjsZeS; BDSFRCVID=ES0OJeC62GCC8JvuTXgXhwQl0uU8apcTH6aoFr5FvLR0p2adhvgAEG0PSf8g0KA-JGzVogKKKgOTHICF_2uxOjjg8UtVJeC6EG0Ptf8g0M5; H_BDCLCKID_SF=tb-qVCLKJIvhDRTvhCcjh-FSMgTBKI62aKDsKx8MBhcqEIL4QUbTeMIB5JOvKhv3ba6QoP5Gfnj2HUbSj4QoQRj0QljrQjOjtD6bBPQj0l5nhMJbb67JDMPFqt6xhq5y523iab3vQpnzEpQ3DRoWXPIqbN7P-p5Z5mAqKl0MLPbtbb0xXj_0-nDSHHLDJ58D3f; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; H_PS_PSSID=1460_31672_21121_32045_31321_30823_32111_22157; BD_HOME=1; delPer=0; BD_CK_SAM=1; PSINO=5; sugstore=1; H_PS_645EC=84f3fVlTBHX7oq1GLtiG7gqX9u%2Bnylf5mBIFSiinT3BJTrtXdTCGAe1Ixbw; BDSVRTM=0' \
//  --compressed
func ParseCurlFromFile(path string) (curl *pcurl.Curl, err error) {
	if path == "" {
		return nil, errors.New("path 不能为空")
	}
	var file *os.File
	if file, err = os.Open(path); err != nil {
		logrus.Errorf("[%s] open file fail: err=[%+v]", util.RunFuncName(), err)
		return nil, err
	}
	defer file.Close()
	var dataBytes []byte // 文件中的内容
	if dataBytes, err = ioutil.ReadAll(file); err != nil {
		logrus.Errorf("[%s] read file fail: err=[%+v]", util.RunFuncName(), err)
		return nil, err
	}
	curl = pcurl.ParseString(string(dataBytes))
	
	return
}

