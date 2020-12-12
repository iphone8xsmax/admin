package models

import "time"

type BaseModel struct {
	ID	  int		`gorm:"size:10;primary_key;AUTO_INCREMENT" json:"id"`		//主键ID
	MTime time.Time `gorm:"column:mtime;default:null"          json:"mtime"`    //更新时间
	CTime time.Time `gorm:"column:ctime;default:null"          json:"ctime"`	//创建时间
}
