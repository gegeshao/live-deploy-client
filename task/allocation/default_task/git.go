package default_task

import (
  "live-deploy-client/schema"
  "os/exec"
)
/**
{
id: xxx,
type: "GitClone"
action: clone
content: git url
}
**/
func (maid *DefaultTaskMaid) GitClone(task *schema.Task)(bool, string){
  cmd := exec.Command("git" ,"clone", task.Content)
  cmd.Path = cwd
}
