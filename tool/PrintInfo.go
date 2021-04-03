package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PrintInfo(ctx *gin.Context,v interface{}){
	ctx.JSON(http.StatusOK,map[string]interface{}{
		"data":v,
	})
}

