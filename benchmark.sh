# /bin/bash

DOC=docs/benchmark.md

go build


document(){
echo $1 >> $DOC
}
datasize(){
image_dir=$1
size=$(du -hs $image_dir)

document "#### TestData\n\n${size}\n"
}
benchmark(){

format=$1
parallel=$2
quality=$3
image_dir=$4

START=$(python -c 'import time; print(time.time() * 1000)')
CMD="./imagic compress -q ${quality} -p ${parallel} \
             -s .comp -o ./output \
             ${image_dir}/*
"
eval $CMD
END=$(python -c 'import time; print(time.time() * 1000)')
RUNTIME=$(python -c "print(int(${END} - ${START}))")
document "- parallel=${parallel},quality=${quality}: ${RUNTIME} ms"
}

rm $DOC
document "# Imagic Benchmark\n"
document "## Environment\n"
document "MacBook Pro (13-inch, 2016)\n"

document "## Compress\n"

document "### JPEG\n"
IMAGE_DIR=./testdata/benchmark/jpg
datasize $IMAGE_DIR
document "#### Result\n\n"

document "##### Parallel\n"
benchmark jpg 1 70 $IMAGE_DIR
benchmark jpg 3 70 $IMAGE_DIR
benchmark jpg 9 70 $IMAGE_DIR

document "\n\n##### Quality\n"
benchmark jpg 10 90 $IMAGE_DIR
benchmark jpg 10 70 $IMAGE_DIR
benchmark jpg 10 30 $IMAGE_DIR


document "\n\n### PNG\n"
IMAGE_DIR=./testdata/benchmark/png
datasize $IMAGE_DIR
document "#### Result\n\n"

document "##### Parallel\n"
benchmark png 1 70 $IMAGE_DIR
benchmark png 3 70 $IMAGE_DIR
benchmark png 9 70 $IMAGE_DIR

document "\n\n##### Quality\n"
benchmark png 10 90 $IMAGE_DIR
benchmark png 10 70 $IMAGE_DIR
benchmark png 10 30 $IMAGE_DIR
