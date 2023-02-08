package.path = "/Users/ray/Documents/GO_PATH/src/gsnippet/?.lua;" .. package.path
package.cpath = "/Users/ray/Documents/GO_PATH/src/gsnippet/?.so;" .. package.cpath

-- APP_PATH = debug.getinfo(1).short_src
-- APP_PATH = string.sub(APP_PATH, 0, -13)
-- print(APP_PATH)
-- package.path = APP_PATH .. "/?.lua;" .. package.path

require "lualib.string"
require "lualib.func"
-- require "lib.userobjs"
-- require "lib.allianceobjs"
-- require "lualib.game"
-- require "lib.mergezone"

ptb = require "lualib.ptb"
-- json = require "cjson.safe"
-- warn = require "lib.warn"
-- CONFIGCFG = require "config.configCfg"
-- PLATFORM = getConfig('baseCfg.APPPLATFORM')
-- BIGPALTID = getConfig('baseCfg.AppPlatid')
-- formula = require "lib.formula"

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

    print("lua 返回结果", result)
    return result.data.req.."&ret:"..result.data.num, 1
end

function call_go_func()
    print(double(20))
end


