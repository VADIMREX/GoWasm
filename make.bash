#!/usr/bin/env bash
# check cloned go lang repo
if [ -d "go" ]; then
  cd go
  git fetch
  git pull
  cd ..
else
  git clone https://github.com/golang/go.git  
fi
# check compiled go lang
cd go
if [ -d "bin" ]; then
  echo 1
else
  cd src
  GOOS=js GOARCH=wasm GOROOT_FINAL=/usr/local/go ./make.bash
fi