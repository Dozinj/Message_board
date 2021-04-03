package main

import (
	"MessageBoard/Controll"
	"MessageBoard/tool"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	//解析文件
	cfg, err := tool.ParseConfig("./Config/app.Json")
	if err != nil {
		panic(err)
		return
	}

	_, err = tool.OrmEngine(cfg)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	register(engine)
	engine.Run()

}

func register(engine  *gin.Engine){
	new(Controll.UserController).Router(engine)
	new(Controll.MessageController).Router(engine)
}