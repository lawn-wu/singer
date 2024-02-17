package request

import (
	"meituan-api/domain"
)

type PoiOpenRequest struct {
	*BaseRequest
}

func NewPoiOpenRequest(systemParam *domain.SystemParam) (request PoiOpenRequest) {
	request = PoiOpenRequest{
		BaseRequest: &BaseRequest{
			Uri:               "/api/v1/poi/open",
			RequestMethodType: domain.REQUEST_METHOD_TYPE_POST,
			SystemParam:       systemParam,
		},
	}

	return
}
