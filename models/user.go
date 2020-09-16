package models

import (
	"fmt"
	"time"

	"github.com/ezedinff/appzoid/database"
	"github.com/ezedinff/appzoid/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User user model
type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *User) create() map[string]interface{} {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	database.GetDB().Create(user)
	if user.ID <= 0 {
		fmt.Println("Failed to create the user")
	}
	response := utils.Message(true, "")
	response["user"] = user
	return response
}

// GetUsers get all users
func GetUsers() []*User {
	users := make([]*User, 0)
	result := database.GetDB().Find(&users)
	if result.Error != nil {
		fmt.Println("Error!")
	}
	return users
}
