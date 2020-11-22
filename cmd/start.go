package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "my-first-ls",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Hello World")
	},
}

// Execute executes startCmd
func Execute() {
	if err := startCmd.Execute(); err != nil {
		// TODO: use zap and errors
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
