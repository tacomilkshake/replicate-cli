package api

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

func ApiGet(url string, TokenKey string) (string, error) {
	client := resty.New()
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Token "+TokenKey).
		Get(url)
	if response.Status() == "401 Unauthorized" {
		return "", errors.New("unauthorized: specify Replicate API token with --token or REPLICATE_TOKEN environment variable")
	} else if response.Status() == "404 Not Found" {
		return "", errors.New("not found")
	} else if err != nil {
		return "", err
	}
	return response.String(), nil
}

func ApiPost(url string, body string, TokenKey string) (string, error) {
	client := resty.New()
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Token "+TokenKey).
		SetBody(body).
		Post(url)
	if err != nil {
		return "", err
	} else if response.Status() == "401 Unauthorized" {
		return "", errors.New("unauthorized: specify Replicate API token with --token or REPLICATE_TOKEN environment variable")
	} else if response.Status() == "404 Not Found" {
		return "", errors.New("model version not found")
	} else if response.Status() == "400 Bad Request" {
		return "", errors.New("model version not found") // TODO: Issue: Replicate API should replace this response as a 404 and leave 400 for malformed requests.
	}
	return response.String(), nil
}
