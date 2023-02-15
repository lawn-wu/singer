package medicine

import "meituan-api/domain"

type Client struct {
	domain.Config
}

var systemParam = &domain.SystemParam{}

func NewClient(config domain.Config) *Client {
	systemParam = &domain.SystemParam{
		AppId:             config.AppId,
		AppSecret:         config.AppSecret,
		AppPoiAccessToken: config.AppPoiAccessToken,
		AppPoiCode:        config.AppPoiCode,
	}
	return &Client{config}
}
