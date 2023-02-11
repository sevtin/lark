for i=1,  #KEYS do
    redis.call('hmset', KEYS[i], ARGV[i*2-1], ARGV[i*2])
end
return 0