package errs

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

var (
	ErrUserNotFound = NewHttp(http.StatusNotFound, "юзер не найден").
		SetErrorCode("auth.error.user_not_found")
)

func TestHttp(t *testing.T) {

	// check translation
	{
		ctx := &gin.Context{}

		err := func() error {
			return ErrUserNotFound
		}()
		if err != nil {
			SetGinError(ctx, ErrUserNotFound)
			return
		}
	}

	ctx := &gin.Context{}

	SetGinError(ctx, ErrUserNotFound)

	raw, _ := json.Marshal(ErrUserNotFound)
	fmt.Println(string(raw))

	fmt.Println(ErrUserNotFound.Error())
}
