package main

import (
	"errors"
	"github.com/bigbigliu/web_app/cmd/app"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/spf13/cobra"
	"os"

	"github.com/bigbigliu/web_app/cmd/api"
)

var rootCmd = &cobra.Command{
	Use:          "go-admin",
	Short:        "go-admin",
	SilenceUsage: true,
	Long:         `go-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(app.StartCmd)
}

// Execute : apply commands
func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
