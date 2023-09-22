package errs

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	err := NewHttp(http.StatusNotFound, "item not found")

	assert.Equal(t, "code: 404 | msg: item not found", err.Error())

	n := UnmarshalCode(err)
	assert.Equal(t, 404, n)
}
