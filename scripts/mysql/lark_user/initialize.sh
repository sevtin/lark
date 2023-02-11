#!/bin/bash

MYSQL_USER=${MYSQL_USER:-root}
MYSQL_PASSWORD=${MYSQL_PASSWORD:-lark2022}
MYSQL_HOST="127.0.0.1"
MYSQL_PORT=3306
MYSQL_DB="lark_user"
SCRIPT_PATH=$(cd $(dirname $0);pwd)
SQL_FILE="/sqls"

folder=$SCRIPT_PATH$SQL_FILE
for file in ${folder}/*
do
  mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -D${MYSQL_DB} < ${file}
done
