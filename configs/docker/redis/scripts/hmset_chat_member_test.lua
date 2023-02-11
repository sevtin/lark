local KEYS = {"LK:CHAT_MEMBER","LK:CHAT_MEMBER_INFO","LK:CHAT_MEMBER_INFO","LK:CHAT_MEMBER_INFO","LK:CHAT_MEMBER_INFO","LK:CHAT_MEMBER_INFO"}
local ARGV = {1,"10000,1,0",2,"{\"uid\":2}",3,"{\"uid\":3}",4,"{\"uid\":4}",5,"{\"uid\":5}",6,"{\"uid\":6}"}

for i=1,  #KEYS do
    print(i,KEYS[i], ARGV[i*2-1], ARGV[i*2])
end
return 0

--redis-cli -h 127.0.0.1 -p 63791 -a lark2022 script load "
--for i=1,  #KEYS do
--    redis.call('hmset', KEYS[i], ARGV[i*2-1], ARGV[i*2])
--end
--return 0
--"