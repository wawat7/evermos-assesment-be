package helper

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
