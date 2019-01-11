package cmd

import (
	"github.com/spf13/cobra"
)

var resizeWidth int
var resizeHeight int

func init() {
	resizeCmd.Flags().IntVarP(&resizeWidth,
		"width", "w", 0,
		"Width of output image, default adaptive",
	)
	resizeCmd.Flags().IntVarP(&resizeHeight,
		"height", "l", 0,
		"Height of output image, default adaptive",
	)

	rootCmd.AddCommand(resizeCmd)
}

var resizeCmd = &cobra.Command{
	Use:  "resize",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if resizeWidth == 0 && resizeHeight == 0 {
			resizeWidth = 1024
		}
		execCmd(
			args, "resize", 0,
			resizeWidth, resizeHeight, nil, "", "",
		)
	},
}
