package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"statuss"`
	Data   interface{} `json:"data"`
}
