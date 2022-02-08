package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Status    int        `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"createAt" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updateAt" gorm:"column:updated_at;"`
}
