package mipush

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	errorLog  *log.Logger
	ErrResult = errors.New("error result")

	v2URI      string = "https://api.xmpush.xiaomi.com/v2"
	v3URI      string = "https://api.xmpush.xiaomi.com/v3"
	defaultURI string = v2URI

	global *MiPush
)

func SetErrorLog(err *log.Logger) {
	errorLog = err
}

func logf(format string, args ...interface{}) {
	if errorLog != nil {
		errorLog.Printf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

type MiPush struct {
	appID     string
	appKey    string
	appSecret string
}

func Init(id, key, secret string) {
	global = NewMiPush(id, key, secret)
}

func NewMiPush(id, key, secret string) *MiPush {
	return &MiPush{id, key, secret}
}

/*
{
"result": "ok",
"trace_id": "Xcm13b72483633997339os",
"code": 0,
"data": {
"id": "scm13b72483633997343Bs"
},
"description": "成功",
"info": "Received push messages for 1 REGID"
}
*/
type result struct {
	Result      string `json:"result"`
	TraceID     string `json:"trace_id"`
	Code        int    `json:"code"`
	Description string `json:"description"`
	Info        string `json:"info"`
	Data        struct {
		ID string `json:"id"`
	} `json:"data"`
}

func httpPost(url string, params url.Values, headers map[string]string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	cli := &http.Client{
		Transport: tr,
	}

	var ioParams io.Reader
	ioParams = bytes.NewReader([]byte(params.Encode()))
	req, err := http.NewRequest(http.MethodPost, url, ioParams)

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		logf("cli.Do(req) error(%v)", err)
		return nil, err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logf("ioutil.ReadAll(resp.Body) error(%v)", err)
		return nil, err
	}
	return buff, nil
}
