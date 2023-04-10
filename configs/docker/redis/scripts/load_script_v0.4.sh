#!/usr/bin/env bash

SERVER_HOST="10.0.115.14"
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

MYSQL_PORT=3307
MYSQL_DB="lark"
TRUNCATE_01="truncate table messages;"

mysql -h"$SERVER_HOST" -P"$MYSQL_PORT" -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -D${MYSQL_DB} -e "$TRUNCATE_01"

for (( i = 7001; i <= 7004; i++ )); do
    redis-cli -h $SERVER_HOST -p $i FLUSHDB
done

for ((i=7001; i<=7008; i++))
do
#if [ $i -le 7004 ]; then
#  redis-cli -h $SERVER_HOST -p $i --scan --pattern "LK:*" | xargs -L 1 redis-cli -h $SERVER_HOST -p $i del
#fi

redis-cli -h $SERVER_HOST -p $i script load "
if #KEYS~=1 or #ARGV~=1 then
    return 'PARAM_ERROR'
end

local key = KEYS[1]
local ex = ARGV[1]

if redis.call('EXISTS', key)==1 then
    return 'EXISTED'
else
    return redis.call('SET', key, 1, 'EX', ex)
end
"
done


#tail -f /dev/null