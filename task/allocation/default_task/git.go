package default_task

import (
  "live-deploy-client/schema"
  "live-deploy-client/utils"
	"fmt"
  "path"
  "strings"
)


/**
{
id: xxx,
type: "Git"
action: clone
content: xxx
}
**/
func (maid *DefaultTaskMaid) Git(task *schema.Task)(bool, string){
	switch task.Action {
	case "clone":
		return gitClone(task)
	case "checkout":
		return gitCheckout(task)
	default:
		return false, fmt.Sprintf("Git undefine action type: %s", task.Action)
	}

}

/**
{
id: xxx,
type: "Git"
action: clone
content: git url
}
**/
func gitClone(task *schema.Task)(bool, string){
  config:=utils.GetConfig()
  return execute(config.System.ProjectDir, "git", "clone", task.Content)
}

func gitPull(projectName string)(bool, string){
	config:=utils.GetConfig()
	return execute(path.Join(config.System.ProjectDir, projectName), "git", "pull", "origin", "master")
}

/**
{
id: xxx,
type: "Git"
action: checkout
content: projectName, commit hash
}
**/
func gitCheckout(task *schema.Task)(bool, string){
	config:=utils.GetConfig()
	content := strings.Split(task.Content, ",")
	if len(content)!=2 {
	  return false, "content is illegal"
  }
	pullSuccess, pullResult := gitPull(content[0])
	if !pullSuccess{
	  return false,pullResult
  }
	return execute(path.Join(config.System.ProjectDir, content[0]), "git", "checkout", content[1])
}