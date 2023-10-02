package entity

import (
	"github.com/Anupam-dagar/baileys/entity"
	"github.com/oleiade/reflections"
)

type Poll struct {
	entity.BaseModel
	Title string `gorm:"column:title" json:"title"`
}

type PollPtr = *Poll

func (p *Poll) GetModel() interface{} {
	return &Poll{}
}

func (p *Poll) SetCol(field string, val interface{}) error {
	return reflections.SetField(p, field, val)
}
