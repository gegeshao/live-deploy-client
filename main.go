package main

import (
	"live-deploy-client/utils"
	"log"
	"live-deploy-client/task"
  "fmt"
)

func main(){
	config, err := utils.InitConfig("./conf/config.yaml")
	if err != nil {
	  fmt.Println(err,999)
		log.Fatalln(err)
		return
	}
	//校验密钥
  if len(config.PrivateKey) != 24{
    log.Fatalln("密钥错误!")
  }

  if err:=task.Check(); err!=nil{
    log.Fatalln("密钥错误!")
  }


	//taskTimer := make(chan int)
	//go func(){
	//	for {
	//		taskTimer <- 1
	//		time.Sleep(1 * time.Second)
	//	}
 	//}()
  //
 	//func(){
 	//	for{
 	//		<-taskTimer
 	//		fmt.Println(time.Now().Format("15:04:05"))
 	//		time.Sleep(2 * time.Second)
	//	}
	//}()

}