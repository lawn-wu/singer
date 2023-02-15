package request

import (
	"encoding/json"
	"errors"
	"meituan-api/domain"
	"meituan-api/response"
	"meituan-api/util"
)

type BaseRequest struct {
	Uri               string
	RequestMethodType domain.RequestMethodType
	SystemParam       *domain.SystemParam
	ApplicationParams domain.SortedMaps
}

func (r *BaseRequest) DoRequest() (requestBody []byte, err error) {
	if r.SystemParam == nil {
		err = errors.New("SystemParam is nil")
		return
	}

	systemParamMap := util.ConvertSystemParamToSortedMap(r.SystemParam)
	urlForGenSig := domain.UrlApp.GenUrlForGenSig(r.Uri, systemParamMap, r.ApplicationParams)
	sig := util.GenSig(urlForGenSig)
	urlPrefix := domain.UrlApp.GenUrlPrefix(r.Uri)
	requestBody, err = domain.Request(urlPrefix, r.GenUrlForGetRequest(urlPrefix, systemParamMap, r.ApplicationParams),
		sig, systemParamMap, r.ApplicationParams, r.RequestMethodType)
	if err != nil {
		return
	}

	var responseParam response.SystemResponse
	err = json.Unmarshal(requestBody, &responseParam)
	if err != nil {
		return
	}

	if responseParam.Error.Msg != "" {
		err = errors.New(responseParam.Error.Msg)
		return
	}

	return
}

func (r *BaseRequest) SetApplicationParams(reflectObject any) {
	r.ApplicationParams = util.GetStructField(reflectObject, "meituan")
}

func (r *BaseRequest) GenUrlForGetRequest(
	urlPrefix string,
	systemParams domain.SortedMaps,
	applicationParams domain.SortedMaps,
) (baseUrl string) {
	uriParamStr := "app_id=" + systemParams.Get("app_id") + "&timestamp=" + systemParams.Get("timestamp")
	accessToken := systemParams.Get("access_token")
	if accessToken != "" {
		uriParamStr = uriParamStr + "&" + "access_token" + "=" + accessToken
	}

	if len(applicationParams) > 0 {
		for _, param := range applicationParams {
			if param.Value != "" {
				uriParamStr = uriParamStr + "&" + param.Key + "=" + param.Value
			}
		}
	}

	baseUrl = urlPrefix + "?" + uriParamStr

	return baseUrl
}
