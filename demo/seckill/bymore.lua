--
-- Created by IntelliJ IDEA.
-- To change this template use File | Settings | File Templates.
--

-- decrby
-- 127.0.0.1:6379> set tdecrby 3
-- OK
-- 127.0.0.1:6379> decrby tdecrby 2
-- (integer) 1
-- 127.0.0.1:6379> decrby tdecrby 1
-- (integer) 0
-- 127.0.0.1:6379> decrby tdecrby 1
-- (integer) -1

local key = KEYS[1]
local val = ARGV[1]

local stock = redis.call("GET", key)
if (tonumber(stock) <= 0)
then
    return -1
else
    local descrstock = redis.call("DESCRBY", key, val)
    if (tonumber(descrstock) >= 0)
    then
        return decrstock
    else
        redis.call("INCRBY", key, val)
        return -2
    end
end