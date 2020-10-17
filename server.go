package response

import "errors"

type serverErrorResponse struct {
	Message    string      `json:"message"`
	Code       string      `json:"code"`
	StatusCode int         `json:"-"`
	Errors     interface{} `json:"errors"`
}

// NewServerErrorResponse builds an instance of an error
func newServerErrorResponse(message string, statusCode int, code string, errors interface{}) Response {
	return &serverErrorResponse{
		Message:    message,
		Code:       code,
		StatusCode: statusCode,
		Errors:     errors,
	}
}

func (s *serverErrorResponse) GetStatusCode() int {
	return s.StatusCode
}

func (s *serverErrorResponse) GetMessage() string {
	return s.Message
}

func (s *serverErrorResponse) GetCode() string {
	return s.Code
}

func (s *serverErrorResponse) GetBody() ([]byte, error) {
	return nil, errors.New("Server Error does not body")
}
