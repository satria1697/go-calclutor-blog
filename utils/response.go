package utils

type Response struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}
