package main

import (
	"live-deploy-client/utils"
	"log"
	"fmt"
	"time"
)

func main(){
	_, err := utils.InitConfig("./conf/config.yaml")
	if err != nil {
		log.Fatalln(err)
		return
	}
	taskTimer := make(chan int)
	go func(){
		for {
			taskTimer <- 1
			time.Sleep(1 * time.Second)
		}
 	}()

 	func(){
 		for{
 			<-taskTimer
 			fmt.Println(time.Now().Format("15:04:05"))
 			time.Sleep(2 * time.Second)
		}
	}()

}