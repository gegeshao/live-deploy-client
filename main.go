package main

import (
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "live-deploy-client/vm"
  "log"
	"live-deploy-client/task"
  "time"
)

func init(){
  config, err := utils.InitConfig("./conf/config.yaml")
  if err != nil {
    log.Fatalln(err)
    return
  }

  //校验密钥
  if len(config.System.PrivateKey) != 24{
    log.Fatalln("密钥错误!")
  }
  if err:=task.Check(); err!=nil{
    log.Println(err)
    log.Fatalln("密钥错误!")
  }
  //初始化虚拟机
  if err := vm.InitLuaVM(); err!=nil{
    log.Fatal(err)
  }
  //初始化sqlite3
  if err:= schema.InitDriver(); err!=nil{
    log.Fatalln(err)
    return
  }
}

func main(){
	//task.Get()
	//阅读 《Go语言高级编程》 ，改为匿名struct可以减少内存空间使用
	taskTimer := make(chan struct{})
	go func(){
		for {
			taskTimer <- struct{}{}
			time.Sleep(5 * time.Second)
		}
 	}()

 	func(){
 		for{
 			<-taskTimer
 			task.Get()
		}
	}()
}