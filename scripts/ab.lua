--[[
    nginx 模块
    模块必须实现一下两个方法， deploy,  undeploy
]]--
gosystem = require('gosystem')
ab = {}

local function writeFile(filepath, filecontent)
    file = io.open(filepath, "w+")
    if(io.type(file) == nil) then error("没有权限打开:"..filepath) end
    file:write(filecontent)
    file:close()
end

function ab.deploy(trackID, trackKey, content)
    config = gosystem.getConfig()
    nginxConfigPath = gosystem.path().join(config["nginx_config_path"], "ab-id-"..trackID .. ".conf")
    status, errorInfo = pcall(writeFile, nginxConfigPath, content)
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
        result= msg.."\n"..result
    }
end

function ab.undeploy(trackID, trackKey, content)
    config = gosystem.getConfig()
    nginxConfigPath = gosystem.path().join(config["nginx_config_path"], "ab-id-"..trackID .. ".conf")
    os.remove(nginxConfigPath)
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
        result= msg.."\n"..result
    }
end