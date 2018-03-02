package default_task

import (
  "log"
  "testing"
)

func TestGoExecute(t *testing.T){
  success, result:=execute("/Users/hyh/go-work/src/live-deploy-client/projects/", "git", "clone", "git@github.com:huyinghuan/aliyun-voice.git")
  if !success{
    log.Println(result)
    t.Fail()
  }else{
    log.Println(result)
  }
}