#!/usr/bin/env bash

set -e
mkdir -p output
go build -o output/read-track .
./output/read-track -conf=./
