#!/usr/bin/env bash

source ./proto_dir.cfg

length=${#all_proto}
for (( i=0; i <= $length; i++ )); do
  proto=${all_proto[$i]}
  if [ -z "$proto" ]; then
    continue
  fi
#  protoc -I ../../../ -I ./ --go_out=plugins=grpc:. $proto
  protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative $proto
  s=`echo $proto | sed 's/ //g'`
  v=${s//proto/pb.go}
  protoc-go-inject-tag -input=./$v
  echo "protoc --go_out=plugins=grpc:." $proto
done
echo "proto file generate success..."
