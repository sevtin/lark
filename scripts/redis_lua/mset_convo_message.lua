for i=1,  #KEYS do
    if i==1 then
        redis.call('SET', KEYS[i], ARGV[i+1], "EX", ARGV[1])
    else
        redis.call('SET', KEYS[i], ARGV[i+1])
    end
end
return 0