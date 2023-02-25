if #KEYS+2~=#ARGV then
    return 0
end

for i=1,  #KEYS do
    redis.call('hmset', KEYS[i], ARGV[1], ARGV[i+2])
    redis.call('Expire', KEYS[i], ARGV[2])
end
return 0
