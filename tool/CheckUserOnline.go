package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//检测用户是否已在线
func CheckUserOnline (ctx *gin.Context)string{
	cookie,err:=ctx.Request.Cookie("login")
	if err!=nil{
		fmt.Println("test: ", err)
		return ""
	}
	return cookie.Value
}