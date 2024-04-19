#!/bin/bash

exe_ver="1.0"
exe_name="LightSimulator"
build_date="$(date "+%Y-%m-%d %H:%M:%S")"

function DoBuild()
{
  export GOOS="${1}"
  export GOARCH="${2}"

  local out_suffix="_${1}"
  local out_ext=""
  if [ "${1}" == "windows" ]; then
    out_suffix=""
    out_ext=".exe"
  fi
  local out_filename="${exe_name}${out_suffix}${out_ext}"

  echo "Building ${out_filename} for ${1}/${2}"
  go build -ldflags "-X 'main.gVersion=${exe_ver}' -X 'main.gBuildDate=${build_date}'" -o "$out_filename" || exit 10
}

# Build binaries
[ -f "debug" ] && rm -f "debug"
DoBuild "windows" "amd64"
DoBuild "linux" "amd64"
DoBuild "darwin" "amd64"

echo "DONE!"
