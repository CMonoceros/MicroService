package response

import (
	"SnowBrick-Backend/common/errcode"
	"SnowBrick-Backend/common/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/pkg/errors"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(ctx *gin.Context, data interface{}, err error) {
	code := http.StatusOK
	realCode := errcode.Cause(err)

	ctx.Render(code, render.JSON{
		Data: Response{
			Code:    realCode.Code(),
			Message: realCode.Message(),
			Data:    data,
		},
	})
}

func CatchPanicFunc(ctx *gin.Context) func() {
	return func() {
		if r := recover(); r != nil {
			err := errors.WithStack(r.(error))
			log.Error("recover panic error(%+v)", err)

			JSON(ctx, nil, err)
		}
	}
}
