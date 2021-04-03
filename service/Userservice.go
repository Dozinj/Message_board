package service

import (
	dao2 "MessageBoard/dao"
	"MessageBoard/model"
	"MessageBoard/tool"
	"database/sql"
	"net/http"
	"strconv"
	"time"
)

type UserService struct {

}

//检查用户是否存在
func (us UserService)QueryUser(name string) bool{
	Ud:=dao2.Userdao{tool.DbEngine}
	flag:=Ud.QueryName(name)
	if !flag{
		return false
	}
	return true
}

//插入注册信息----返回ture注册成功
func (us UserService) RegisterUser(name,pwd string)bool{
	userinfo:=model.User{
		Name: name,
		Password: pwd,
		RegisterTime: time.Now().Unix(),
	}
	Ud:=dao2.Userdao{tool.DbEngine}
	//将注册信息插入数据库
	result:=Ud.InsertInfo(userinfo)
	return result>0
}

//用户登录
func (us UserService)Login(name,password string)*http.Cookie{
	Ud:=dao2.Userdao{tool.DbEngine}
	Pwd:=Ud.QueryPassword(name)
	if password!=Pwd{
		return nil
	}

	cookie:=&http.Cookie{
		Name: "login",
		Value:strconv.Itoa(Ud.QueryId(name)),
		MaxAge: 300,
		Path: "/",
		HttpOnly: true,
		Secure: false,
	}
	return cookie
}

//修改用户密码
func (us UserService)ChangePws(name,pwd string)(sql.Result,error) {
	Ud := dao2.Userdao{tool.DbEngine}
	return Ud.ChangePwd(name,pwd)
}

//根据用户名返回id
func (us UserService)QueryId(name string)int{
	ud:=dao2.Userdao{tool.DbEngine}
	return ud.QueryId(name)
}



