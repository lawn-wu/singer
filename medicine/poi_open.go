package medicine

import (
	"encoding/json"
	"meituan-api/request"
	"meituan-api/response"
)

func (r *Client) PoiOpen() (responseParams response.SystemResponse, err error) {
	poiOpenRequest := request.NewPoiOpenRequest(systemParam)

	body, err := poiOpenRequest.DoRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &responseParams)
	if err != nil {
		return
	}

	return
}
