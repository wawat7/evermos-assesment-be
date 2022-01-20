package helper

import "encoding/json"

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {
	return Response{
		Data: data,
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  status,
		},
	}
}

func ConvertDataToJsonString(data interface{}) string {
	jsonByte, err := json.Marshal(data)
	PanicIfError(err)
	return string(jsonByte)
}
