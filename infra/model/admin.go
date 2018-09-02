package model

import (
	"FIFA-World-Cup-Backstage-Management-System/infra/crypto"
	"FIFA-World-Cup-Backstage-Management-System/infra/init"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/swaggo/swag/example/celler/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Admin struct {
	gorm.Model
	Name              string `gorm:"type:varchar;column:name"`
	AuthToken         string `gorm:"type:varchar;column:auth_token"`
	EncryptedPassword string `gorm:"type:varchar;column:encrypted_password"`
	Phone             string `gorm:"type:varchar;column:phone"`
	State             string `gorm:"type:varchar;column:state"`
}

type AdminSerializer struct {
	ID                uint       `json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	Name              string     `json:"name"`
	AuthToken         string     `json:"auth_token"`
	EncryptedPassword string     `json:"encrypted_password"`
	Phone             string     `json:"phone"`
	State             string     `json:"state"`
}

func (a *Admin) Serializer() AdminSerializer {
	return AdminSerializer{
		ID:                a.ID,
		CreatedAt:         a.CreatedAt,
		UpdatedAt:         a.UpdatedAt,
		DeletedAt:         a.DeletedAt,
		Name:              a.Name,
		AuthToken:         a.AuthToken,
		EncryptedPassword: a.EncryptedPassword,
		Phone:             a.Phone,
		State:             a.State,
	}
}

func (a *Admin) UpdatePassword(password string) {
	var newPassword string
	// 传入的秘密加密处理
	newPassword = crypto.PasswordEncrypted(password)
	oldAuthToken := a.AuthToken
	a.EncryptedPassword = newPassword

	var admin model.Admin
	if dbError := init.POSTGRES.Where("auth_token = ?", oldAuthToken).First(&admin).Error; dbError != nil {
		fmt.Println("no exists user", dbError)
		return
	}

	if dbError := init.POSTGRES.Save(&admin).Error; dbError != nil {
		fmt.Println("update error", dbError)
		return
	}
}

// 密码检查
func (a *Admin) PasswordCheck(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
