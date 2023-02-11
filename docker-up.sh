#!/usr/bin/env bash
# 初始化执行一次

rm -rf ./configs/docker/mysql/*
rm -rf ./configs/docker/redis/*

mkdir -p ./configs/docker/redis/scripts

cp -Rp ./scripts/mysql/* ./configs/docker/mysql/
cp -Rp ./scripts/redis_lua/* ./configs/docker/redis/scripts/

#docker-compose -f docker-compose-elk.yaml up -d
#docker-compose -f docker-compose-flink.yaml up -d
docker-compose -f docker-compose-lark.yaml up -d