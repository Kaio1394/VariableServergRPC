package model

import "time"

type Variable struct {
	ID       uint      `gorm:"primary_key"`
	Key      string    `gorm:"type:varchar(255);unique;not null)"`
	Value    string    `gorm:"type:text;"`
	EditDate time.Time `gorm:"type:timestamp"`
}

func (v Variable) TableName() string {
	return "t_variables"
}
