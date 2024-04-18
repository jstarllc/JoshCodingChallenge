#!/bin/bash

build_date="$(date "+%Y-%m-%d %H:%M:%S")"
exe_name="josh-coding-challenge"

function DoBuild()
{
export GOOS="${1}"
export GOARCH="${2}"

local out_suffix="-${3}"
local out_ext=""
if [ "${1}" == "windows" ]; then
  out_suffix=""
  out_ext=".exe"
fi
local out_filename="${exe_name}${out_suffix}${out_ext}"

echo "Building ${out_filename} for ${1}/${2}"
go build -ldflags "-X 'main.buildDate=${build_date}'" -o "${out_filename}"
}

if [ -f "debug" ]; then
  rm "debug"
fi

DoBuild "darwin" "amd64" "osx-x64"
DoBuild "linux" "arm" "linux-armhf"

echo "DONE!"
