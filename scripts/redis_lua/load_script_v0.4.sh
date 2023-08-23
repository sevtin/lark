#!/usr/bin/env bash

SERVER_HOST="10.0.105.38"
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

for (( i = 7001; i <= 7008; i++ )); do
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

redis-cli -h $SERVER_HOST -p $i script load "
local GetResult = {
    Success = 'Success',
    NotFound = 'NotFound',
    Failed = 'Failed'
}

local SetResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local DelResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local HSetnxResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local HGetResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local HDelResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local DecrResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local DecrbyResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local IncrbyResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local IncrResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local function Get(key)
    local result = redis.call('GET', key)
    if result == false then
        return GetResult.NotFound, 0
    elseif result == nil then
        return GetResult.Error, 0
    else
        if type(result) ~= 'string' then
            return GetResult.Error, 0
        end
        return GetResult.Success, tonumber(result)
    end
end

local function Set(key, value, expire)
    local ok, err = redis.call('SET', key, value)
    if not ok then
        return SetResult.Failed
    end

    if expire > 0 then
        redis.call('EXPIRE', key, expire)
    end

    return SetResult.Success
end

local function Del(key)
    local result = redis.call('DEL', key)
    -- 也可用 result >= 0 判断
    if type(result) == 'number' then
        return DelResult.Success
    else
        return DelResult.Failed
    end
end

local function HSetnx(key, field, value)
    local result = redis.call('HSETNX', key, field, value)
    if result == 1 then
        return HSetnxResult.Success
    else
        return HSetnxResult.Failed
    end
end

local function HGet(key, field)
    local result = redis.call('HGET', key, field)
    if result == false then
        return HGetResult.Failed, 0
    else
        return HGetResult.Success, tonumber(result)
    end
end

local function HDel(key, field)
    local result = redis.call('HDEL', key, field)
    if result == 1 then
        return HDelResult.Success
    else
        return HDelResult.Failed
    end
end

local function Decr(key)
    local result = redis.call('DECR', key)
    if type(result) ~= 'number' then
        return DecrResult.Failed, 0
    end
    return DecrResult.Success, tonumber(result)
end

local function Decrby(key, decrement)
    local result = redis.call('DECRBY', key, decrement)
    if result then
        return DecrbyResult.Success, tonumber(result)
    else
        return DecrbyResult.Failed, 0
    end
end

local function Incrby(key, decrement)
    local result = redis.call('INCRBY', key, decrement)
    if result then
        return IncrbyResult.Success, tonumber(result)
    else
        return IncrbyResult.Failed, 0
    end
end

local function Incr(key)
    local result = redis.call('INCR', key)
    if type(result) ~= 'number' then
        return IncrResult.Failed, 0
    end
    return IncrResult.Success,tonumber(result)
end

-- 回滚剩余红包数量
local function rollbackRemainQuantity(key)
    return Incr(key)
end

-- 回滚剩余红包金额
local function rollbackRemainAmount(key, value)
    return Set(key, value)
end

-- 删除红包领取记录
local function rollbackRecord(key, field)
    return HDel(key, field)
end

-- amount:剩余红包金额
-- num:剩余红包个数
local function  getRandomAmount(amount, num)
    if num == 1 or amount == 0 then
        return amount
    end

    local minAmount = 1
    local maxAmount = math.floor(amount/num)*2 - minAmount

    if maxAmount <= 0 then
        maxAmount = minAmount
    end

    return math.random(minAmount, maxAmount)
end

local function response(status, ...)
    return status .. ':' .. table.concat({...}, ':')
end

local red_env_key = KEYS[1]
local uid = ARGV[1]
local remain_quantity_key = 'LK:RED_ENV:REMAIN_QUANTITY:'..red_env_key
local remain_amount_key = 'LK:RED_ENV:REMAIN_AMOUNT:'..red_env_key
local status_key = 'LK:RED_ENV:STATUS:'..red_env_key
local record_key = 'LK:RED_ENV:RECORD:'..red_env_key
local received_status = 2

-- 成功/失败:描述:剩余红包数量:剩余红包金额:本次发放金额
-- 1、获取红包状态
local result, status = Get(status_key)
if result ~= GetResult.Success then
    return response('FAILED','GET_STATUS_FAILED',0,0,0)
end

-- 红包已领完
if status == received_status then
    return response('FAILED','RED_ENV_OVER',0,0,0)
end

-- 2、是否已经领过红包
local result, received_amount = HGet(record_key, uid)
if result == HGetResult.Success then
    -- 已经领取
    return response('FAILED','RED_ENV_RECEIVED',0,0,0)
end

-- 3、获取红包余额
local result, remain_amount = Get(remain_amount_key)
if result ~= GetResult.Success then
    return response('FAILED','GET_REMAIN_AMOUNT_FAILED',0,0,0)
end

-- 红包余额异常
if remain_amount <= 0 then
    return response('FAILED','REMAIN_AMOUNT_ERROR',0,remain_amount,0)
end

-- 4、剩余红包数量
local result, remain_quantity = Decr(remain_quantity_key)
if result == DecrResult.Failed then
    return response('FAILED','DECR_REMAIN_QUANTITY_FAILED',0,0,0)
end

-- 红包数量异常
if remain_quantity < 0 then
    return response('FAILED','REMAIN_QUANTITY_ERROR',remain_quantity,0,0)
end

-- 5、更新红包余额
local red_env_amount = getRandomAmount(remain_amount, remain_quantity + 1)
local new_remain_amount = remain_amount - red_env_amount
local result = Set(remain_amount_key, new_remain_amount, 0)
if result == SetResult.Failed then
    rollbackRemainQuantity(remain_quantity_key)
    return response('FAILED','SET_REMAIN_AMOUNT_FAILED',remain_quantity+1,remain_amount,0)
end

-- 6、记录红包领取人
local result = HSetnx(record_key, uid, red_env_amount)
if result == HSetnxResult.Failed then
    rollbackRemainQuantity(remain_quantity_key)
    rollbackRemainAmount(remain_amount_key, remain_amount)
    return response('FAILED','HSETNX_RECORD_FAILED',remain_quantity+1,remain_amount,0)
end

if remain_quantity == 0 then
    -- 5、红包已领完
    local result = Set(status_key, received_status, 0)
    if result == SetResult.Failed then
        rollbackRemainQuantity(remain_quantity_key)
        rollbackRemainAmount(remain_amount_key, remain_amount)
        rollbackRecord(record_key, uid)
        return response('FAILED','SET_RED_ENV_STATUS_FAILED',remain_quantity+1,remain_amount,0)
    end
end

return response('SUCCEED','RECEIVED',remain_quantity,new_remain_amount,red_env_amount)
"

redis-cli -h $SERVER_HOST -p $i script load "
local IncrbyResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local IncrResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local SetResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local HDelResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local function Incrby(key, decrement)
    local result = redis.call('INCRBY', key, decrement)
    if result then
        return IncrbyResult.Success, tonumber(result)
    else
        return IncrbyResult.Failed, 0
    end
end

local function Incr(key)
    local result = redis.call('INCR', key)
    if type(result) ~= 'number' then
        return IncrResult.Failed, 0
    end
    return IncrResult.Success,tonumber(result)
end

local function Set(key, value, expire)
    local ok, err = redis.call('SET', key, value)
    if not ok then
        return SetResult.Failed
    end

    if expire > 0 then
        redis.call('EXPIRE', key, expire)
    end

    return SetResult.Success
end

local function HDel(key, field)
    local result = redis.call('HDEL', key, field)
    if result == 1 then
        return HDelResult.Success
    else
        return HDelResult.Failed
    end
end

local function response(status, ...)
    return status .. ':' .. table.concat({...}, ':')
end

local red_env_key = KEYS[1]
local uid = ARGV[1]
local dist_amount = ARGV[2]
local remain_quantity_key = 'LK:RED_ENV:REMAIN_QUANTITY:'..red_env_key
local remain_amount_key = 'LK:RED_ENV:REMAIN_AMOUNT:'..red_env_key
local status_key = 'LK:RED_ENV:STATUS:'..red_env_key
local record_key = 'LK:RED_ENV:RECORD:'..red_env_key
local issued_status = 1

-- 1、还原红包余额
local result, remain_amount = Incrby(remain_amount_key, tonumber(dist_amount))
if result == IncrbyResult.Failed then
    return response('FAILED','INCRBY_REMAIN_AMOUNT_FAILED')
end

-- 2、还原红包数量
local result, remain_quantity = Incr(remain_quantity_key)
if result == IncrResult.Failed then
    return response('FAILED','INCR_REMAIN_QUANTITY_FAILED')
end

-- 3、还原红包状态
local result = Set(status_key, issued_status, 86400)
if result == SetResult.Failed then
    return response('FAILED','SET_STATUS_FAILED')
end

-- 4、删除领取记录
local result = HDel(record_key, uid)
if result == HDelResult.Failed then
    return response('FAILED','HDEL_RECORD_FAILED')
end

return response('SUCCEED','ROLLBACK_SUCCEEDED')
"

done


#tail -f /dev/null