package authentication

import "golang.org/x/crypto/bcrypt"

// EncryptPassword generates a bcrypt salt/hash combination from the given password,
// then returns that salt/hash as a string. If an error occurs, the returned salt/hash
// is empty string "", and a non-nil error is returned.
func EncryptPassword(password string) (string, error) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPw), nil
}

// CheckPassword accepts a salt/hash and a plaintext password. If that password
// is a match to the given salt/hash, returns true. Otherwise, returns false.
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
