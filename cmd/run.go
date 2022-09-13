package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var runCmd = &cobra.Command{
	Use:   "run [version/model] [input]",
	Short: "Run a model.",
	Long:  "Run a model. Provide a model version ID, or a model name, i.e. 'tencentarc/gfpgan'",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("No model and/or input specified. Run 'replicate-cli run [model] [input]'.\n")
		}

		// Get latest version ID if args[0] is a model name, i.e. contains a slash
		var versionId string
		if strings.Contains(args[0], "/") {
			url := "https://api.replicate.com/v1/models/" + args[0] + "/versions"

			client := resty.New()
			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetHeader("Authorization", "Token "+TokenKey).
				Get(url)
			if err != nil {
				return err
			} else if resp.Status() == "401 Unauthorized" {
				return errors.New("Unauthorized. Specify Replicate API token with --token or REPLICATE_TOKEN environment variable.\n")
			} else if resp.Status() == "404 Not Found" {
				return errors.New("Model not found.\n")
			}
			versionId = gjson.Get(resp.String(), "results.0.id").String()
		} else {
			versionId = args[0]
		}

		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Authorization", "Token "+TokenKey).
			SetBody(`{"version":"` + versionId + `","input":{"image":"` + args[1] + `"}}`).
			Post("https://api.replicate.com/v1/predictions")
		if err != nil {
			return err
		} else if resp.Status() == "401 Unauthorized" {
			return errors.New("Unauthorized. Specify Replicate API token with --token or REPLICATE_TOKEN environment variable.\n")
		} else if resp.Status() == "404 Not Found" {
			return errors.New("Model version not found.\n")
		} else if resp.Status() == "400 Bad Request" {
			return errors.New("Model version not found.\n") // TODO: Issue: Replicate API should replace this response as a 404 and leave 400 for malformed requests.
		}
		predictionUrl := gjson.Get(resp.String(), "urls.get").String()
		predictionId := gjson.Get(resp.String(), "id").String()
		fmt.Println("prediction url: https://replicate.com/p/" + predictionId)

		for {
			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetHeader("Authorization", "Token "+TokenKey).
				Get(predictionUrl)
			if err != nil {
				return err
			}
			status := gjson.Get(resp.String(), "status").String()

			s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
			s.Prefix = status + ` `
			s.Start()
			time.Sleep(3 * time.Second)
			s.Stop()

			if status == "succeeded" {
				output := gjson.Get(resp.String(), "output")
				fmt.Println("succeeded.\n", output)
				break
			} else if status == "failed" {
				return errors.New("Prediction failed.\n")
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
