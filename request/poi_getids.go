package request

import (
	"meituan-api/domain"
)

type PoiGetIdsRequest struct {
	*BaseRequest
}

func NewPoiGetIdsRequest(systemParam *domain.SystemParam) (request PoiGetIdsRequest) {
	request = PoiGetIdsRequest{
		BaseRequest: &BaseRequest{
			Uri:               "/api/v1/poi/getids",
			RequestMethodType: domain.REQUEST_METHOD_TYPE_GET,
			SystemParam:       systemParam,
		},
	}

	return
}
