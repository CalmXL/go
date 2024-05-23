package util

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
)

var opts = &password.Options{SaltLen: 10, Iterations: 10000, KeyLen: 30, HashFunction: md5.New}

func GeneratePassword(rawPassword string) string {
	salt, encodedPwd := password.Encode(rawPassword, opts)
	encodePassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	return encodePassword
}

func VerifyPassword(rawPassword string, encodePassword string) bool {
	passwordInfo := strings.Split(encodePassword, "$")
	salt := passwordInfo[2]
	encodePwd := passwordInfo[3]
	isPass := password.Verify(rawPassword, salt, encodePwd, opts)

	return isPass
}
