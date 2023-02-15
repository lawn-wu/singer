package request

import (
	"meituan-api/domain"
)

type PoiMGetRequest struct {
	*BaseRequest
	PoiMGetParam
}

type PoiMGetParam struct {
	AppPoiCodes string `meituan:"app_poi_codes"`
}

func NewPoiMGetRequest(systemParam *domain.SystemParam) (poiMGetRequest PoiMGetRequest) {
	poiMGetRequest = PoiMGetRequest{
		BaseRequest: &BaseRequest{
			Uri:               "/api/v1/poi/mget",
			RequestMethodType: domain.REQUEST_METHOD_TYPE_GET,
			SystemParam:       systemParam,
		},
	}
	return
}
