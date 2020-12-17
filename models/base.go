package models

import (
	"time"
)

type BaseModel struct {
	ID	  	int			`gorm:"size:10;primary_key;AUTO_INCREMENT" json:"id"`		//主键ID
	Status  int 		`gorm:"size:4;default:1"                json:"status"`   //状态{-1：删除；1：正常；}
	MTime 	time.Time 	`gorm:"column:mtime;default:null"          json:"mtime"`    //更新时间
	CTime 	time.Time 	`gorm:"column:ctime;default:null"          json:"ctime"`	//创建时间
}
//验证status是否有效
func (b BaseModel) IsValidStatus() bool{
	return b.Status == 1  || b.Status == -1
}

