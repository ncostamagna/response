package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseSuccessFunction(t *testing.T) {
	data := struct {
		TextValue   string `json:"textValue"`
		NumberValue int    `json:"numberValue"`
	}{"test", 5}
	valueToCompare := `{"message":"error example","code":"test","data":{"textValue":"test","numberValue":5},"meta":{"limit":0,"current":0}}`

	res := NewResponse("error example", 200, "test", data)

	body, err := res.GetBody()

	assert.EqualValues(t, res.GetCode(), "test")
	assert.EqualValues(t, string(body), valueToCompare)
	assert.EqualValues(t, res.GetStatusCode(), 200)
	assert.EqualValues(t, res.GetMessage(), "error example")
	assert.NotNil(t, body)
	assert.Nil(t, err)
}

func TestResponseClientFunction(t *testing.T) {

	res := NewResponse("error example", 400, "test", nil)
	body, err := res.GetBody()

	assert.EqualValues(t, res.GetCode(), "test")
	assert.EqualValues(t, res.GetStatusCode(), 400)
	assert.EqualValues(t, res.GetMessage(), "error example")
	assert.Nil(t, body)
	assert.NotNil(t, err)
}

func TestResponseServerFunction(t *testing.T) {
	res := NewResponse("error example", 500, "test", nil)
	body, err := res.GetBody()

	assert.EqualValues(t, res.GetCode(), "test")
	assert.EqualValues(t, res.GetStatusCode(), 500)
	assert.EqualValues(t, res.GetMessage(), "error example")
	assert.NotNil(t, err)
	assert.Nil(t, body)
}
