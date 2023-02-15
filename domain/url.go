package domain

import (
	sort2 "sort"
	"strings"
)

type Url struct {
	UrlPrefix string
}

var UrlApp = new(Url)

func (r *Url) GenUrlPrefix(uri string) string {
	if r.UrlPrefix == "" {
		env := ENV
		switch env {
		case 0:
			r.UrlPrefix = "http://openapi.b.waimai.test.sankuai.com"
		case 1:
			r.UrlPrefix = "https://waimaiopen.meituan.com"
		case 2:
			r.UrlPrefix = "http://127.0.0.1:9000"
		case 3:
			r.UrlPrefix = "http://openapi.b.waimai.dev.sankuai.com"
		case 4:
			r.UrlPrefix = "http://openapi.waimai.st.sankuai.com"
		case 5:
			r.UrlPrefix = "http://liushikuan-lzdod-sl-e.openapi.waimai.test.sankuai.com"
		}

		if URL != "" {
			r.UrlPrefix = URL
		}
	}

	return r.UrlPrefix + uri
}

func (r *Url) GenUrlForGenSig(uri string, systemParams SortedMaps, applicationParams SortedMaps) (baseUrl string) {
	//params := append(systemParams, applicationParams...)
	var params SortedMaps
	for _, systemParam := range systemParams {
		params = append(params, SortedMap{
			Key:   systemParam.Key,
			Value: systemParam.Value,
		})
	}
	for _, applicationParam := range applicationParams {
		params = append(params, SortedMap{
			Key:   applicationParam.Key,
			Value: applicationParam.Value,
		})
	}
	concatStr := r.ConcatParams(params)
	baseUrl = r.GenUrlPrefix(uri) + "?" + concatStr + systemParams.Get("appSecret")

	return
}

func (*Url) ConcatParams(params SortedMaps) (concatStr string) {
	sort2.Sort(params)

	for _, param := range params {
		if param.Key != "appSecret" {
			concatStr += "&" + param.Key + "=" + param.Value
		}
	}

	concatStr = strings.TrimLeft(concatStr, "&")

	return
}

func (r *Url) GenOnlyHasSysParamsAndSigUrl(urlPrefix string, systemParams SortedMaps, sig string) (baseUrl string) {
	concatStr := r.ConcatParams(systemParams)
	baseUrl = urlPrefix + "?" + concatStr + "&sig=" + sig
	return
}

func (r *Url) GenUrlForGetRequest(urlHasParamsNoSig string, sig string) (baseUrl string) {
	baseUrl = urlHasParamsNoSig + "&sig=" + sig
	return
}
