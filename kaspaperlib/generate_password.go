package kaspaperlib
import (
	"crypto/rand"
)

func generatePassword() (string, error) {
	password := make([]byte, 16)
	_, err := rand.Read(password)
	if err != nil {
		return "", err
	}

	return string(password), nil
}
