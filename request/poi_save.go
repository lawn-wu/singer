package request

import (
	"meituan-api/domain"
)

type PoiSaveRequest struct {
	*BaseRequest
	PoiSaveParam
}

type PoiSaveParam struct {
	AppPoiCode string `meituan:"app_poi_code"`
	Name string `meituan:"name"`
	Address string `meituan:"address"`
	Latitude string `meituan:"latitude"`
	Longitude string `meituan:"latitude"`
	PicUrl string `meituan:"pic_url"`
	PicUrlLarge string `meituan:"pic_url_large"`
	Phone string `meituan:"phone"`
	StandbyTel string `meituan:"standby_tel"`
	ShippingFee float64 `meituan:"shipping_fee"`
	ShippingTime string `meituan:"shipping_time"`
	PromotionInfo string `meituan:"promotion_info"`
	OpenLevel int `meituan:"open_level"`
	IsOnline int `meituan:"is_online"`
	InvoiceSupport int `meituan:"invoice_support"`
	InvoiceMinPrice float64 `meituan:"invoice_min_price"`
	InvoiceDescription string `meituan:"invoice_description"`
	ThirdTagName string `meituan:"third_tag_name"`
	PreBook int `meituan:"pre_book"`
	TimeSelect int `meituan:"time_select"`
	AppBrandCode string `meituan:"app_brand_code"`
	PreBookMinDays int `meituan:"pre_book_min_days"`
	PreBookMaxDays int `meituan:"pre_book_max_days"`
}

func NewPoiSaveRequest(systemParam *domain.SystemParam) (poiSaveRequest PoiSaveRequest) {
	poiSaveRequest = PoiSaveRequest{
		BaseRequest: &BaseRequest{
			Uri:               "/api/v1/poi/save",
			RequestMethodType: domain.REQUEST_METHOD_TYPE_POST,
			SystemParam:       systemParam,
		},
	}
	return
}
