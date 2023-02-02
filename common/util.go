package common

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomOrderNum() (result string) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_=+/?!@#$%^&*"
	result, _ = randomString(letters, 10)

	return
}

func randomString(letters string, length int) (result string, err error) {
	if letters == "" || length == 0 {
		return
	}

	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
