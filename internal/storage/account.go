package storage

import (
	"time"
)

type Account struct {
	ID       int       `gorm:"type:bigint(20) auto_increment;primary_key" json:"id"`
	Email    string    `gorm:"type:varchar(64);unique_index:email_idx" json:"email"`
	Phone    string    `gorm:"type:varchar(64);unique_index:phone_idx" json:"phone"`
	Name     string    `gorm:"type:varchar(32);unique_index:name_idx" json:"name"`
	Password string    `gorm:"type:varchar(32)" json:"password"`
	Birthday time.Time `gorm:"type:timestamp default '1970-01-02 00:00:00'" json:"birthday"`
	Gender   int       `gorm:"type:int(1)" json:"gender"`
	Avatar   string    `gorm:"type:varchar(512)" json:"avatar"`
}
