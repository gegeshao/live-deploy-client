# live-deploy-client

===============================
## 编译

配置好go的工作目录

下载安装 dep: https://github.com/golang/dep
下载安装 gox: https://github.com/mitchellh/gox

```
dep ensure
sh build.sh
```

编译好的文件就在bin目录下

## 内置任务

内置了一些可能需要的任务： UpdateScripts 等


### UpdateScripts

开启该功能需要在 config.yaml中配置`load_default_task`,使客户端能够 禁用/开启 开功能。

下发任务内容

```
{
    id: xxx,
    type: "UpdateScripts"
    action: filename
    content: download-url
}
```

## 任务脚本

任务脚本采用lua 编写, 内置 `gosystem` 模块

### 规范

每个任务脚本必须使用  `任务类型.lua`的形式命名.
任务脚本：

以任务类型nginx举例

```lua
nginx = {} -- nginx变量名称为 任务类型

function nginx.deploy  --任务动作 挂载到 全局变量上
    -- body
    return {
        status=bool, -- 任务是否执行成功
        result=string -- 任务执行结果 这个可以为空字符串, 用于反馈用户操作
    }
    -- 每个任务返回值必须为table
end

--[[

or 其他的任务

function nginx.undeploy
    -- body
     return {
        status=bool,
        result=string
     }
end

...
]]--

```

### gosystem

gosystem包含以下方法


#### execute(command, arg...)

用于执行系统命令, 如  `gosystem.execute("nginx". "-t")`

#### path

返回一个table 包含 如下函数：

##### join

等价与 go的 `path.Join()`. 使用如下

```
path = gosystem.path()

path.join(arg1, arg2...)
```






