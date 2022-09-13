package token

import (
	"errors"
	"os"
)

func GetTokenKey(tokenFlag string) (string, error) {
	var tokenString string

	if tokenFlag != "" {
		tokenString = tokenFlag
	} else {
		if os.Getenv("REPLICATE_TOKEN") != "" {
			tokenString = os.Getenv("REPLICATE_TOKEN")
		} else {
			return "", errors.New("No token provided. Specify Replicate API token with --token or REPLICATE_TOKEN environment variable.\n")
		}
	}

	return tokenString, nil
}
