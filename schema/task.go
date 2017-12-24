package schema

import  (
  _ "github.com/mattn/go-sqlite3"
  "github.com/go-xorm/xorm"
  "fmt"
  "log"
)

type Task struct {
  ID int64  `xorm:"id unique autoincr index pk" json:"-" formam:"-"`
  TrackID      int64  `xorm:"track_id" json:"track_id"`   //可追踪
  TrackKey     int64  `xorm:"track_key" json:"track_key"` //可追踪
  TaskID int64 `xorm:"task_id" json:"task_id"` //任务id
  Content string `xorm:"content" json:"content"` //完成任务结果
  Type string `xorm:"type"` //任务类型
  Action string `xorm:"action"` //任务动作
  Status int `xorm:"task_id"` //-2 失败 未发送  -1 已完成 未发送 1 已完成 已发送  2 失败，已发送
}

type TaskClientFinish struct{
  ID  int64  `json:"id"`  //任务ID
  Status bool `json:"status"` // 任务状态
  Content string `json:"content"` //任务完成结果
}
           

var engine *xorm.Engine

// InitDriver 初始化数据库链接
func InitDriver() error {
  var connectErr error
  engine, connectErr = xorm.NewEngine("sqlite3", "task.db")

  engine.ShowSQL(true)

  if connectErr != nil {
    return connectErr
  }
  if engine.Ping() != nil {
    return fmt.Errorf("数据库链接失败")
  }
  engine.Sync(new(Task))
  log.Printf("数据库连接成功")
  return nil
}

func AddTask(task *Task)error{
  _, err:=engine.InsertOne(task)
  return err
}