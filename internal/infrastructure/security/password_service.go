package security

import "golang.org/x/crypto/bcrypt"

type PasswordService interface {
	Hash(password string) (string, error)
	Check(hashedPassword, password string) error
}

type bcryptPasswordService struct {
	//
}

func NewBcryptPasswordService() PasswordService {
	return &bcryptPasswordService{}
}

func (s *bcryptPasswordService) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *bcryptPasswordService) Check(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
