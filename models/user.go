package models

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name" binding:"required"`
	Email    string `gorm:"type:varchar(255);not null;uniqueIndex" json:"email" binding:"required"`
	Password string `gorm:"type:varchar(255);not null" json:"password,omitempty" binding:"required"`
	Address  string `gorm:"type:string;" json:"address"`
	Role     string `gorm:"type:varchar(255);NOT NULL" json:"role"`
}

func (u *User) BeforeSave(scope *gorm.DB) (err error) {
	u.Email = strings.ToLower(u.Email)
	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			return fmt.Errorf("failed to hash password")
		}
		u.Password = string(hash)
	}
	var find = scope.Where("email = ?", u.Email).First(&u)
	if find.RowsAffected > 0 {
		return fmt.Errorf("email already exists")
	}
	return
}

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) Mask() {
	u.Password = ""
}

func (u *User) GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  u.ID,
		"exp": time.Now().AddDate(0, 0, 7).Unix(),
	})
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
