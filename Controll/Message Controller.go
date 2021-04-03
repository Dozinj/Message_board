package Controll

import (
	"MessageBoard/service"
	"MessageBoard/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type MessageController struct {

}
//资源路径
func (mc *MessageController)Router(engine *gin.Engine){
	engine.GET("/mes/:id/getone",mc.GetOneMes)
	engine.GET("/mes/:id/getmore",mc.ShowComments)
	engine.GET("/mes/:id/getall",mc.ShowAllMes)
	engine.DELETE("mes/:id/delete",mc.DeleteOneComment)
	engine.POST("/mes/anonymous",mc.AnonymousMes)
	engine.POST("/mes/sendmes",mc.SendMes)
	engine.POST("/mes/sendcomment",mc.SendComment)
}

//匿名留言
func (mc *MessageController)AnonymousMes(ctx *gin.Context){
	//检查登录状态
	value:=tool.CheckUserOnline(ctx)
	if value==""{
		tool.PrintInfo(ctx,"用户未登录")
		return
	}
	//获取评论信息
	message:=ctx.PostForm("message")
	ms:=service.MessageService{}
	err:=ms.SendMes("anonymity",message)
	if err!=nil{
		log.Fatal(err)
		return
	}
	tool.PrintInfo(ctx,"匿名留言成功")
}

//获取指定一条信息
func (mc *MessageController)GetOneMes(ctx *gin.Context){
	id:=ctx.Param("id")
	ms:=service.MessageService{}
	err:=ms.ShowOneMes(ctx,id)
	if err!=nil{
		log.Fatal(err)
	}
}

//获取一条留言及其下的全部评论信息
func (mc *MessageController)ShowComments(ctx *gin.Context){
	//获取Pid
	pid:=ctx.Param("id")
	ms:=service.MessageService{}

	err:=ms.ShowOneMes(ctx,pid)
	if err!=nil{
		log.Fatal(err)
		return
	}

	info,err:=ms.Getinfo(pid)
	if err!=nil{
		log.Fatal(err)
		return
	}
	//如果info为nil的话直接退出
	if info==nil{
		return
	}

	fmt.Println(info)
	err=ms.TaowaComment(ctx,info)
	if err!=nil{
		log.Fatal(err)
		return
	}
}

//对一个留言发送一条评论
func (mc *MessageController)SendComment(ctx *gin.Context){
	//检查登录状态
	valueId:=tool.CheckUserOnline(ctx)
	if valueId==""{
		tool.PrintInfo(ctx,"用户未登录")
		return
	}

	pid:=ctx.PostForm("id")
	message:=ctx.PostForm("message")
	ms:=service.MessageService{}
	username:=ms.QueryName(valueId)

	err:=ms.SendComment(pid,username,message)
	if err!=nil{
		log.Fatal(err)
		return
	}
	tool.PrintInfo(ctx, "评论成功")
}

//找出所有留言信息
func (mc *MessageController)ShowAllMes(ctx *gin.Context){
	ms:=service.MessageService{}
	err:=ms.ShowAllMes(ctx)
	if err!=nil{
		log.Fatal(err)
		return
	}
}

//删除一条留言
func (mc *MessageController)DeleteOneComment(ctx *gin.Context) {
	id := ctx.Param("id")
	ms := service.MessageService{}
	result,err := ms.DeleteMes(id)
	if err != nil {
		log.Fatal(err)
		return
	}
	rowAff,_:=result.RowsAffected()
	if rowAff==0{
		tool.PrintInfo(ctx,"评论删除失败")
		return
	}
	tool.PrintInfo(ctx, "评论删除成功")
}

//发送一条留言
func (mc *MessageController)SendMes(ctx *gin.Context){
	valueId:=tool.CheckUserOnline(ctx)
	if valueId==""{
		tool.PrintInfo(ctx,"用户未登录")
		return
	}

	message:=ctx.PostForm("message")
	ms:=service.MessageService{}
	username:=ms.QueryName(valueId)

	err:=ms.SendMes(username,message)
	if err!=nil{
		log.Fatal(err)
		return
	}
	tool.PrintInfo(ctx,"实名留言成功")
}
