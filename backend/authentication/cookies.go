package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	ErrInvalidValue = errors.New("invalid value")
)

func WriteEncrypted(w http.ResponseWriter, cookie http.Cookie, secretKey []byte) error {
	// Create a new AES cipher block.
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return fmt.Errorf("failed to create AES cipher block: %w", err)
	}

	// Wrap the cipher block in Galois Counter Mode.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	// Create a unique nonce containing 12 random bytes.
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Prepare the plaintext input for encryption.
	plaintext := fmt.Sprintf("%s:%s", cookie.Name, cookie.Value)

	// Encrypt the data.
	encryptedValue := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)

	// Encode the encrypted value to base64 for safe transmission.
	cookie.Value = base64.URLEncoding.EncodeToString(encryptedValue)

	// Write the cookie.
	http.SetCookie(w, &cookie)
	return nil
}

func ReadEncrypted(r *http.Request, name string, secretKey []byte) (string, error) {
	// Retrieve the cookie.
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve cookie: %w", err)
	}

	// Decode the base64-encoded encrypted value.
	encryptedValue, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return "", fmt.Errorf("failed to decode encrypted value: %w", err)
	}

	// Create a new AES cipher block.
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher block: %w", err)
	}

	// Wrap the cipher block in Galois Counter Mode.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	// Validate the length of the encrypted value.
	nonceSize := aesGCM.NonceSize()
	if len(encryptedValue) < nonceSize {
		return "", ErrInvalidValue
	}

	// Separate the nonce and the ciphertext.
	nonce := encryptedValue[:nonceSize]
	ciphertext := encryptedValue[nonceSize:]

	// Decrypt and authenticate the data.
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", ErrInvalidValue
	}

	// Parse the plaintext to retrieve the cookie name and value.
	expectedName, value, ok := strings.Cut(string(plaintext), ":")
	if !ok || expectedName != name {
		return "", ErrInvalidValue
	}

	return value, nil
}

func CreateCookie(name, value, path string, duration time.Duration) http.Cookie {
	return http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Expires:  time.Now().Add(duration),
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
	}
}