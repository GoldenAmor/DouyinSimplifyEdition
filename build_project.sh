#!/usr/bin/env bash

for i in {user,video,favorite,comment,relation}
do
  cd cmd/$i/
  sh build.sh
  chmod +x output/bin/dousheng.$i
  cd -
done

mkdir output

for i in {user,video,favorite,comment,relation}
do
  mkdir output/$i
  cp -rf cmd/$i/output output/$i/output
done

cd cmd/api
go build main.go
cd -
mkdir output/api
cp cmd/api/main output/api/main
