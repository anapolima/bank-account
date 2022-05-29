package models

type Response struct {
	Data     interface{} `json:"data"`
	Messages []string    `json:"messages"`
}

type ResponseInfo struct {
	StatusCode int      `json:"statusCode"`
	Data       Response `json:"data"`
}
