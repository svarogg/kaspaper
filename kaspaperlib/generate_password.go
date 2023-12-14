package kaspaperlib
import (
	"crypto/rand"
)

func generatePassword() string {
	password := make([]byte, 16)
	_, err := rand.Read(password)
	if err != nil {
		return nil, err
	}

	return string(password), nil
}
