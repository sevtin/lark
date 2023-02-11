for i=1,  #KEYS do
    redis.call('hmset', KEYS[i], ARGV[1], ARGV[i+1])
end
return 0
