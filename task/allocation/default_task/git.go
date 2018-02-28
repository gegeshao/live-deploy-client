package default_task

import (
  "bytes"
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "os/exec"
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
type: "GitClone"
action: clone
content: git url
}
**/
func (maid *DefaultTaskMaid) GitClone(task *schema.Task)(string, bool){
  config:=utils.GetConfig()
  execute(config.System.ProjectDir, "git", "clone", task.Content)
}
