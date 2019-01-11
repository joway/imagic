package cmd

import (
	"github.com/spf13/cobra"
)

var resizeWidth int
var resizeHeight int

func init() {
	resizeCmd.Flags().IntVarP(&resizeWidth,
		"width", "w", 1024,
		"number of concurrent tasks",
	)
	resizeCmd.Flags().IntVarP(&resizeHeight,
		"height", "l", 0,
		"number of concurrent tasks",
	)

	rootCmd.AddCommand(resizeCmd)
}

var resizeCmd = &cobra.Command{
	Use:  "resize",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		execCmd(
			args, "resize", 0,
			resizeWidth, resizeHeight, nil, "", "",
		)
	},
}
