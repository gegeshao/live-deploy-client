package default_task

import (
  "bytes"
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "os/exec"
	"fmt"
)

func execute(cwd string, command string, params ...string)(result string, success bool){
  cmd := exec.Command(command, params...)
  cmd.Path = cwd
  var errBuf bytes.Buffer
  var outBuf bytes.Buffer
  cmd.Stdout = &outBuf
  cmd.Stderr = &errBuf
  startErr := cmd.Start()
  waitErr := cmd.Wait()
  err := errBuf.String()
  out := outBuf.String()
  if err != ""{ result = result + err}
  if out != ""{ result = result + out}
  if startErr != nil || waitErr != nil {
    return result, false
  }
  return result, true
}

/**
{
id: xxx,
type: "Git"
action: clone
content: xxx
}
**/
func (maid *DefaultTaskMaid) Git(task *schema.Task)(string, bool){
	switch task.Action {
	case "clone":
		return clone(task)
	case "pull":
		return pull(task)
	case "checkout":
		return checkout(task)
	default:
		return fmt.Sprintf("undefine action type: %s", task.Action), false
	}

}

/**
{
id: xxx,
type: "clone"
action: clone
content: git url
}
**/
func clone(task *schema.Task)(string, bool){
  config:=utils.GetConfig()
  return execute(config.System.ProjectDir, "git", "clone", task.Content)
}
/**
{
id: xxx,
type: "pull"
action: pull
content: project name
}
**/
func pull(task *schema.Task)(string, bool){
	config:=utils.GetConfig()
	return execute(config.System.ProjectDir, "git", "pull", task.Content)
}

/**
{
id: xxx,
type: "pull"
action: pull
content: commit hash
}
**/
func checkout(task *schema.Task)(string, bool){
	config:=utils.GetConfig()
	return execute(config.System.ProjectDir, "git", "checkout", task.Content)
}