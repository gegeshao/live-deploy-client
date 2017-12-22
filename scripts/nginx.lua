--[[
    nginx 模块
    模块必须实现一下两个方法， deploy,  undeploy
]]--

nginx = {}
local function nginxTest()
    local testOk = os.execute("nginx -t")
    local msg = ""

    if (testOk == nil) then testOk = false end

    if testOk == 0 then teskOk = true
        elseif testOk == 1 then testOk = false
    end

    if testOk == false then msg = "nginx 测试失败" else msg = "nginx 测试成功" end
    return testOk, msg
end

local function nginxReload()
    local testOk, msg = nginxTest()
    if not testOk then return {status=testOk, result=msg} end

    local reloadOk = os.execute("nginx -s reload")
    local reloadMsg = ""

    if (reloadOk == nil) then reloadOk = false end

    if reloadOk == 0 then reloadOk = true
    elseif reloadOk == 1 then reloadOk = false
    end

    if reloadOk == false then reloadMsg = "nginx reload失败" else reloadMsg = "nginx reload成功" end
    return {
        status=reloadOk,
        result=reloadMsg
    }
end

function nginx.deploy(trackID, trackKey, Content)
    return nginxReload()
end

function nginx.undeploy(trackID, trackKey)
end

return nginx