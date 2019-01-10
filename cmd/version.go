package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var IMAGIC_VERSION = "v0.1.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Imagic %s\n", IMAGIC_VERSION)
	},
}
