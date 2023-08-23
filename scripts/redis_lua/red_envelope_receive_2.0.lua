local SetResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local GetResult = {
    Success = 'Success',
    NotFound = 'NotFound',
    Failed = 'Failed'
}

local DelResult = {
    Success = 'Success',
    Failed = 'Failed'
}

local DecrResult = {
    Success = 'Success',
    Failed = 'Failed'
}

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

local function Del(key)
    local result = redis.call('DEL', key)
    -- 也可用 result >= 0 判断
    if type(result) == 'number' then
        return DelResult.Success
    else
        return DelResult.Failed
    end
end

local function Decr(key)
    local result = redis.call('DECR', key)
    if type(result) ~= 'number' then
        return DecrResult.Failed,0
    end
    return DecrResult.Success,tonumber(result)
end

-- 回滚剩余红包数量
local function rollbackRemainQuantity(key)
    return redis.call('INCR', key)
end

-- 回滚剩余红包金额
local function rollbackRemainAmount(key, value)
    return redis.call('SET', key, value)
end

-- 删除红包领取记录
local function rollbackRecord(key)
    return redis.call('DEL', key)
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
local record_key = 'LK:RED_ENV:RECORD:'..red_env_key..':'..uid
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
local result, received_amount = Get(record_key)
if result == GetResult.Failed then
    return response('FAILED','GET_RECEIVED_AMOUNT_FAILED',0,0,0)
end

-- 已经领取
if received_amount > 0 then
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
local result = Set(record_key, red_env_amount,86400)
if result == SetResult.Failed then
    rollbackRemainQuantity(remain_quantity_key)
    rollbackRemainAmount(remain_amount_key, remain_amount)
    return response('FAILED','SET_RECORD_FAILED',remain_quantity+1,remain_amount,0)
end

if remain_quantity == 0 then
    -- 5、红包已领完
    local result = Set(status_key, received_status, 0)
    if result == SetResult.Failed then
        rollbackRemainQuantity(remain_quantity_key)
        rollbackRemainAmount(remain_amount_key, remain_amount)
        rollbackRecord(record_key)
        return response('FAILED','SET_RED_ENV_STATUS_FAILED',remain_quantity+1,remain_amount,0)
    end
end

return response('SUCCEED','RECEIVED',remain_quantity,new_remain_amount,red_env_amount)