package cmd

import (
	"fmt"
	"github.com/joway/imagic/pkg/image"
	"github.com/joway/imagic/pkg/util"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

var resizeParallelCh chan int
var resizeWidth int
var resizeHeight int

func init() {
	resizeCmd.Flags().IntVarP(&resizeWidth, "width", "w", 1024, "number of concurrent tasks")
	resizeCmd.Flags().IntVarP(&resizeHeight, "height", "l", 0, "number of concurrent tasks")

	rootCmd.AddCommand(resizeCmd)
}

var resizeCmd = &cobra.Command{
	Use:  "resize",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resizeParallelCh = make(chan int, parallel)
		wg := &sync.WaitGroup{}

		files := getFilesFromPatterns(args)
		for _, filename := range files {
			resizeParallelCh <- 1
			wg.Add(1)
			outputFileName := getOutputFileName(filename, suffix, output)
			img, err := image.NewImageFromPath(filename)
			if err != nil {
				util.LogError(err)
				return
			}

			go func() {
				defer wg.Done()

				outImg, err := img.Resize(resizeWidth, resizeHeight)
				if err != nil {
					util.LogError(err)
					return
				}
				if err := outImg.Write(outputFileName); err != nil {
					util.LogError(err)
					return
				}

				util.LogInfo(
					fmt.Sprintf("%s resized", filepath.Base(outputFileName)),
				)

				<-resizeParallelCh
			}()
		}
		wg.Wait()
	},
}
