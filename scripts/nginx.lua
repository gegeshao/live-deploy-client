--[[
    nginx 模块
    模块必须实现一下两个方法， deploy,  undeploy
]]--
gosystem = require('gosystem')
nginx = {}
function nginx.deploy(trackID, trackKey, Content)
    ok, msg = gosystem.execute("nginx", "-t")
    if not ok then
    return {
        status = ok,
        result=msg
    }
    end
    reloadOk, result = gosystem.execute("nginx", "-s", "reload")
    return {
        status = reloadOk,
        result=result
    }
end

function nginx.undeploy(trackID, trackKey)
end