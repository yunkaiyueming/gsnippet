package main

import (
	"crypto/cipher"
	"fmt"

	"golang.org/x/crypto/blowfish"
)

func Auth(input, salt string) (bool, error) {
	cp, err := blowfish.NewSaltedCipher([]byte(input), []byte(salt))
	if err != nil {
		return false, err
	}

	cryptedPwd := make([]byte, 8)
	cp.Encrypt(cryptedPwd, cryptedPwd)

	fmt.Println([]byte(input), cryptedPwd)
	if input == string(cryptedPwd) {
		fmt.Println(true)
		return true, nil
	} else {
		fmt.Println(false)
		return false, nil
	}
}

func main() {
	fmt.Println(blowfishEncrypt([]byte("qwe@123"), []byte("$2a$07$L2wAcTqxNVYSYKkMcmnEb.6XHjWucmogUVR6uKrqF13SnDkFC6lCK")))
}
