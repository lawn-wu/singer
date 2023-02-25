package response

type SystemResponse struct {
	Error struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}
	Data interface{} `json:"data"`
}

type PoiMGetResponse struct {
	Data []PoiMGetItemResponse `json:"data"`
}

type PoiMGetItemResponse struct {
	AppPoiCode         string  `json:"app_poi_code"`
	Name               string  `json:"name"`
	Address            string  `json:"address"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	PicUrl             string  `json:"pic_url"`
	PicUrlLarge        string  `json:"pic_url_large"`
	Phone              string  `json:"phone"`
	ShippingFee        float64 `json:"shipping_fee"`
	ShippingTime       string  `json:"shipping_time"`
	PromotionInfo      string  `json:"promotion_info"`
	InvoiceSupport     int     `json:"invoice_support"`
	InvoiceMinPrice    float64 `json:"invoice_min_price"`
	InvoiceDescription string  `json:"invoice_description"`
	OpenLevel          int     `json:"open_level"`
	IsOnline           int     `json:"is_online"`
	Ctime              int     `json:"ctime"`
	Utime              int     `json:"utime"`
	ThirdTagName       string  `json:"third_tag_name"`
	PreBook            int     `json:"pre_book"`
	TimeSelect         int     `json:"time_select"`
	LogisticsCodes     string  `json:"logistics_codes"`
	PreBookMinDays     int     `json:"pre_book_min_days"`
	PreBookMaxDays     int     `json:"pre_book_max_days"`
}
