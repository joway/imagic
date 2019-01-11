package cmd

import (
	"github.com/joway/imagic/pkg/image"
	"github.com/joway/imagic/pkg/util"
	"github.com/spf13/cobra"
)

var watermarkTextureFn string
var watermarkXStr string
var watermarkYStr string

func init() {
	watermarkCmd.Flags().StringVarP(
		&watermarkTextureFn, "texture",
		"t", "",
		"texture filename",
	)
	watermarkCmd.Flags().StringVarP(
		&watermarkXStr, "X",
		"x", "+0",
		"pos X",
	)
	watermarkCmd.Flags().StringVarP(
		&watermarkYStr, "Y",
		"y", "+0",
		"pos Y",
	)

	rootCmd.AddCommand(watermarkCmd)
}

var watermarkCmd = &cobra.Command{
	Use:  "watermark",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		texture, err := image.NewImageFromPath(watermarkTextureFn)
		if err != nil {
			util.LogFatal(err)
		}

		execCmd(
			args, "watermark", 0,
			0, 0, texture, watermarkXStr, watermarkXStr,
		)
	},
}
