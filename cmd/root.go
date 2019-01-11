package cmd

import (
	"github.com/joway/imagic/pkg/util"
	"github.com/spf13/cobra"
)

var parallel int
var output string
var suffix string

func init() {
	rootCmd.PersistentFlags().IntVarP(&parallel, "parallel", "p", 4, "Number of parallel tasks")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output directory")
	rootCmd.PersistentFlags().StringVarP(&suffix, "suffix", "s", "", "Suffix of precessed image filename, format: [base_filename][suffix].[ext]")
}

var rootCmd = &cobra.Command{
	Use:   "imagic",
	Short: "An easy and fast tool to process images.",
	Long: `
Imagic is an easy and fast tool to process images.

Created by Joway Wang (https://joway.io).
Issues on https://github.com/joway/imagic/issues.
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// Execute root cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		util.LogFatal(err)
	}
}
