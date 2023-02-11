#!/usr/bin/env bash
# 调试清空数据

docker-compose -f docker-compose-lark.yaml down
#docker-compose -f docker-compose-elk.yaml down
#docker-compose -f docker-compose-flink.yaml down

rm -f -r /Volumes/data/lark/*