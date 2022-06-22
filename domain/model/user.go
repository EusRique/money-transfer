package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID        uint       `gorm:"primaryKey;autoIncrement:true" json:"id" valid:"-"`
	Name      string     `gorm:"type:varchar(56)" json:"name" valid:"notnull"`
	LastName  string     `gorm:"type:varchar(56)" json:"last_name" valid:"notnull"`
	Document  string     `gorm:"type:varchar(56)" json:"document" valid:"notnull"`
	Email     string     `gorm:"type:varchar(56)" json:"email" valid:"notnull"`
	Password  string     `gorm:"type:varchar(56)" json:"password" valid:"notnull"`
	IsSeller  bool       `gorm:"type:bool" json:"is_seller" valid:"-"`
	CreatedAt time.Time  `gorm:"type:timestamp;autoCreateTime;default:CURRENT_TIMESTAMP" json:"created_at" valid:"-"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime" json:"updated_at" valid:"-"`
}

func (user *User) IsValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	return nil
}
