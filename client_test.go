package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientFunction(t *testing.T) {

	res := newClientErrorResponse("error example", 400, "test", nil)
	body, err := res.GetBody()

	assert.EqualValues(t, res.GetCode(), "test")
	assert.EqualValues(t, res.GetStatusCode(), 400)
	assert.EqualValues(t, res.GetMessage(), "error example")
	assert.Nil(t, body)
	assert.NotNil(t, err)
}
