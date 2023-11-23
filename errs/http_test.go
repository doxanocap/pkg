package errs

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	err := NewHttp(http.StatusNotFound, "item not found")
	assert.Equal(t, "code: 404 | msg: item not found", err.Error())

	code := UnmarshalCode(err)
	assert.Equal(t, http.StatusNotFound, code)
	msg := UnmarshalMsg(err)
	assert.Equal(t, "item not found", msg)

	err = NewHttp(http.StatusBadRequest, "invalid request")
	assert.Equal(t, "code: 400 | msg: invalid request", err.Error())

	code = UnmarshalCode(err)
	assert.Equal(t, http.StatusBadRequest, code)
	msg = UnmarshalMsg(err)
	assert.Equal(t, "invalid request", msg)

	err = NewHttp(http.StatusUnauthorized, "user is not authorized to this request")
	assert.Equal(t, "code: 401 | msg: user is not authorized to this request", err.Error())

	code = UnmarshalCode(err)
	assert.Equal(t, http.StatusUnauthorized, code)
	msg = UnmarshalMsg(err)
	assert.Equal(t, "user is not authorized to this request", msg)

	err1 := New("code: 1 hello world")
	assert.NotEqual(t, "code: hello world", UnmarshalMsg(err1))
}
