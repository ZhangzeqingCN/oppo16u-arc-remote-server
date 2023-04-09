package main

import "time"

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type DataResult[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type MyModel struct {
	Id   uint   `gorm:"autoIncrement;primaryKey" json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id       uint   `gorm:"autoIncrement;primaryKey" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `gorm:"unique" json:"email"`
	Token    string `json:"token"`
}

type Task struct {
	Id              uint `gorm:"autoIncrement;primaryKey" json:"id"`
	Name            string
	Description     string
	User            *User  `gorm:"references:Id;foreignKey:UserId;constraint:OnUpdate:Cascade,OnDelete:Set Null;"`
	UserId          uint   /* 任务的所有者 */
	PreTaskOfId     uint   /* 作为谁的前置任务 */
	PreTasks        []Task `gorm:"references:Id;foreignKey:PreTaskOfId" json:"preTasks"` /* 前置任务 */
	SubTaskOfId     uint   /* 作为谁的子任务 */
	SubTasks        []Task `gorm:"references:Id;foreignKey:SubTaskOfId" json:"subTasks"` /* 子任务 */
	Manual          bool   /* 需要人操作 */
	DeviceExclusive bool   /* 独占设备 */
}

type Device struct {
	Id   uint   `gorm:"autoIncrement;primaryKey" json:"id"`
	Name string `json:"name"`
}

type DeviceTaskRelation struct {
	Id       uint `gorm:"autoIncrement;primaryKey" json:"id"`
	DeviceId uint
	TaskId   uint
	Device   *Device `gorm:"references:Id;foreignKey:DeviceId" json:"device"`
	Task     *Task   `gorm:"references:Id;foreignKey:TaskId" json:"task"`
}

type UserRelation struct {
	Id      uint `gorm:"autoIncrement;primaryKey" json:"id"`
	User1Id uint
	User2Id uint
	User1   *User `gorm:"references:Id;foreignKey:User1Id" json:"user1"`
	User2   *User `gorm:"references:Id;foreignKey:User2Id" json:"user2"`
	Tag     string
}

type Notification struct {
	Id                     uint `gorm:"autoIncrement;primaryKey" json:"id"`
	ReceiverId             uint
	Receiver               *User `gorm:"references:Id;foreignKey:ReceiverId" json:"receiver"`
	FromUserId             uint
	FromDeviceId           uint
	FromUser               *User   `gorm:"references:Id;foreignKey:FromUserId" json:"fromUser"`
	FromDevice             *Device `gorm:"references:Id;foreignKey:FromDeviceId" json:"fromDevice"`
	Context                string
	Read                   bool
	NotificationArriveTime time.Time
}
