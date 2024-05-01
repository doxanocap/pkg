package errs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetGinError handles privateError and optional publicError
// privateErr might be also an HttpError, however if not it just puts error into gin.Ctx
// Better use to write publicErr if known what privateError is not HttpError type
func SetGinError(ctx *gin.Context, privateErr error, publicErr ...error) {
	if privateErr == nil {
		return
	}

	// this error will be parsed and logged
	_ = ctx.Error(privateErr)

	// if private error is also HttpError
	httpError := UnmarshalError(privateErr)
	if httpError.StatusCode != 0 {
		ctx.JSON(httpError.StatusCode, httpError)
		return
	}

	if publicErr == nil {
		return
	}

	// this error will be shown as output
	httpError = UnmarshalError(publicErr[0])
	if httpError.StatusCode != 0 {
		ctx.JSON(httpError.StatusCode, httpError)
	} else {
		ctx.Status(http.StatusInternalServerError)
	}
	return
}

func GetGinPrivateErr(ctx *gin.Context) error {
	if ctx.Errors == nil {
		return nil
	}
	return ctx.Errors[0]
}

func GetGinPublicErr(ctx *gin.Context) error {
	if len(ctx.Errors) < 2 {
		return nil
	}
	return ctx.Errors[1]
}

func SetGinErrorWithStatus(ctx *gin.Context, status int, err error) {
	if err == nil {
		return
	}
	_ = ctx.Error(err)
	ctx.Status(status)
	return
}
