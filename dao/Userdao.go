package dao

import (
	"MessageBoard/model"
	"MessageBoard/tool"
	"database/sql"
	"fmt"
	"log"
)

type Userdao struct {
	*tool.Orm
}

//修改用户密码
func (ud *Userdao)ChangePwd(name,pwd string)(sql.Result,error){
	sql:="update user set password=? where name=?"
	result,err:=ud.Exec(sql,pwd,name)
	//修改密码时需检查用户名正确性
	if err!=nil||!ud.QueryName(name){
		fmt.Println(err)
		return nil, err
	}
	return result,nil
}

//检查用户名是否存在
func (ud *Userdao)QueryName(name string)bool {
	student := new(model.User)
	flag,err:=ud.SQL("select * from user where name=?",name).Get(student)
	if err!=nil{
		log.Fatal(err)
	}
	return flag
}

//根据用户名返回用户密码
func (ud *Userdao)QueryPassword(name string)string{
	var student model.User
	flag,err:=ud.SQL("select * from user where name=?",name).Get(&student)
	if err!=nil{
		log.Fatal(err)
	}
	if !flag{
		return ""
	}
	return student.Password
}

//根据用户名返回用户id
func (ud *Userdao)QueryId(name string)int{
	user:=new(model.User)
	flag,err:=ud.SQL("select id from user where name=?",name).Get(user)
	if err!=nil{
		log.Fatal(err)
	}
	if !flag{
		return -1
	}
	return user.Id
}

//插入注册信息
func (ud *Userdao)InsertInfo(user model.User)int64{
	result,err:=ud.InsertOne(&user)
	if err!=nil{
		log.Fatal(err)
	}
	return result
}




