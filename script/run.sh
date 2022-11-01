#!/usr/bin/env bash

mkdir -p output
go build -o output/read-track ./cmd/read-track
./output/read-track
