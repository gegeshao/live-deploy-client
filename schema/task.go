package schema

type Task struct {
  ID int64  `xorm:"id unique autoincr index pk" json:"-" formam:"-"`
}