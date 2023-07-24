package request

import (
	"meituan-api/domain"
)

type PoiGetIdsRequest struct {
	*BaseRequest
}

func NewGetIdsRequest(systemParam *domain.SystemParam) (poiMGetRequest PoiGetIdsRequest) {
	poiMGetRequest = PoiGetIdsRequest{
		BaseRequest: &BaseRequest{
			Uri:               "/api/v1/poi/getids",
			RequestMethodType: domain.REQUEST_METHOD_TYPE_GET,
			SystemParam:       systemParam,
		},
	}
	return
}
