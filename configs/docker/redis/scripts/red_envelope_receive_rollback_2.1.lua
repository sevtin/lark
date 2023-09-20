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