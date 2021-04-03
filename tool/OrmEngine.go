package tool

import (
	"MessageBoard/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

//建立了链接后，还要在目录下再建立数据库
func  OrmEngine(cfg *DatabaseConfig)(*Orm ,error){
	database:=cfg
	//连接字符串
	conn:=database.User+":"+database.Password+"@tcp("+database.Host+":"+database.Port+")/"+database.DbName+"?charset"+database.CharSet
	//conn="root:123456@tcp(127.0.0.1:3306)/message_board?charset=utf8"
	engine,err:=xorm.NewEngine(database.Driver,conn)
	if err!=nil{
		return nil, err
	}

	engine.ShowSQL(database.ShowSql)

	//创建数据表
	err=engine.Sync(new(model.User),
		new(model.MessageInfo))

	if err!=nil{
		return nil,err
	}

	orm:=new(Orm)
	orm.Engine=engine

	//数据库表引擎
	DbEngine=orm
	return orm,nil
}

