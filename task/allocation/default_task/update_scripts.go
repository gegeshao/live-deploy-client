package default_task

import (
  "fmt"
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "live-deploy-client/vm"
  "log"
  "strings"
)


/*
{
    id: xxx,
    type: "UpdateScripts"
    action: filename
    content: download-url
}
*/
func (maid *DefaultTaskMaid) UpdateScripts(task *schema.Task)(bool, string){
  filename:=task.Action
  url:=task.Content
  //保证安全
  filename = strings.Replace(filename, "/", "_", -1)
  if err:= utils.Download(url, "scripts/"+filename); err!=nil{
    log.Println(err)
    return false, err.Error()
  }
  //加载新的脚本
  if err := vm.LoadScript(filename); err!=nil{
    return false, fmt.Sprintf("加载%s脚本失败: %v", filename, err)
  }
  return true, filename
}




