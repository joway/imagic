package cmd

import (
	"fmt"
	"github.com/joway/imagic/pkg/image"
	"github.com/joway/imagic/pkg/util"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

var compressParallelCh chan int
var compressQuality int

func init() {
	compressCmd.Flags().IntVarP(&compressQuality,
		"quality", "q", 70,
		"Quantization of image compression",
	)
	compressCmd.Flags().IntVarP(&parallel,
		"parallel", "p", 4,
		"Number of concurrent tasks",
	)
	compressCmd.Flags().StringVarP(&output,
		"output", "o", "",
		"Output directory to write precessed images",
	)
	compressCmd.Flags().StringVarP(&suffix,
		"suffix", "s", "",
		"Suffix of precessed image filename",
	)

	rootCmd.AddCommand(compressCmd)
}

var compressCmd = &cobra.Command{
	Use:  "compress",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		compressParallelCh = make(chan int, parallel)
		wg := &sync.WaitGroup{}

		files := getFilesFromPatterns(args)
		for _, filename := range files {
			compressParallelCh <- 1
			wg.Add(1)

			outputFileName := getOutputFileName(filename, suffix, output)
			img, err := image.NewImageFromPath(filename)
			if err != nil {
				util.LogError(err)
				return
			}
			go func() {
				defer wg.Done()

				outImg, err := img.Compress(compressQuality)
				if err != nil {
					util.LogError(err)
					return
				}
				if err := outImg.Write(outputFileName); err != nil {
					util.LogError(err)
					return
				}

				util.LogInfo(fmt.Sprintf("%s compressed", filepath.Base(outputFileName)))

				<-compressParallelCh
			}()
		}
		wg.Wait()
	},
}
