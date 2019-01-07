package cmd

import (
	"fmt"
	"github.com/joway/imagic/pkg/image"
	"github.com/joway/imagic/pkg/util"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var radio *float32
var output *string
var suffix *string

func init() {
	radio = compressCmd.Flags().Float32P("radio", "r", 1.0, "Source directory to read from")
	output = compressCmd.Flags().StringP("output", "o", "", "Source directory to read from")
	suffix = compressCmd.Flags().StringP("suffix", "s", "", "Source directory to read from")

	rootCmd.AddCommand(compressCmd)
}

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress images",
	Long:  `Compress images`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filesPattern := args[0]
		files, err := filepath.Glob(filesPattern)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		for _, filename := range files {
			var outputPath string
			if *output != "" {
				absOutput, _ := filepath.Abs(*output)
				outputPath = absOutput
			} else {
				outputPath = filepath.Dir(filename)
			}
			baseFilename := filepath.Base(filename)
			dotPos := util.IndexOf(baseFilename, ".", true)
			if dotPos == -1 {
				continue
			}
			outputFilename := outputPath + "/" + baseFilename[:dotPos] + *suffix + baseFilename[dotPos:]

			img, err := image.NewImageFromPath(filename)
			if err != nil {
				log.Fatal(err)
			}
			outImg, err := img.Compress(0.5)
			if err != nil {
				log.Fatal(err)
			}
			err = outImg.Write(outputFilename)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(outputFilename)
		}
	},
}
