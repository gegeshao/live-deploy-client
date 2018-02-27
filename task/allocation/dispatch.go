package allocation

import (
  "live-deploy-client/schema"
)

func Dispatch(task *schema.Task) schema.TaskClientFinish{
  /*判断是否为内置任务*/
  //内置任务
  if exist, status, result := DoDefalutTask(task); exist{
    return schema.TaskClientFinish{
      ID: task.TaskID,
      Status: status,
      Result:  result,
    }
  }else{
    return DoCustomTask(task)
  }
  //定制任务
}