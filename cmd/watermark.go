package cmd

import (
	"fmt"
	"github.com/joway/imagic/pkg/image"
	"github.com/joway/imagic/pkg/util"
	"github.com/spf13/cobra"
	"path/filepath"
	"sync"
)

var watermarkParallelCh chan int
var watermarkTextureFn string
var watermarkX int
var watermarkY int

func init() {
	watermarkCmd.Flags().StringVarP(
		&watermarkTextureFn, "texture",
		"t", "",
		"texture filename",
	)
	watermarkCmd.Flags().IntVarP(
		&watermarkX, "X",
		"x", 0,
		"pos X",
	)
	watermarkCmd.Flags().IntVarP(
		&watermarkY, "Y",
		"y", 0,
		"pos Y",
	)

	rootCmd.AddCommand(watermarkCmd)
}

var watermarkCmd = &cobra.Command{
	Use:  "watermark",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		watermarkParallelCh = make(chan int, parallel)
		wg := &sync.WaitGroup{}

		texture, err := image.NewImageFromPath(watermarkTextureFn)
		if err != nil {
			util.LogFatal(err)
		}
		files := getFilesFromPatterns(args)
		for _, filename := range files {
			watermarkParallelCh <- 1
			wg.Add(1)
			outputFileName := getOutputFileName(filename, suffix, output)
			img, err := image.NewImageFromPath(filename)
			if err != nil {
				util.LogError(err)
				return
			}

			go func() {
				defer wg.Done()

				outImg, err := img.WaterMark(texture, watermarkX, watermarkY)
				if err != nil {
					util.LogError(err)
					return
				}
				if err := outImg.Write(outputFileName); err != nil {
					util.LogError(err)
					return
				}

				util.LogInfo(
					fmt.Sprintf("%s watermarked", filepath.Base(outputFileName)),
				)

				<-watermarkParallelCh
			}()
		}
		wg.Wait()
	},
}
