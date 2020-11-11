package response

// Response object
type Response interface {
	GetStatusCode() int
	GetMessage() string
	GetCode() string
	GetBody() ([]byte, error)
	GetData() interface{}
}

// Meta object
type Meta struct {
	Limit       int `json:"limit"`
	CurrentPage int `json:"current"`
}

// NewResponse - make a new response
func NewResponse(message string, statusCode int, code string, data interface{}) Response {

	switch {
	case statusCode < 300:
		return newSuccessResponse(message, statusCode, code, data)
	case statusCode < 500:
		return newClientErrorResponse(message, statusCode, code, data)
	default:
		return newServerErrorResponse(message, statusCode, code, data)
	}

}
