package cmd

import (
	"os"

	"github.com/jamiesteven/replicate-cli/internal/pkg/token"
	"github.com/spf13/cobra"
)

var TokenFlag string
var TokenKey string

var rootCmd = &cobra.Command{
	Use:   "replicate [command]",
	Short: "replicate-cli",
	Long: `replicate-cli 💫 A command line interface for Replicate ❤️
Version 0.2.2 by Jamie Steven (github.com/jamiesteven/replicate-cli)`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		tokenResp, err := token.GetTokenKey(TokenFlag)
		if err != nil {
			return err
		}
		TokenKey = tokenResp
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&TokenFlag, "token", "t", "", "Replicate API token. Uses REPLICATE_TOKEN environment variable if not specified.")
}
