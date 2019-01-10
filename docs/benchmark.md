# Imagic Benchmark

## Environment

MacBook Pro (13-inch, 2016)

## Compress

### JPEG

#### TestData

 83M ./testdata/benchmark/jpg

#### Result


##### Parallel

- parallel=6,quality=70: 9940 ms
- parallel=15,quality=70: 11480 ms
- parallel=30,quality=70: 9770 ms


##### Quality

- parallel=30,quality=90: 11650 ms
- parallel=30,quality=70: 22690 ms
- parallel=30,quality=30: 23110 ms


### PNG

#### TestData

 37M ./testdata/benchmark/png

#### Result


##### Parallel

- parallel=6,quality=70: 5890 ms
- parallel=15,quality=70: 7690 ms
- parallel=30,quality=70: 6940 ms
##### Quality

- parallel=30,quality=90: 6650 ms
- parallel=30,quality=70: 6070 ms
- parallel=30,quality=30: 5260 ms
