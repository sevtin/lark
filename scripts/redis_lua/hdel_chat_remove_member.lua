--- 弃用
for i=1,  #ARGV do
    if redis.call('hdel', KEYS[1], ARGV[i]) == 1 then
        return 1
    end
end
return 0