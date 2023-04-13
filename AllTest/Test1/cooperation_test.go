package Test1

import (
	"sort"
	"testing"
)

type Task struct {
	taskType      string
	urgency       int
	importance    int
	estimatedTime int
	assignedUser  *User
}

type User struct {
	name          string
	habits        string
	abilities     string
	availability  int
	assignedTasks []*Task
	importance    int
}

type Device struct {
	deviceType    string
	location      string
	function      string
	availableTime int
	inUse         bool
}

type TaskFlow struct {
	step      int
	device    *Device
	user      *User
	startTime int
	endTime   int
}

type TaskFlowPlan struct {
	taskFlows []*TaskFlow
}

type ByStartTime []*TaskFlow

func (a ByStartTime) Len() int           { return len(a) }
func (a ByStartTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStartTime) Less(i, j int) bool { return a[i].startTime < a[j].startTime }

func assignTaskToUser(task *Task, users []*User) {
	var maxScore int
	var bestUser *User
	for _, user := range users {
		if user.availability < task.estimatedTime {
			continue
		}
		score := task.urgency*user.importance + user.availability - task.estimatedTime
		if score > maxScore {
			maxScore = score
			bestUser = user
		}
	}
	if bestUser != nil {
		bestUser.assignedTasks = append(bestUser.assignedTasks, task)
		task.assignedUser = bestUser
		bestUser.availability -= task.estimatedTime
	}
}

func generateTaskFlowPlan(tasks []*Task, users []*User, devices []*Device) *TaskFlowPlan {
	for _, task := range tasks {
		assignTaskToUser(task, users)
	}
	var taskFlows []*TaskFlow
	step := 1
	for _, device := range devices {
		sort.Slice(users, func(i, j int) bool {
			return users[i].availability < users[j].availability
		})
		for _, user := range users {
			if user.availability < device.availableTime {
				continue
			}
			for _, task := range user.assignedTasks {
				if task.assignedUser != user || task.estimatedTime > device.availableTime || device.inUse {
					continue
				}
				taskFlow := &TaskFlow{
					step:      step,
					device:    device,
					user:      user,
					startTime: device.availableTime - task.estimatedTime,
					endTime:   device.availableTime,
				}
				taskFlows = append(taskFlows, taskFlow)
				step++
				device.availableTime -= task.estimatedTime
				device.inUse = true
				break
			}
			if device.inUse {
				break
			}
		}
		device.inUse = false
	}
	sort.Sort(ByStartTime(taskFlows))
	return &TaskFlowPlan{taskFlows: taskFlows}
}

func TestPr1(t *testing.T) {
	user1 := &User{
		name:         "Alice",
		habits:       "Cooking on weekdays",
		abilities:    "Baking",
		availability: 8,
	}
	user2 := &User{
		name:         "Charlie",
		habits:       "Cleaning on Saturdays",
		abilities:    "Gardening",
		availability: 6,
	}
	user3 := &User{
		name:         "Bob",
		habits:       "Grocery shopping on Sundays",
		abilities:    "Painting",
		availability: 4,
	}
	device1 := &Device{
		deviceType:    "Oven",
		location:      "Kitchen",
		function:      "Baking",
		availableTime: 10,
		inUse:         false,
	}
	device2 := &Device{
		deviceType:    "Washing machine",
		location:      "Laundry room",
		function:      "Washing clothes",
		availableTime: 8,
		inUse:         false,
	}
	task1 := &Task{
		taskType:      "Bake a cake",
		urgency:       3,
		importance:    4,
		estimatedTime: 2,
		assignedUser:  nil,
	}
	task2 := &Task{
		taskType:      "Wash clothes",
		urgency:       2,
		importance:    3,
		estimatedTime: 1,
		assignedUser:  nil,
	}
	task3 := &Task{
		taskType:      "Paint a picture",
		urgency:       4,
		importance:    3,
		estimatedTime: 3,
		assignedUser:  nil,
	}
	tasks := []*Task{task1, task2, task3}
	users := []*User{user1, user2, user3}
	devices := []*Device{device1, device2}
	taskFlowPlan := generateTaskFlowPlan(tasks, users, devices)
	for _, taskFlow := range taskFlowPlan.taskFlows {
		t.Logf("Step %d: User %s uses %s from %d to %d to %s\n", taskFlow.step, taskFlow.user.name, taskFlow.device.deviceType, taskFlow.startTime, taskFlow.endTime, taskFlow.device.function)
	}
}
