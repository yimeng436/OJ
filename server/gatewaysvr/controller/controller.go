package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yimeng436/OJ/common"
)

// @Summary		测试接口
// @Description	描述信息
// @Param			id	path		int		true	"Account ID"
// @Produce  json
// @Success		200	{object}	constant.Response
// @Router			/test/{id} [get]
func Test(ctx *gin.Context) {
	ctx.JSON(200, common.Response{
		Code: 200,
		Data: "adsda",
		Msg:  "success",
	})
}
