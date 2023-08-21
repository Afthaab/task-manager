package domain

import (
	"time"
)

type Task struct {
	Taskid      uint      `json:"Taskid" gorm:"primaryKey;autoIncrement:true;unique"`
	User        User      `gorm:"ForeignKey:uid"`
	Uid         uint      `json:"uid"`
	Taskname    string    `json:"taskname"`
	Description string    `json:"description"`
	Duedate     string    `json:"duedate"`
	Duetime     string    `json:"duetime"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
