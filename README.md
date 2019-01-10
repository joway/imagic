# Imagic

An easy and fast tool to process images.

## Install

```shell
go get github.com/joway/imagic
```

## Usage

### Supported format

- png
- jpeg

### CLI Options

```
imagic -h
```

### Compress

```shell
imagic compress -q 70 -p 10 \
	-s .comp -o ./output \
	./testdata/**/*.png
```

### Resize

```shell
imagic resize -w 320 -p 10 \
	-s .comp -o ./output \
	./testdata/**/*.jpg
```

## Benchmark

[benchmark.md](docs/benchmark.md)

## Acknowledgement

- [libimagequant-go](https://github.com/joway/libimagequant-go)
- [imaging](https://github.com/disintegration/imaging)






