package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id       string
	email    string
	password string
}
