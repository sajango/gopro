package common

import "time"

type SQLModel struct {
	Id int `json:"id" gorm:"column:id;"`
	//FakeId    int        `json:"fake_id" gorm:"column:fake_id;"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"created_at"`
}
