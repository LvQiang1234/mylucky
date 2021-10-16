package mylucky

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
	"time"
)

const passwordLength = 256

type password [passwordLength]byte

func init() {
	rand.Seed(time.Now().Unix())
}

func (password *password) String() string {
	return base64.StdEncoding.EncodeToString(password[:])
}

func ParseLittlePassword(passwordString string) (*password, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(passwordString))
	if err != nil || len(bs) != passwordLength {
		return nil, errors.New("不合法的密码")
	}
	password := password{}
	copy(password[:], bs)
	return &password, nil
}

func RandLittlePassword() string {
	intArr := rand.Perm(passwordLength)
	password := &password{}
	for i, v := range intArr {
		password[i] = byte(v)
		if i == v {
			return RandLittlePassword()
		}
	}
	return password.String()
}


