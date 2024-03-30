// Helper functions for the backend package
package backend

import (
	"fmt"
	"strings"
	"unicode"
)


// InitMessage prints a message when the server starts
func InitMessage() {
	fmt.Printf("===============================================\n")
	fmt.Printf("Starting Realtime forum\n")
	fmt.Printf("Server is running on port: 8080\n")
	fmt.Printf("Press Ctrl+C to stop the server\n")
	fmt.Printf("===============================================\n")
}

// Helper function to check password strength
func CheckPasswordStrength(password string) bool {
	if len(password) < 4 {
		return false
	}

	hasLower := false
	hasDigit := false

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		}

		if hasLower && hasDigit {
			return true
		}
	}

	return false
}

func NicknameCheck(un string) bool {
	if strings.Contains(un, "@") {
		return true
	} else {
		return false
	}
}
