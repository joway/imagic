# Imagic Benchmark

## Environment

MacBook Pro (13-inch, 2016)

## Compress

### JPEG

#### TestData

 83M ./testdata/benchmark/jpg

#### Result


##### Parallel

- parallel=1,quality=70: 40750 ms
- parallel=3,quality=70: 20110 ms
- parallel=9,quality=70: 25000 ms


##### Quality

- parallel=10,quality=90: 26280 ms
- parallel=10,quality=70: 22070 ms
- parallel=10,quality=30: 19340 ms


### PNG

#### TestData

 37M ./testdata/benchmark/png

#### Result


##### Parallel

- parallel=1,quality=70: 12650 ms
- parallel=3,quality=70: 5990 ms
- parallel=9,quality=70: 5230 ms


##### Quality

- parallel=10,quality=90: 5800 ms
- parallel=10,quality=70: 5950 ms
- parallel=10,quality=30: 4940 ms
