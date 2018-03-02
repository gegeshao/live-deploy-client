package default_task

import (
  "live-deploy-client/schema"
  "live-deploy-client/utils"
	"fmt"
  "path"
  "strings"
)

/*
*
*/
func (maid *DefaultTaskMaid) PM2(task *schema.Task)(bool, string){
	switch task.Action {
  case "start":
    return start(task)
	case "restart":
		return restart(task)
	default:
		return false, fmt.Sprintf("PM2 undefine action type: %s", task.Action)
	}
}

/**
id: xxx,
type: "PM2"
action: restart
content: projectname
*/

func restart(task *schema.Task)(bool, string){
  config:=utils.GetConfig()
  projectName := task.Content
  command := []string{"restart", projectName}
  return execute(path.Join(config.System.ProjectDir,task.Content), "pm2", command...)
}

/**
id: xxx,
type: "PM2"
action: start
content: projectName, filepath
*/

func start(task *schema.Task)(bool, string){
  config:=utils.GetConfig()
  content := strings.Split(task.Content, ",")
  if len(content)!=2 {
    return false, "content is illegal"
  }
  projectName := content[0]
  filepath := content[1]
  command:= []string{"start", filepath, "--name", projectName}
  return execute(path.Join(config.System.ProjectDir,projectName), "pm2", command...)
}