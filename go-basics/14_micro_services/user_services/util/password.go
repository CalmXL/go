package util

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
)

var opts = &password.Options{10, 10000, 30, md5.New}

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

// func main() {
// 	// Using the default options
// 	salt, encodedPwd := password.Encode("generic password", nil)
// 	check := password.Verify("generic password", salt, encodedPwd, nil)
// 	fmt.Println(check) // true

// 	// Using custom options
// 	salt, encodedPwd = password.Encode("generic password", options)
// 	check = password.Verify("generic password", salt, encodedPwd, options)
// 	fmt.Println(check) // true
// }
