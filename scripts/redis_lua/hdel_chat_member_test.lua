local KEYS = {"LK:CHAT:MEMBERS_HASH:3333336666669999990",
              "LK:CHAT:MEMBERS_HASH:3333336666669999990",
              "LK:CHAT:MEMBER_INFO:3333336666669999990",
              "LK:CHAT:MEMBER_INFO:3333336666669999990"}
local ARGV = {1,2,1,2 }

for i=1,  #KEYS do
    print(i,KEYS[i], ARGV[i])
end
return 0

--redis-cli -h 127.0.0.1 -p 63791 -a lark2022 script load "
--for i=1,  #KEYS do
--    redis.call('hdel', KEYS[i], ARGV[i])
--end
--return 0
--"