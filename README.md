# Imagic

## Purpose

A easy and fast tool to process images.

### Support image format

- png
- jpeg

## Install

```shell
go get github.com/joway/imagic
```

## Usage

### compress

```shell
imagic compress --ratio 0.7 --parallel 10 --suffix .comp --output ./testdata/output ./testdata/*.png
```

### resize

```shell
imagic resize --ratio 0.7 --parallel 10 --suffix .comp --output ./testdata/output ./testdata/*.png
imagic resize --width 320  *.png
```
