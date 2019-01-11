package cmd

import (
	"fmt"
	"github.com/joway/imagic/pkg/image"
	"github.com/joway/imagic/pkg/util"
	"sync"
)

func execCmd(
	patterns []string, cmd string,
	quality,
	width, height int,
	texture *image.Image, x, y string,
) {
	parallelCh := make(chan int, parallel)
	wg := &sync.WaitGroup{}

	files := getFilesFromPatterns(patterns)
	total := float64(len(files))
	cur := 0.0
	for _, filename := range files {
		parallelCh <- 1
		wg.Add(1)

		img, err := image.NewImageFromPath(filename)
		if err != nil {
			util.LogFatal(err)
		}

		outputFileName := getOutputFileName(filename, suffix, output)
		go func() {
			defer wg.Done()
			defer func() { cur++ }()

			var outImg *image.Image
			switch cmd {
			case "compress":
				outImg, err = img.Compress(quality)
			case "resize":
				outImg, err = img.Resize(width, height)
			case "watermark":
				outImg, err = img.WaterMark(texture, x, y)
			}

			if err != nil {
				util.LogError(err)
			} else {
				if err := outImg.Write(outputFileName); err != nil {
					util.LogError(err)
				}
			}

			util.LogNail(fmt.Sprintf("%d%%", int(((cur+1)/total)*100)))

			<-parallelCh
		}()
	}

	wg.Wait()
}
