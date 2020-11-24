#!/bin/bash

RUN_NAME="nilmusic.service"

mkdir -p output/bin output/conf
find conf/ -type f ! -name "*_local.*" | xargs -I{} cp {} output/conf/

go build -o output/bin/${RUN_NAME}
