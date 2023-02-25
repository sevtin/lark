#!/usr/bin/env bash

SERVER_HOST="10.0.117.113"

redis-cli -h $SERVER_HOST -p 63791 -a lark2022 -n 0 --scan --pattern "LK:*" | xargs -L 5000 redis-cli -h $SERVER_HOST -p 63791 -a lark2022 -n 0 DEL


MYSQL_PORT=13306
MYSQL_USER=${MYSQL_USER:-root}
MYSQL_PASSWORD=${MYSQL_PASSWORD:-lark2022}
MYSQL_DB="lark_user"

TRUNCATE_01="truncate table users;"
TRUNCATE_02="truncate table chats;"
TRUNCATE_03="truncate table chat_members;"
TRUNCATE_04="truncate table chat_invites;"
TRUNCATE_05="truncate table avatars;"

mysql -h"$SERVER_HOST" -P"$MYSQL_PORT" -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -D${MYSQL_DB} -e "$TRUNCATE_01 $TRUNCATE_02 $TRUNCATE_03 $TRUNCATE_04 $TRUNCATE_05"

MYSQL_PORT=13307
MYSQL_DB="lark_chat_msg"
TRUNCATE_01="truncate table messages;"

mysql -h"$SERVER_HOST" -P"$MYSQL_PORT" -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -D${MYSQL_DB} -e "$TRUNCATE_01"