package service

import (
	"MessageBoard/model"
	dao2"MessageBoard/dao"
	"MessageBoard/tool"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type MessageService struct {

}

//获取一条信息
func (ms *MessageService)Getinfo(id string)(*model.Info,error){
	md:=dao2.Messagedao{tool.DbEngine}
	return md.GetInfo(id)
}

//输出套娃评论
func (ms *MessageService)TaowaComment(ctx *gin.Context,myinfo *model.Info)error{
	md:=dao2.Messagedao{tool.DbEngine}
	return md.TraverSon(ctx,myinfo)
}

//显示出一条留言
func (ms *MessageService)ShowOneMes(ctx *gin.Context,id string)error{
	md:=dao2.Messagedao{tool.DbEngine}
	return md.GetOneMes(ctx,id)
}

//显示下一层的所有留言
func (ms *MessageService)ShowAllMes(ctx *gin.Context)error{
	md:=dao2.Messagedao{tool.DbEngine}
	return md.ShowMes(ctx)
}

//删除留言
func (ms *MessageService)DeleteMes(id string)(sql.Result,error){
	md:=dao2.Messagedao{tool.DbEngine}
	return md.DeleteMes(id)
}

//发送留言
func (ms *MessageService)SendMes(name ,message string )error{
	md:=dao2.Messagedao{tool.DbEngine}
	return md.InsertMes(name,message)
}

//发送评论
func (ms *MessageService)SendComment(pid string,name string,message string)error{
	md:=dao2.Messagedao{tool.DbEngine}
	return md.InsertComment(pid,name,message)
}

//根据用户ID返回信息
func (ms *MessageService)QueryName(id string)string{
	md:=dao2.Messagedao{tool.DbEngine}
	return md.QueryName(id)
}