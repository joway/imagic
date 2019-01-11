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
	compressCmd.Flags().IntVarP(&parallel,
		"parallel", "p", 4,
		"Number of parallel tasks",
	)
	compressCmd.Flags().StringVarP(&output,
		"output", "o", "",
		"Output directory",
	)
	compressCmd.Flags().StringVarP(&suffix,
		"suffix", "s", "",
		"Suffix of precessed image file",
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
