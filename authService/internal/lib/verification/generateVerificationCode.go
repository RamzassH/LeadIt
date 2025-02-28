package verification

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateVerificationCode() (string, error) {
	b := make([]byte, 4)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b)[:6], nil
}
