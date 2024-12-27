package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Errors for validation
var (
	ErrInvalidTokenFormat = errors.New("invalid token format")
	ErrInvalidTokenType   = errors.New("invalid token type")
)

// ValidateAndExtractUserID validates a token and extracts the userID.
// `tokenName` should be "access_token" or "refresh_token" to distinguish token types.
func ValidateAndExtractUserID(r *http.Request, tokenName string, secretKey []byte, expectedType string) (int, error) {
	// Read and decrypt the token
	plaintext, err := ReadEncrypted(r, tokenName, secretKey)
	if err != nil {
		return 0, fmt.Errorf("failed to read or decrypt token: %w", err)
	}

	// Parse the plaintext value
	parts := strings.Split(plaintext, "|")
	if len(parts) != 2 {
		return 0, ErrInvalidTokenFormat
	}

	// Extract userID and type
	var userID int
	var tokenType string
	for _, part := range parts {
		keyValue := strings.Split(part, ":")
		if len(keyValue) != 2 {
			return 0, ErrInvalidTokenFormat
		}
		switch keyValue[0] {
		case "userID":
			fmt.Sscanf(keyValue[1], "%d", &userID)
		case "type":
			tokenType = keyValue[1]
		default:
			return 0, ErrInvalidTokenFormat
		}
	}

	// Validate token type
	if tokenType != expectedType {
		return 0, ErrInvalidTokenType
	}

	// Return the extracted userID
	return userID, nil
}