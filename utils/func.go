package utils

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"log"
)

//对用户密码进行加密存储
func ScryptPw(pwd string) string {
	salt := []byte{3, 48, 209, 38, 48, 22, 29, 99} //随机的盐值
	hashPw, err := scrypt.Key([]byte(pwd), salt, 16384, 8, 1, 10)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(hashPw)
}
