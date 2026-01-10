package cmd

import (
	"fmt"

	"github.com/kanaya/jobwatch/cli/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize jobwatch configuration",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Generate(config.DefaultFilename, config.DefaultTemplate)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("created: %s\n", config.DefaultFilename)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
