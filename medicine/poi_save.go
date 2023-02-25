package medicine

import (
	"encoding/json"
	"meituan-api/request"
	"meituan-api/response"
)

func (r *Client) PoiSave(params request.PoiSaveParam) (responseParams response.SystemResponse, err error) {
	poiSaveRequest := request.NewPoiSaveRequest(systemParam)
	poiSaveRequest.PoiSaveParam = params
	poiSaveRequest.SetApplicationParams(poiSaveRequest.PoiSaveParam)

	body, err := poiSaveRequest.DoRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &responseParams)
	if err != nil {
		return
	}

	return
}
