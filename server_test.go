package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerFunction(t *testing.T) {
	res := newServerErrorResponse("error example", 500, "test", nil)
	body, err := res.GetBody()

	assert.EqualValues(t, res.GetCode(), "test")
	assert.EqualValues(t, res.GetStatusCode(), 500)
	assert.EqualValues(t, res.GetMessage(), "error example")
	assert.NotNil(t, err)
	assert.Nil(t, body)
}
