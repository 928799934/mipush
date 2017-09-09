package mipush

import (
	"encoding/json"
	"net/url"
	"strings"
)

// 推送消息给指定的一些 regid
func SendRegIDs(message Message, packageName string, regIDs []string) (string, string, error) {
	return global.SendRegIDs(message, packageName, regIDs)
}

// 推送消息给指定的regid
func SendRegID(message Message, packageName, regID string) (string, string, error) {
	return global.SendRegID(message, packageName, regID)
}

// 推送消息给指定的一些 regid
func (this *MiPush) SendRegIDs(message Message, packageName string, regIDs []string) (string, string, error) {
	return this.SendRegID(message, packageName, strings.Join(regIDs, ","))
}

// 推送消息给指定的regid
func (this *MiPush) SendRegID(message Message, packageName, regID string) (string, string, error) {
	params := url.Values(message)
	params.Set("registration_id", regID)
	params.Set("restricted_package_name", packageName)
	uri := defaultURI + "/message/regid"
	headers := make(map[string]string)
	headers["Authorization"] = "key=" + this.appSecret
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	jsonData, err := httpPost(uri, params, headers)
	if err != nil {
		logf("httpPost(%v, %v, %v) error(%v)", uri, params, headers, err)
		return "", "", err
	}
	res := result{}
	if err := json.Unmarshal(jsonData, &res); err != nil {
		logf("json.Unmarshal(%s,&res) error(%v)", jsonData, err)
		return "", "", err
	}
	if res.Code != 0 {
		// 21301:认证失败
		// 21305:缺失必选参数
		// 10016:缺失必选参数
		logf("result(%s)", jsonData)
		return "", "", ErrResult
	}
	return res.TraceID, res.Data.ID, nil
}
