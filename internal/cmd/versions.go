package cmd

import (
	"errors"

	"github.com/jamiesteven/replicate-cli/internal/pkg/models"
	"github.com/jamiesteven/replicate-cli/internal/pkg/tables"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var versionsCmd = &cobra.Command{
	Use:   "versions [model]",
	Short: "Get versions for a model",
	Long:  "Get versions for a model, e.g. 'tencentarc/gfpgan'.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("no model specified - run 'replicate-cli versions [model]'")
		}

		response, err := models.GetModelVersions(args[0], TokenKey)
		if err != nil {
			return err
		}

		versions := gjson.Get(response, "@this").Array()

		table := tables.CreateTable()
		table.SetHeader([]string{"Version ID", "Published On"})
		for _, value := range versions {
			table.Append([]string{
				gjson.Get(value.String(), "id").String(),
				gjson.Get(value.String(), "created_at").String(),
			})
		}
		table.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionsCmd)
}
