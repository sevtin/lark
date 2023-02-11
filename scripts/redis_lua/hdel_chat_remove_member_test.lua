local KEYS = {"LK:CHAT_MEMBER"}
local ARGV = {1,2,3,4,5,6}

for i=1,  #ARGV do
    print(i,KEYS[1], ARGV[i])
end
return 0


--redis-cli -h 127.0.0.1 -p 63791 -a lark2022 script load "for i=1,  #KEYS do
--    if redis.call('hdel', KEYS[i], ARGV[i]) == 1 then
--        return 1
--    end
--end
--return 0"