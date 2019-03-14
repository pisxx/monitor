#!/bin/bash

export GOOS=linux
pkg=$(echo $1 |cut -d . -f 1)
echo -e "Building $pkg binary for Linux"
go build -o $pkg-linux $1