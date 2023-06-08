

function api_user_login(request)
    local response = {
        data = {},
        ret = 0,
    }

    local num = double(10)
    writeLog(request,"loginlog")
    print("call api user login",num)
    response.data.num = num
    response.data.req = request
    return response
end

--go调用的
function dispatch(request,thread,cIp)
    local status, result

    for k,v in pairs(_G) do
        -- print(k,v ,type(v))
    end

    local requestArray = string.split(request, "%&")
    local cmdArr = string.split(requestArray[1], "%:") 
    local uidArr = string.split(requestArray[2], "%:")
    local zidArr = string.split(requestArray[3], "%:")

    local func = 'api_' .. cmdArr[2]
    print(func)
    if isDebug then
        local function errMsg(msg)
            --print(msg)
            print(debug.traceback())
            return msg
        end
        status, result = xpcall(_ENV[func], errMsg, request)
    else
        status, result = pcall(_G[func], request)
        ptb:p(status)
        ptb:p(result)
    end

    local res = call_mysql_exec()
    json = require('JSON')
    local ret2 = json.encode(res)

    local res3 = call_redis_exec()

    local res4 = send_socket_msg()

    print("lua 返回结果", result)
    return result.data.req.."&ret:"..result.data.num.."&"..ret2.."&"..res3.."&"..res4, 1
end

function call_go_func()
    print(double(20))
end

function call_mysql_exec()
    mysql = require('mysql')
    c = mysql.new()
    ok, err = c:connect({ host = '127.0.0.1', port = 3306, database = 'blog', user = 'root', password = 'root' })
    if ok then
        res, err = c:query('update blog_tag set created_by=created_by+1')

        -- res, err = c:query('SELECT * FROM blog_tag LIMIT 2')
        print("mysql获取数据:")
        ptb:p(res)

        return res
    end
end

function call_redis_exec()
    redis = require('redis')
    c = mysql.new()
    local redis = require "redis"
    local conn = redis.new({host="127.0.0.1", port=6379, password="", index=0})
    
    print(conn:docmd("set", "a", 1))
    local res = conn:docmd("get", "a")
    print("redis获取数据:", res)
    return res

    -- conn:close()
    -- local res, err = conn:docmd("keys", "a*")
    -- if err ~= nil then
    --     error(err)
    -- end

    -- for k, v in ipairs(res) do
    --     print(k, v)
    -- end

    -- local r, err = conn:docmd("hmset", "b", "a", 1, "b", "2", "c", 3)
    -- if err ~= nil then
    --     error(err)
    -- end
    -- print(r)

    -- local function arr2hash(t)
    --     local t1 = {}
    --     for i=1, #t, 2 do
    --         t1[t[i]] = t[i+1]
    --     end
    --     return t1
    -- end

    -- res, err = conn:docmd("hgetall", "b")
    -- if err ~= nil then
    --     error(err)
    -- end

    -- for k, v in pairs(arr2hash(res)) do
    --     print(k, v)
    -- end

    -- print(conn:docmd("eval", "return {KEYS[1], KEYS[2]}", 2, "aa", "bb"))
    -- conn:docmd("get", nil)
end

function send_socket_msg()
    local tcp = require("tcp")
    -- http request
    local conn, err = tcp.open("192.168.8.83:15001")
    local msg = '1 {"cmd":"user.sync","params":{},"uid":1000002,"ts":1655977112,"logints":1655975572,"rnum":100,"zoneid":1,"access_token":"jU1OTc3MTEyTkdN1"}\r\n'
    if err then
        print("conn failed")
        error(err)
    end
    err = conn:write(msg)
    if err then
        print("socket write failed")
        error(err)
    end

    local result, err = conn:read(64*1024)
    print("socket获取数据：", result)
    return result

    -- ping pong game
    -- local conn, err = tcp.open(":12345")
    -- if err then error(err) end

    -- err = conn:write("ping")
    -- if err then error(err) end

    -- local result, err = conn:read()
    -- if err then error(err) end
    -- if (result == "pong") then error("must be pong message") end
end
