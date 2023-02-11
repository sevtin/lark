#!/usr/bin/env bash

source ./cfg/services.cfg
source ./cfg/docker-composes.cfg

rm -f -r ../build
mkdir -p ../build/bin

cp -Rp run ../build/run
cp -Rp ../configs ../build/run/configs
cp -Rp cfg ../build/cfg

LARK_APP=$(dirname "$PWD")
BUILD_BIN=${LARK_APP}"/build/bin"

length=${#composes}
for (( i=0; i <= $length; i++ )); do
  compose=${composes[$i]}
  if [ -z ${compose} ];then
      continue
  fi
  cp ${LARK_APP}/${compose} ../build/run
done
echo "copy success..."