# Imagic Benchmark

## Environment

MacBook Pro (13-inch, 2016)

## Compress

### JPEG

#### TestData

 83M ./testdata/benchmark/jpg

#### Result


##### Parallel

- parallel=6,quality=70: 21210 ms
- parallel=15,quality=70: 19450 ms
- parallel=30,quality=70: 19660 ms


##### Quality

- parallel=30,quality=90: 20890 ms
- parallel=30,quality=70: 25180 ms
- parallel=30,quality=30: 24310 ms


### PNG

#### TestData

 37M ./testdata/benchmark/png

#### Result


##### Parallel

- parallel=6,quality=70: 6590 ms
- parallel=15,quality=70: 6750 ms
- parallel=30,quality=70: 6430 ms


##### Quality

- parallel=30,quality=90: 6850 ms
- parallel=30,quality=70: 7060 ms
- parallel=30,quality=30: 7220 ms
