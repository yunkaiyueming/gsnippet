function writeLog(message,path)
    path = path or ''
    message = message or ''
    local date = os.date('%Y%m%d')
    local fileName = "/Users/ray/Documents/GO_PATH/src/gsnippet/lualib/"..path..date .. '.log'

    message = message .. '\r\n'
    local now   = os.date('%Y-%m-%d %H:%M:%S')
    message = now ..': '..message
    local f = io.open(fileName, "a+")
    if f then
        f:write(message)
        f:close()
    end
end

-- 打乱数组
function table.rand(arr)

    local arr_size=#arr
    local tmp_arr={}

    setRandSeed()

    for i=1,arr_size do
        local rd=rand(1,arr_size+1-i)
        table.insert(tmp_arr,arr[rd])
        table.remove(arr,rd)
    end

    return tmp_arr
end

function setRandSeed()
    if _GAMEVARS.isseed == 0 then
        local socket = require("socket.core")
        math.randomseed( socket.gettime()*1000 )
        _GAMEVARS.isseed = 1
    end
end

function rand(m,n)
    setRandSeed()
    math.random(m,n); math.random(m,n); math.random(m,n)
    return math.random(m,n)
end

-- n > m
function randByNum(m,n,num,noNum)
    local res = {}

    local result = {}
    for k=m,n do
        if tonumber(k) ~= tonumber(noNum) then
            table.insert(result,k)
        end
    end

    for i=1,num do
        local length = table.length(result)
        if tonumber(length) >0 then
            local tmpKey = rand(1,length)
            table.insert(res,result[tmpKey])
            table.remove(result,tmpKey)
        end
    end

    return res
end

-- table 长度
function table.length(array)
    local len = 0
    if type(array) == 'table' then
        for _ in pairs(array) do
            len = len + 1
        end
    end

    return len
end

-- 取table的子数组
-- args:
-- arr 数组（一定是无索引table）
-- i 起始位
-- n 长度
function table.subArray(arr, i, n)
    local res = {}
    for j = i, i + n-1 do
        if j > #arr then
            break
        end
        table.insert(res, arr[j])
    end
    return res
end


-- 获取当前日期时间戳
function getTimeZoneByDate(dtable)
    local zone = getConfig('baseCfg.TIMEZONE') or 0

    local now = os.time()
    local diff = os.difftime(os.time(os.date("!*t", now)),now)

    local restime = os.time(dtable) - zone * 3600 - diff

    return restime
end