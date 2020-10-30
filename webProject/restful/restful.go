package restful

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Fail(err error, c *gin.Context) {
	var (
		apiErr *ApiCode
		errno    = UnknownErr
	)
	if errors.As(err, &apiErr) {
		errno = apiErr.Code
	}
	RenderJson(errno, nil, err.Error(), c)
}

func RenderJson(errno int, data interface{}, error string, c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		errno,
		data,
		error,
	})
}

func NewApiCode(code int, msg string) *ApiCode {
	return  &ApiCode {
		Code: code,
		Msg: msg,
	}
}

func WithDataset(data interface{}, err error, c *gin.Context) {
	if err == nil {
		WithData(data, c)
	} else {
		Fail(err, c)
	}
}

func WithData(data interface{}, c *gin.Context) {
	RenderJson(SUCCESS, data, "", c)
}