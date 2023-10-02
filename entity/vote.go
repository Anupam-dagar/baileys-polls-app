package entity

import (
	"github.com/Anupam-dagar/baileys/entity"
	"github.com/oleiade/reflections"
)

type Vote struct {
	entity.BaseModel
	PollId       string `gorm:"column:poll_id" json:"pollId"`
	PollOptionId string `gorm:"column:poll_option_id" json:"pollOptionId"`
}

type VotePtr = *Vote

func (v *Vote) GetModel() interface{} {
	return &Vote{}
}

func (v *Vote) SetCol(field string, val interface{}) error {
	return reflections.SetField(v, field, val)
}
