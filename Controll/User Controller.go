package Controll

import (
	"MessageBoard/service"
	"MessageBoard/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {

}

func (uc *UserController)Router(engine *gin.Engine){
	engine.POST("user/register",uc.Register)
	engine.POST("user/login",uc.Login)
	engine.POST("user/change_pw",uc.ChangePw)
	engine.GET("user/logout",uc.Logout)
}

//Register
func (uc *UserController) Register(context *gin.Context){
	name:=context.PostForm("name")
	password:=context.PostForm("password")
	fmt.Println("UserInfo: ",name,password)

	Us:=new(service.UserService)
	if Us.QueryUser(name){
		tool.PrintInfo(context,"该用户已经存在")
		return
	}

	if !Us.RegisterUser(name,password){
		tool.PrintInfo(context,"注册失败")
		return
	}
	tool.PrintInfo(context,"注册成功")
}

//Login
func (uc *UserController) Login(ctx *gin.Context){
	//检查是否已在线
	name:=ctx.PostForm("name")
	Us:=new(service.UserService)

	//检测用户信息是否正确
	password:=ctx.PostForm("password")
	fmt.Println("userinfo",name,password)


	if !Us.QueryUser(name){
		tool.PrintInfo(ctx,"用户名错误")
		return
	}

	cookie:=Us.Login(name,password)
	if cookie==nil{
		tool.PrintInfo(ctx,"密码错误")
		return
	}

	//先检查信息正确性再检查登录状态
	id:=strconv.Itoa(Us.QueryId(name))
	value:=tool.CheckUserOnline(ctx)
	if value==id{
		tool.PrintInfo(ctx,"用户已在线")
		return
	}

	//开始登录写入cookie
	http.SetCookie(ctx.Writer,cookie)
	tool.PrintInfo(ctx,name+"登录成功")
}

//ChangePw
func (uc *UserController) ChangePw(context *gin.Context){
	name:=context.PostForm("name")
	password:=context.PostForm("password")
	Us:=service.UserService{}

	//_,err:=Us.ChangePws(name,password)
	result,err:=Us.ChangePws(name,password)
	if err!=nil||result==nil{
		tool.PrintInfo(context,"请确认用户名是否正确")
		return
	}

	count,_:=result.RowsAffected()
	if count==0{
		tool.PrintInfo(context,"请提交修改后的密码")
		return
	}
	tool.PrintInfo(context,"修改密码成功")
}


//Logout
func (uc *UserController) Logout (ctx *gin.Context){
	//处在未登录状态
	if value:=tool.CheckUserOnline(ctx);value==""{
		tool.PrintInfo(ctx,"账号未登录")
		return
	}
	cookie,err:=ctx.Request.Cookie("login")
	if err!=nil{
		tool.PrintInfo(ctx,"获取cookie失败")
		return
	}
	//终止cookie运行
	cookie.MaxAge=-1
	http.SetCookie(ctx.Writer,cookie)
	tool.PrintInfo(ctx,"退出登录成功")
}

