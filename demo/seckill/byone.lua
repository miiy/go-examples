--
-- Created by IntelliJ IDEA.
-- To change this template use File | Settings | File Templates.
--

-- redis lua脚本：eval script numkeys key [key ...] arg [arg ...]
--
-- 127.0.0.1:6379> eval "return redis.call('set', KEYS[1], KEYS[2])" 2 foo bar
-- OK
-- 127.0.0.1:6379> get foo
-- "bar"
--
-- 127.0.0.1:6379> eval "return redis.call('set', KEYS[1], ARGV[1])" 1 foo bar
-- OK
-- 127.0.0.1:6379> get foo
-- "bar"

-- set
--
-- 127.0.0.1:6379> sadd tset 1
-- (integer) 1
-- 127.0.0.1:6379> sadd tset 1
-- (integer) 0
--
-- list
--
-- 127.0.0.1:6379> lpush tlist 1
-- (integer) 1
-- 127.0.0.1:6379> lpop tlist
-- "1"
-- 127.0.0.1:6379> lpop tlist
-- (nil)





local stockKey  = KEYS[1]
local buyersKey = KEYS[2]
local uid       = KEYS[3]

local result = redis.call("sadd", buyersKey, uid)
if (tonumber(result) == 1)
then
    local stock = redis.call("lpop", stockKey)
    if stock
    then
        return 1
    else
        return -1
    end
else
    return -3
end


