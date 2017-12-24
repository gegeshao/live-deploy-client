--[[
    nginx 模块
    模块必须实现一下两个方法， deploy,  undeploy
]]--
gosystem = require('gosystem')
nginx = {}

local function writeFile(filepath, filecontent)
    file = io.open(filepath, "w+")
    if(io.type(file) == nil) then error("没有权限打开:"..filepath) end
    file:write(filecontent)
    file:close()
end

function nginx.deploy(trackID, trackKey, content)
    config = gosystem.getConfig()
    nginxConfigPath = gosystem.path().join(config["nginx_config_path"], "id-"..trackID .. ".conf")
    status, errorInfo = pcall(writeFile, nginxConfigPath, content)
    print(status, errorInfo)
    if(not status) then return {status = false, result = errorInfo} end
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