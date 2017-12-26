package main

import (
  "live-deploy-client/schema"
  "live-deploy-client/utils"
	"log"
	"live-deploy-client/task"
  "time"
)



func main(){
	config, err := utils.InitConfig("./conf/config.yaml")
	if err != nil {
		log.Fatalln(err)
		return
	}
  if err:= schema.InitDriver(); err!=nil{
    log.Fatalln(err)
	  return
  }
	//校验密钥
  if len(config.PrivateKey) != 24{
    log.Fatalln("密钥错误!")
  }
  if err:=task.Check(); err!=nil{
  	log.Println(err)
    log.Fatalln("密钥错误!")
  }

	//task.Get()
	taskTimer := make(chan int)
	go func(){
		for {
			taskTimer <- 1
			time.Sleep(5 * time.Second)
		}
 	}()

  L, err:=task.CreateLuaVM()
  if err !=nil{
    log.Fatal(err)
  }
 	func(){
 		for{
 			<-taskTimer
 			task.Get(L)
		}
	}()

}