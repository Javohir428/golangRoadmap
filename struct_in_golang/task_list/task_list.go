package main

import (
	"fmt"
	"time"
)

// Defining a struct type
type Task struct {
	ID        int
	Name      string
	Status    string
	Priority  string
	CreatedAt string
	CreatedBy string
	DueDate   string
}

func Create(task *Task) {
	ID++
	task.ID = ID
	task.CreatedAt = CreatedTime
	TaskList = append(TaskList, *task)
}

func Update(task *Task) {
	for i := 0; i < len(TaskList); i++ {
		attr := &TaskList[i]
		if attr.ID == task.ID {

			attr.CreatedAt = task.CreatedAt

			if task.Name != "" {
				attr.Name = task.Name
			}
			if task.Status != "" {
				attr.Status = task.Status
			}
			if task.Priority != "" {
				attr.Priority = task.Priority
			}
			if task.CreatedBy != "" {
				attr.CreatedBy = task.CreatedBy
			}
			if task.DueDate != "" {
				attr.DueDate = task.DueDate
			}
			break
		}

		if i >= len(TaskList)-1 {
			fmt.Println("Record not found")
		}
	}
}

func Delete(id int) {
	for i := 0; i < len(TaskList); i++ {
		attr := TaskList[i]
		if attr.ID == id {
			for k, v := range TaskList {
				if id == v.ID {
					TaskList = append(TaskList[:k], TaskList[k+1:]...)
				}
			}
			break
		}
		if i >= len(TaskList)-1 {
			fmt.Println("Record not found")
		}
	}

}

func Get(id int) *Task {
	for i := 0; i < len(TaskList); i++ {
		attr := TaskList[i]
		if attr.ID == id {
			for k, v := range TaskList {
				if id == v.ID {
					return &TaskList[k]
				}
			}
			break
		}
		if i >= len(TaskList)-1 {
			fmt.Println("Record not found")
		}
	}
	return nil
}

func GetAll() []Task {
	return TaskList
}

var ID int = -1
var TaskList []Task
var CreatedTime string = time.Now().Format(time.RFC850)

func main() {

	task := &Task{Name: "asd", Status: "Loading", Priority: "high", CreatedBy: "me", DueDate: "2022"}
	task2 := &Task{ID: 1, Status: "Loaded", Priority: "low"}
	Create(task)
	Create(task)
	Create(task)
	Create(task)
	Create(task)
	Update(task2)

	Delete(1)

	Get(3)
	fmt.Println(GetAll())
}
