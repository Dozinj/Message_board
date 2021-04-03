package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type DatabaseConfig struct {
	AppName string `json:"app_name"`
	Driver string `json:"driver"`
	User string `json:"user"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
	DbName string `json:"db_name"`
	CharSet string `json:"charset"`
	ShowSql bool `json:"show_sql"`

}

var _cfg *DatabaseConfig

func GetConfig()*DatabaseConfig{
	return _cfg
}

//解析Config里的文件
func ParseConfig (path string)(*DatabaseConfig ,error){
	file,err:=os.Open(path)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	reader:=bufio.NewReader(file)
	decode:=json.NewDecoder(reader)

	if err=decode.Decode(&_cfg);err!=nil{
		return nil,err
	}

	return _cfg,nil

}
