for i=1,  #KEYS do
    redis.call('hdel', KEYS[i], ARGV[i])
end
return 0