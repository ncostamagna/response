package response

import "errors"

type clientErrorResponse struct {
	Message    string      `json:"message"`
	Code       string      `json:"code"`
	StatusCode int         `json:"-"`
	Errors     interface{} `json:"errors"`
}

// NewClientErrorResponse builds an instance of an error
func newClientErrorResponse(message string, statusCode int, code string, errors interface{}) Response {
	return &clientErrorResponse{
		Message:    message,
		Code:       code,
		StatusCode: statusCode,
		Errors:     errors,
	}
}

func (s *clientErrorResponse) GetStatusCode() int {
	return s.StatusCode
}

func (s *clientErrorResponse) GetMessage() string {
	return s.Message
}

func (s *clientErrorResponse) GetCode() string {
	return s.Code
}

func (s *clientErrorResponse) GetBody() ([]byte, error) {
	return nil, errors.New("Client Error does not body")
}
