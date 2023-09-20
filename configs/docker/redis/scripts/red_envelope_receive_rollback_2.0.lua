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

local IncrResult = {
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

local function Incr(key)
    local result = redis.call('INCR', key)
    if type(result) ~= 'number' then
        return IncrResult.Failed,0
    end
    return IncrResult.Success,tonumber(result)
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
local record_key = 'LK:RED_ENV:RECORD:'..red_env_key..':'..uid
local issued_status = 1

-- 1、获取红包余额
local result, remain_amount = Get(remain_amount_key)
if result ~= GetResult.Success then
    return response('FAILED','GET_REMAIN_AMOUNT_FAILED')
end

-- 2、更新红包余额
local result = Set(remain_amount_key, remain_amount + tonumber(dist_amount), 86400)
if result ~= SetResult.Success then
    return response('FAILED','SET_REMAIN_AMOUNT_FAILED')
end

-- 3、还原红包数量
local result,remain_quantity = Incr(remain_quantity_key)
if result ~= IncrResult.Success then
    return response('FAILED','INCR_REMAIN_QUANTITY_FAILED')
end

-- 4、还原红包状态
local result = Set(status_key, issued_status, 86400)
if result ~= SetResult.Success then
    return response('FAILED','SET_STATUS_FAILED')
end

-- 5、删除领取记录
local result = Del(record_key)
if result ~= DelResult.Success then
    return response('FAILED','DEL_RECORD_FAILED')
end

return response('SUCCEED','ROLLBACK_SUCCEEDED')