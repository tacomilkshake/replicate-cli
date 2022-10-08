package cmd

import (
	"errors"

	"github.com/jamiesteven/replicate-cli/internal/pkg/models"
	"github.com/jamiesteven/replicate-cli/internal/pkg/tables"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var infoCmd = &cobra.Command{
	Use:   "info [model]",
	Short: "Get info for a model",
	Long:  "Get info for a model, e.g. 'tencentarc/gfpgan'.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("no model specified - run 'replicate-cli info [model]'")
		}

		model, err := models.GetModel(args[0], TokenKey)
		if err != nil {
			return err
		}

		modelKeys := gjson.Get(model, "@keys").Array()
		modelValues := gjson.Get(model, "@values").Array()

		var props [][]string
		for n, value := range modelKeys {
			props = append(props, []string{
				value.String(),
				modelValues[n].String(),
			})
		}
		// Remove latest_version, the last element, from props
		props = props[:len(props)-1]

		table := tables.CreateTable()
		table.AppendBulk(props)
		table.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
