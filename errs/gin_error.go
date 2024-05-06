package errs

import (
	"github.com/gin-gonic/gin"
)

// SetGinError sets error in gin.Context
// and if it is HttpError, it sets right headers and writes message into body
// Best use if you are not sure which type of error you are handling
func SetGinError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	// this error will be parsed and logged
	_ = ctx.Error(err)

	// if error is HttpError
	httpError := UnmarshalError(err)
	if httpError.StatusCode != 0 {
		ctx.JSON(httpError.StatusCode, httpError)
		return
	}
	return
}

func SetBothErrors(ctx *gin.Context, publicErr *HttpError, privateErr error) {
	if publicErr == nil {
		return
	}

	_ = ctx.Error(privateErr)
	ctx.JSON(publicErr.StatusCode, publicErr)
	_ = ctx.Error(publicErr)
	return
}

func GetGinPrivateErr(ctx *gin.Context) error {
	if len(ctx.Errors) == 0 {
		return nil
	}
	return ctx.Errors[0]
}

func GetGinPublicErr(ctx *gin.Context) *HttpError {
	if len(ctx.Errors) == 0 {
		return nil
	} else if len(ctx.Errors) == 1 {
		httpError := UnmarshalError(ctx.Errors[1])
		if httpError.StatusCode != 0 {
			return httpError
		}
		return nil
	}

	httpError := UnmarshalError(ctx.Errors[1])
	if httpError.StatusCode == 0 {
		return nil
	}
	return httpError
}
