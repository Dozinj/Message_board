package model

type MessageInfo struct {
	Id int  `xorm:"pk autoincr" json:"id"`
	Pid int  `xorm:"int" json:"pid"`
	Name string `xorm:"varchar(20)" json:"name"`
	Message string `xorm:"varchar(100)" json:"password"`
	CommentNum int `xorm:"bigint" json:"comment_num"`
	Time int64 `xorm:"bigint" json:"time"`
}
