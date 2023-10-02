package entity

import (
	"github.com/Anupam-dagar/baileys/entity"
	"github.com/oleiade/reflections"
)

type PollOption struct {
	entity.BaseModel
	PollId string `gorm:"column:poll_id" json:"pollId"`
	Title  string `gorm:"column:title" json:"title"`
}

type PollOptionPtr = *PollOption

func (p *PollOption) GetModel() interface{} {
	return &PollOption{}
}

func (p *PollOption) SetCol(field string, val interface{}) error {
	return reflections.SetField(p, field, val)
}
