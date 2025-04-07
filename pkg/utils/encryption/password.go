package encryption

import (
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	const cost = 10
	pepper := viper.GetString("password_hash_salt")

	pepperedPassword := password + pepper

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(pepperedPassword), cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedBytes), nil
}
