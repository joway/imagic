# Imagic

<img width="128px" src="logo.png" alt="logo">

![GitHub release](https://img.shields.io/github/tag/joway/imagic.svg?label=release)
[![Travis Build](https://travis-ci.org/joway/imagic.svg?branch=master)](https://travis-ci.org/joway/imagic)
[![Go Report Card](https://goreportcard.com/badge/github.com/joway/imagic)](https://goreportcard.com/report/github.com/joway/imagic)
[![codecov](https://codecov.io/gh/joway/imagic/branch/master/graph/badge.svg)](https://codecov.io/gh/joway/imagic)
[![CircleCI](https://circleci.com/gh/joway/imagic.svg?style=shield)](https://circleci.com/gh/joway/imagic)

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

### Watermark

```shell
imagic watermark -x -600 -y -100 \
	-t ./testdata/images/texture.png \
	-p 10 -o ./output \
	./testdata/**/*.jpg
```

![watermark](watermark.png)

## Benchmark

[benchmark.md](docs/benchmark.md)

## Acknowledgement

- [libimagequant-go](https://github.com/joway/libimagequant-go)
- [imaging](https://github.com/disintegration/imaging)
- [mergi](https://github.com/noelyahan/mergi)


