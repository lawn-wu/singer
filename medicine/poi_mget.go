package medicine

import (
	"encoding/json"
	"meituan-api/request"
	"meituan-api/response"
)

func (r *Client) PoiMGet(params request.PoiMGetParam) (responseParams response.PoiMGetResponse, err error) {
	poiMGetRequest := request.NewPoiMGetRequest(systemParam)
	poiMGetRequest.PoiMGetParam = params
	poiMGetRequest.SetApplicationParams(poiMGetRequest.PoiMGetParam)

	body, err := poiMGetRequest.DoRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &responseParams)
	if err != nil {
		return
	}

	return
}
