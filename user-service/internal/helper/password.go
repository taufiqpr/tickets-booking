package helper

import "golang.org/x/crypto/bcrypt"

type PasswordHelper struct{}

func NewPasswordHelper() *PasswordHelper {
	return &PasswordHelper{}
}

func (p *PasswordHelper) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (p *PasswordHelper) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
