package dao

import (
	"MessageBoard/model"
	"MessageBoard/tool"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Messagedao struct {
  	*tool.Orm
}

//打印出给自己留言的所有信息
func (md *Messagedao)TraverSon(ctx *gin.Context,Myinfo *model.Info)error{
	//找到目标结点
	mesInfo:=new(model.MessageInfo)
	rows,err:=md.SQL("select * from message_info where pid =?",Myinfo.Id).Rows(mesInfo)
	if err!=nil{
		fmt.Println("rowserr:",err)
		return err
	}

	//找到自己的所有子节点
	defer rows.Close()
	for rows.Next(){
		sonMessage:=new(model.MessageInfo)
		if err=rows.Scan(sonMessage);err!=nil{
			fmt.Println("err:",err)
			return err
		}
		fmt.Println(sonMessage) //test

		sonInfo:=model.Info{
			Id: sonMessage.Id,
			Username: sonMessage.Name,
			Message: sonMessage.Message,
		}

		//如果遍历到了自己的评论则跳过
		if sonInfo.Username==Myinfo.Username{
			continue
		}

		//打印出遍历的子节点信息
		ctx.JSON(http.StatusOK,gin.H{
			"originName":Myinfo.Username,
			"originMes":Myinfo.Message,
			"commentId":sonMessage.Id,
			"commentName":sonMessage.Name,
			"commentMes":sonMessage.Message,
			"commentNum":sonMessage.CommentNum,
			"commentTime":sonMessage.Time,
		})
		if err=rows.Err();err!=nil {
			return err
		}
		//递归调用遍历
		err=md.TraverSon(ctx,&sonInfo)
		if err!=nil{
			log.Fatal(err)
			return err
		}
	}
	return nil
}

//根据Id 返回一个结点信息--info struct
func (md *Messagedao)GetInfo(id string)(*model.Info, error){
	mesInfo:=new(model.MessageInfo)
	flag,err:=md.SQL("select name,message from message_info where id=?",id).Get(mesInfo)
	if err!=nil{
		log.Fatal(err)
		return nil, err
	}
	if flag {
		info := new(model.Info)
		//类型转换
		info.Id,_= strconv.Atoi(id)
		info.Username = mesInfo.Name
		info.Message = mesInfo.Message
		return info,nil
	}
	return nil, err
}

//根据Id打印出一条评论信息
func (md *Messagedao)GetOneMes(ctx *gin.Context,id string)error{
	mesInfo:=new(model.MessageInfo)
	flag,err:=md.SQL("select * from message_info where id=?",id).Get(mesInfo)
	if err!=nil{
		log.Fatal(err)
		return err
	}

	if flag {
		ctx.JSON(http.StatusOK, gin.H{
			"id":          id,
			"name":        mesInfo.Name,
			"message":     mesInfo.Message,
			"time":        mesInfo.Time,
			"comment_num": mesInfo.CommentNum,
		})
		return nil
	}
	tool.PrintInfo(ctx,"用户ID未找到")
	return nil
}

//找出所有评论信息
func (md *Messagedao)ShowMes(ctx *gin.Context)error{
	mesInfo:=new(model.MessageInfo)
	//找到留言者信息
	rows,err:=md.SQL("select * from message_info where id=pid").Rows(mesInfo)
	if err!=nil{
		log.Fatal(err)
		return err
	}
	defer rows.Close()
	for rows.Next(){
		sonMes:=new(model.MessageInfo)

		err=rows.Scan(sonMes)
		if err!=nil{
			log.Fatal(err)
			return err
		}
		ctx.JSON(http.StatusOK,gin.H{
			"id":sonMes.Id,
			"name":sonMes.Name,
			"message":sonMes.Message,
			"time":sonMes.Time,
			"commentnum":sonMes.CommentNum,
		})

		if err=rows.Err();err!=nil {
			return err
		}
	}
	return nil
}

//根据Id删除数据表中一条信息
func (md *Messagedao)DeleteMes(id string)(sql.Result,error){
	sql:="delete from message_info where id=?"
	result,err:=md.Exec(sql,id)
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	return result,nil
}

//向数据库中插入一条留言
func (md *Messagedao)InsertMes(name,message string)error{
	insertTime:=time.Now().Unix()
	mesInfo:=model.MessageInfo{
		Name: name,
		Message: message,
		Time: insertTime,
		CommentNum: 0,
	}
	_,err:=md.InsertOne(mesInfo)
	if err!=nil{
		log.Fatal(err)
		return err
	}

	//获取刚刚插入留言的节点
	sonInfo:=new(model.MessageInfo)
	_,err=md.SQL("select id from message_info where name=?",name).Get(sonInfo)
	if err!=nil{
		log.Fatal(err)
		return err
	}
	var id int
	id=sonInfo.Id

	//将新插入信息设置为一个父节点
	sql:="update message_info set pid =? where id=?"
	_,err=md.Exec(sql,id,id)
	if err!=nil{
		log.Fatal(err)
		return err
	}
	return nil
}

//插入一条评论信息----传入pid在指定人下评论
func (md *Messagedao)InsertComment(pid string,name string,message string)error{
	//存入数据库
	id,_:=strconv.Atoi(pid)

	sonInfo:=model.MessageInfo{
		Pid: id,
		Name: name,
		Message: message,
		Time: time.Now().Unix(),
		CommentNum: 0,
	}
	_,err:=md.InsertOne(sonInfo)
	if err!=nil{
		log.Fatal(err)
		return err
	}

	//修改评论数量
	mesInfo:=new(model.MessageInfo)
	_,err=md.SQL("select comment_num from message_info where id=?",pid).Get(mesInfo)
	if err!=nil{
		log.Fatal(err)
		return err
	}
	commentnum:=mesInfo.CommentNum
	commentnum++

	sql:="update message_info set comment_num=? where id=?"
	_,err=md.Exec(sql,commentnum,pid)
	if err!=nil{
		log.Fatal(err)
		return err
	}
	return nil
}

//根据用户ID返回username
func (md *Messagedao)QueryName(ID string)string{
	id,_:=strconv.Atoi(ID)
	mesInfo:=new(model.MessageInfo)
	//获取真实姓名需从user表中获取
	_,err:=md.SQL("select name from user where id=?",id).Get(mesInfo)
	if err!=nil{
		log.Fatal(err)
	}
	return mesInfo.Name
}
