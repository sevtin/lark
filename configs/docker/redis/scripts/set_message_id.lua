---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by saeipi.
--- DateTime: 2023/1/11 2:20 PM
---

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