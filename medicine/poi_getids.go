package medicine

import (
	"encoding/json"
	"meituan-api/request"
	"meituan-api/response"
)

func (r *Client) PoiGetIds() (responseParams response.SystemResponse, err error) {
	poiGetIdsRequest := request.NewGetIdsRequest(systemParam)

	body, err := poiGetIdsRequest.DoRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &responseParams)
	if err != nil {
		return
	}

	return
}
