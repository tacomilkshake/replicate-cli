package predictions

import (
	"github.com/jamiesteven/replicate-cli/internal/pkg/api"
	"github.com/tidwall/gjson"
)

func StartPrediction(body string, TokenKey string) (string, error) {
	url := "https://api.replicate.com/v1/predictions"

	response, err := api.ApiPost(url, body, TokenKey)
	if err != nil {
		return "", err
	}

	predictionId := gjson.Get(response, "id").String()
	return predictionId, nil
}

func CheckPrediction(predictionId string, TokenKey string) (string, error) {
	url := "https://api.replicate.com/v1/predictions/" + predictionId

	response, err := api.ApiGet(url, TokenKey)
	if err != nil {
		return "", err
	}

	return response, nil
}
