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
			return "", errors.New("no token provided. specify Replicate API token with --token or REPLICATE_TOKEN environment variable")
		}
	}

	return tokenString, nil
}
