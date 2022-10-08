package models

import (
	"errors"
	"strings"

	"github.com/jamiesteven/replicate-cli/internal/pkg/api"
	"github.com/tidwall/gjson"
)

func CheckModelString(model string) bool {
	if strings.Contains(model, "/") {
		return true
	} else {
		return false
	}
}

func GetVersionId(input string, TokenKey string) (string, error) {
	if CheckModelString(input) {
		var err error
		version, err := GetModelLatestVersion(input, TokenKey)
		if err != nil {
			return "", err
		}
		return version, nil
	} else {
		return input, nil
	}
}

func GetModel(model string, TokenKey string) (string, error) {
	if !CheckModelString(model) {
		return "", errors.New("invalid model name, use [username]/[modelname]")
	}

	url := "https://api.replicate.com/v1/models/" + model

	model, err := api.ApiGet(url, TokenKey)
	if err != nil {
		return "", err
	}

	return model, nil
}

func GetModelLatestVersion(model string, TokenKey string) (string, error) {
	url := "https://api.replicate.com/v1/models/" + model

	response, err := api.ApiGet(url, TokenKey)
	if err != nil {
		return "", err
	}

	latest := gjson.Get(response, "latest_version.id").String()
	return latest, nil
}

func GetModelVersion(model string, version string, TokenKey string) (string, error) {
	if !CheckModelString(model) {
		return "", errors.New("invalid model syntax")
	}

	url := "https://api.replicate.com/v1/models/" + model + "/versions/" + version

	response, err := api.ApiGet(url, TokenKey)
	if err != nil {
		return "", err
	}

	versions := gjson.Get(response, "@this").String()
	return versions, nil
}

func GetModelVersions(model string, TokenKey string) (string, error) {
	if !CheckModelString(model) {
		return "", errors.New("invalid model syntax")
	}

	url := "https://api.replicate.com/v1/models/" + model + "/versions"

	response, err := api.ApiGet(url, TokenKey)
	if err != nil {
		return "", err
	}

	versions := gjson.Get(response, "results").String()
	return versions, nil
}
