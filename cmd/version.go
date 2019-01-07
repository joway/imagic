package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version number of imagic",
	Long:  `Version number of imagic`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Imagic v0.0.1")
	},
}
