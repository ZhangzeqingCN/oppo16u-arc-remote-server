package main

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type DataResult[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type User struct {
	Id       uint   `gorm:"autoIncrement;primaryKey" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `gorm:"unique" json:"email"`
	Token    string `json:"token"`
}

type Task struct {
	Id          uint `gorm:"autoIncrement;primaryKey"`
	Name        string
	User        *User  `gorm:"references:Id;foreignKey:UserId;constraint:OnUpdate:Cascade,OnDelete:Set Null;"`
	UserId      uint   /* 任务的所有者 */
	PreTaskOfId uint   /* 作为谁的前置任务 */
	PreTasks    []Task `gorm:"references:Id;foreignKey:PreTaskOfId"` /* 前置任务 */
	SubTaskOfId uint   /* 作为谁的子任务 */
	SubTasks    []Task `gorm:"references:Id;foreignKey:SubTaskOfId"` /* 子任务 */
}

type Device struct {
	Id   uint `gorm:"autoIncrement;primaryKey"`
	Name string
}

type DeviceTaskRelation struct {
	Id       uint `gorm:"autoIncrement;primaryKey"`
	DeviceId uint
	TaskId   uint
	Device   *Device `gorm:"references:Id;foreignKey:DeviceId"`
	Task     *Task   `gorm:"references:Id;foreignKey:TaskId"`
}

type UserRelation struct {
	Id      uint `gorm:"autoIncrement;primaryKey"`
	User1Id uint
	User2Id uint
	User1   *User `gorm:"references:Id;foreignKey:User1Id"`
	User2   *User `gorm:"references:Id;foreignKey:User2Id"`
	Tag     string
}

type Notification struct {
	Id         uint `gorm:"autoIncrement;primaryKey"`
	ReceiverId uint
	Receiver   *User `gorm:"references:Id;foreignKey:ReceiverId"`
	Content    string
}
