package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var imageVersion = "v0.2.2"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Imagic %s\n", imageVersion)
	},
}
