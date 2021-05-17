// Ref: https://stackoverflow.com/questions/39481826/generate-6-digit-verification-code-with-golang

package codegen

import (
	"crypto/rand"
)

const otpChars = "1234567890"

// GenActCode generates an activation registration code
func GenActCode(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}
	return string(buffer), nil
}
