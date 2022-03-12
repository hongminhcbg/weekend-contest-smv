package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/hongminhcbg/weekend-contest-smv/backend/erp"
	"net/http"
	"strconv"
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

func GetQueryInt(ctx *gin.Context, q string, defaultVal int) int {
	qValue, ok := ctx.GetQuery(q)
	if !ok || len(qValue) == 0 {
		return defaultVal
	}

	qValueInt, err := strconv.Atoi(qValue)
	if err != nil {
		return defaultVal
	}

	return  qValueInt
}