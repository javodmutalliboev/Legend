package response

type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   string `json:"data"`
}
