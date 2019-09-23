package dtos

type Response struct {
	Success bool        `json:"success"`
	Message string      `josn:"message"`
	Data    interface{} `json:"data"`
}
