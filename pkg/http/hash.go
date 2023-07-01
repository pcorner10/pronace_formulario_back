package http

import "golang.org/x/crypto/bcrypt"

type PasswordHash interface {
	HashPassword(password string) (string, error)
	ComparePasswords(hashedPassword, password string) bool
}

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// ComparePasswords compara una contraseña en texto plano con un hash y devuelve true si coinciden o false si no.
func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
