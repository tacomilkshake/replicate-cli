package cmd

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var versionsCmd = &cobra.Command{
	Use:   "versions [model]",
	Short: "Get versions for a specific model.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("No model specified. Run 'replicate-cli versions [model]'.\n")
		}

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

		results := gjson.Get(resp.String(), "results").Array()

		for _, value := range results {
			id := gjson.Get(value.String(), "id")
			created_at := gjson.Get(value.String(), "created_at")
			fmt.Println(id, created_at)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionsCmd)
}
