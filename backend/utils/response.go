package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/hongminhcbg/weekend-contest-smv/backend/erp"
	"net/http"
)

func ResponseError(ctx *gin.Context, err *erp.RuntimeError) {
	ctx.JSON(err.Code/1000,
		gin.H{
			"meta": gin.H{
				"code":    err.Code,
				"message": err.Message,
			},
		})
}

func ResponseData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":    http.StatusOK,
			"message": "success",
		},
		"data": data,
	})
}
