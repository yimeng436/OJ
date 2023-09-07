package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode = 0
	Error       = 1
)

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Total int         `json:"total"`
}

func SuccessWithPage(ctx *gin.Context, data interface{}, msg string, total int) {
	ctx.JSON(http.StatusOK, Response{
		Code:  SuccessCode,
		Data:  data,
		Msg:   msg,
		Total: total,
	})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code: SuccessCode,
		Data: data,
		Msg:  msg,
	})
}
func Fail(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code: Error,
		Msg:  msg,
	})
}
