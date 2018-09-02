package crypto

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// 检查签名
func CheckSignInName(value string) bool {
	if value == "" {
		return false
	}
	trimValue := strings.TrimSpace(value)
	if len(trimValue) < 8 {
		return false
	}
	return true
}

//自动生成token
func GenerateToken() string {
	b2 := make([]byte, 200)
	_, err := rand.Read(b2)
	if err != nil {
		return fmt.Sprintf("%x", b2) //输出１６进制数
	}
	return ""
}

// 密码加密
func PasswordEncrypted(value string) string {
	password := []byte(string(value))
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", hashedPassword)
}
