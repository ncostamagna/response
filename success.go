package response

import "encoding/json"

type successResponse struct {
	Message    string      `json:"message"`
	Code       string      `json:"code"`
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data"`
	Meta       Meta        `json:"meta"`
}

func newSuccessResponse(message string, statusCode int, code string, data interface{}) Response {
	return &successResponse{
		Message:    message,
		Code:       code,
		StatusCode: statusCode,
		Data:       data,
	}
}

func (s *successResponse) GetStatusCode() int {
	return s.StatusCode
}

func (s *successResponse) GetMessage() string {
	return s.Message
}

func (s *successResponse) GetCode() string {
	return s.Code
}

func (s *successResponse) GetBody() ([]byte, error) {
	return json.Marshal(s)
}
