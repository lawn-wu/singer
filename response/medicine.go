package response

type SystemResponse struct {
	Error struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}
	Data interface{} `json:"data"`
}
