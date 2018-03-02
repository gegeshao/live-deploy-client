package allocation

import (
  "live-deploy-client/schema"
  "log"
  "testing"
)



func TestDoGitClone(t *testing.T) {
  finishTask := Dispatch(&schema.Task{
    Type: "Git",
    Action:"clone",
    Content: "git@github.com:huyinghuan/aliyun-voice.git",
  })
  if !finishTask.Status{
    log.Println(finishTask.Result)
    t.Fail()
  }else{
    log.Println(finishTask.Result)
  }
}


func TestDOGitCheckout(t *testing.T){
  finishTask := Dispatch(&schema.Task{
    Type: "Git",
    Action:"checkout",
    Content: "aliyun-voice,9faab911",
  })
  if !finishTask.Status{
    log.Println(finishTask.Result)
    t.Fail()
  }else{
    log.Println(finishTask.Result)
  }
}

func TestDOPM2Start(t *testing.T){
  finishTask := Dispatch(&schema.Task{
    Type: "PM2",
    Action:"start",
    Content: "test,index.js",
  })
  if !finishTask.Status{
    log.Println(finishTask.Result)
    t.Fail()
  }else{
    log.Println(finishTask.Result)
  }
}

func TestDOPM2Restart(t *testing.T){
  finishTask := Dispatch(&schema.Task{
    Type: "PM2",
    Action:"restart",
    Content: "test",
  })
  if !finishTask.Status{
    log.Println(finishTask.Result)
    t.Fail()
  }else{
    log.Println(finishTask.Result)
  }
}

//测试不存在的脚本任务
func TestDonotExistTask(t *testing.T){
  finishTask := Dispatch(&schema.Task{
    Type: "PM3",
    Action:"start",
  })
  if finishTask.Status != false{
    t.Fail()
  }
  finishTask = Dispatch(&schema.Task{
    Type: "PM2",
    Action:"hehe",
  })
  if finishTask.Status != false{
    t.Fail()
  }
  finishTask = Dispatch(&schema.Task{
    Type: "nginx",
    Action:"zzzz",
  })
  if finishTask.Status != false{
    t.Fail()
  }
}