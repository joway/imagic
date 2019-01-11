package cmd

import (
	"github.com/spf13/cobra"
)

var compressQuality int

func init() {
	compressCmd.Flags().IntVarP(&compressQuality,
		"quality", "q", 70,
		"Quality of image compression",
	)

	rootCmd.AddCommand(compressCmd)
}

var compressCmd = &cobra.Command{
	Use:  "compress",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		execCmd(
			args, "compress", compressQuality,
			0, 0, nil, "", "",
		)
	},
}
