package allocation

import (
  "log"
  "live-deploy-client/schema"
)
func TaskFail(task *schema.Task, errMsg string)schema.TaskClientFinish{
  log.Println("task失败", errMsg)
  //记录完成任务到本地日志
  task.Status = 0
  task.Result = errMsg
  schema.AddTask(task)
  return schema.TaskClientFinish{
    ID: task.TaskID,
    Status: false,
    Result:errMsg,
  }
}