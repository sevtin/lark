#!/usr/bin/env bash

nohup redis-server --requirepass lark2022 --appendonly yes &
sleep 30

redis-cli -h 127.0.0.1 -p 6379 -a lark2022 script load "
for i=1,  #KEYS do
    redis.call('hdel', KEYS[i], ARGV[i])
end
return 0
"

redis-cli -h 127.0.0.1 -p 6379 -a lark2022 script load "
for i=1,  #KEYS do
    redis.call('hmset', KEYS[i], ARGV[i*2-1], ARGV[i*2])
end
return 0
"

redis-cli -h 127.0.0.1 -p 6379 -a lark2022 script load "
for i=1,  #KEYS do
    redis.call('hmset', KEYS[i], ARGV[1], ARGV[i+1])
end
return 0
"

redis-cli -h 127.0.0.1 -p 6379 -a lark2022 script load "
for i=1,  #KEYS do
    if i==1 then
        redis.call('SET', KEYS[i], ARGV[i+1], \"EX\", ARGV[1])
    else
        redis.call('SET', KEYS[i], ARGV[i+1])
    end
end
return 0
"

redis-cli -h 127.0.0.1 -p 6379 -a lark2022 script load "
for i=1,  #KEYS do
    redis.call('SET', KEYS[i], ARGV[i+1], \"EX\", ARGV[1])
end
return 0
"

redis-cli -h 127.0.0.1 -p 6379 -a lark2022 script load "
if #ARGV~=#KEYS*2 then
    return 0
end

local k = 1
for i=1,  #KEYS do
    redis.call('SET', KEYS[i], ARGV[k], \"EX\", ARGV[k+1])
    k = k+2
end
return 0
"

redis-cli -h 127.0.0.1 -p 6379 -a lark2022 script load "
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

redis-cli -h 127.0.0.1 -p 6379 -a lark2022 script load "
local prefix = KEYS[1]
local chatId = KEYS[2]
local timestamp = tonumber(KEYS[3])
local key = \"\"
for i=1, #ARGV do
    key = string.format(\"%s%d\",prefix,ARGV[i])
    redis.call('ZADD', key, timestamp, chatId)
end
return 0
"

tail -f /dev/null