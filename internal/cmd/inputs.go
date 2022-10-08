package cmd

import (
	"errors"
	"sort"
	"strconv"

	"github.com/jamiesteven/replicate-cli/internal/pkg/models"
	"github.com/jamiesteven/replicate-cli/internal/pkg/tables"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var inputsCmd = &cobra.Command{
	Use:   "inputs [model] [?version]",
	Short: "Get inputs for a model (or model version)",
	Long:  "Get inputs for a model (or model version). Provide a model name, i.e. 'tencentarc/gfpgan', or a model and version.",
	RunE: func(cmd *cobra.Command, args []string) error {

		var response string

		if len(args) > 1 {
			model, err := models.GetModelVersion(args[0], args[1], TokenKey)
			if err != nil {
				return err
			}
			version := gjson.Get(model, "openapi_schema").String()
			response = version
		} else if len(args) == 1 {
			model, err := models.GetModel(args[0], TokenKey)
			if err != nil {
				return err
			}
			version := gjson.Get(model, "latest_version.openapi_schema").String()
			response = version
		} else {
			return errors.New("no model specified - run 'replicate-cli inputs [model] [?version]'")
		}

		propKeys := gjson.Get(response, "components.schemas.Input.properties.@keys").Array()
		propValues := gjson.Get(response, "components.schemas.Input.properties")

		var props [][]string
		for _, value := range propKeys {
			propValue := propValues.Get(value.String()).String()
			props = append(props, []string{
				value.String(),
				gjson.Get(propValue, "type").String(),
				gjson.Get(propValue, "default").String(),
				gjson.Get(propValue, "description").String(),
				gjson.Get(propValue, "enum").String(),
				gjson.Get(propValue, "x-order").String(),
			})
		}

		// Sort by x-order
		sort.Slice(props, func(i, j int) bool {
			before, _ := strconv.Atoi(props[i][5])
			after, _ := strconv.Atoi(props[j][5])
			return before < after
		})

		table := tables.CreateTable()
		table.SetHeader([]string{"Property", "Type", "Default", "Description", "E-Num"})
		table.SetRowLine(true)
		for _, value := range props {
			// Remove x-order from appended row
			table.Append(value[:len(value)-1])
		}
		table.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(inputsCmd)
}
