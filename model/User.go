package model

type User struct {
	Id int `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"varchar(20)" json:"name"`
	Password string `xorm:"varchar(15)" json:"password"`
	RegisterTime int64`xorm:"bigint" json:"register_time"`
}

