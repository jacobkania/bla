package authentication

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPw), nil
}

func CheckPassword(hashedPw string, password string) bool {
	if len(hashedPw) == 0 {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(password))
	if err != nil {
		return false
	}

	return true
}
