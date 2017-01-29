#!/bin/sh

# ./build.sh v0.0.2
if [ $# -ne 1 ]; then
  echo "input version number"
  exit 1
fi

version=$1
echo "version: ${version}"

r_dir="releases/ohacan-${version}"
mkdir -p ${r_dir}

cd ohacan
for os in linux windows darwin
do
  arch="${os}_amd64"
  dir="ohacan_${arch}-${version}"
  echo "* build ${arch} ${dir}"
  mkdir -p $dir
  GOOS=$os GOARCH=amd64 go build -o ohacan main.go
  mv ohacan $dir
  tar czvf "${dir}.tar.gz" $dir 2> /dev/null
  rm -rf $dir
  mv "${dir}.tar.gz" ../${r_dir}
done

echo $r_dir
