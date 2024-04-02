package errs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	ErrUserNotFound = NewHttp(http.StatusNotFound, "user not found").
		AddTranslation("kz", "юзер табылмады").
		AddTranslation("ru", "юзер не найден")
)

func TestHttp(t *testing.T) {
	httpErr := NewHttp(http.StatusNotFound, "item not found")
	assert.Equal(t, "item not found", httpErr.Error())

	httpErr = UnmarshalError(ErrUserNotFound)

	assert.Equal(t, "user not found", httpErr.Error())

	// check translation
	assert.Equal(t, "юзер табылмады", httpErr.InLanguage("kz").Error())
	{
		err := someServiceCall()
		if err != nil {
			httpErr := UnmarshalError(err)

			switch httpErr.StatusCode {
			case http.StatusInternalServerError:
				return
			default:
				fmt.Println(httpErr.InLanguage("kz"))
			}
		}
	}

}

func someServiceCall() error {
	return ErrUserNotFound
}
