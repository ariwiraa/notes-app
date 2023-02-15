package domain

import "time"

type Notes struct {
	ID        int32 `gorm:"primarykey auto_increment"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
