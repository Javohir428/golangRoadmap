package main

import (
	"os"
	"testing"

	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var cm TaskManagerI

func TestMain(m *testing.M) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	cm = NewTaskManager(db)

	os.Exit(m.Run())
}

func TestAdd(t *testing.T) {

	task := &Task{
		Name:      "Web service for HR",
		Status:    "Created",
		Priority:  "High",
		CreatedBy: "HR",
		DueDate:   "10-12-2021",
	}
	id, errAdd := cm.Add(task)
	if errAdd != nil {
		fmt.Println("Error while adding task: ", errAdd)
	}

	result, errGet := cm.Get(id)
	if errGet != nil {
		fmt.Println("Error while getting task: ", errGet)
	}
	expect := *task

	if result.ID != expect.ID {
		t.Errorf("Record not added")
	}
}

func TestUpdate(t *testing.T) {

	task := &Task{
		Name:      "Web service for HR",
		Status:    "Created",
		Priority:  "High",
		CreatedBy: "HR",
		DueDate:   "10-12-2021",
	}

	task2 := &Task{
		Name:      "Web service for HR",
		Status:    "Done",
		Priority:  "Low",
		CreatedBy: "HR",
		DueDate:   "10-12-2021",
	}

	id, errAdd := cm.Add(task)
	if errAdd != nil {
		fmt.Println("Error while adding task: ", errAdd)
	}
	errUpdate := cm.Update(id, task2)
	if errUpdate != nil {
		fmt.Println("Error while adding task: ", errUpdate)
	}

	result, _ := cm.Get(id)
	expect := *task2

	isEqual := (result.Name == expect.Name &&
		result.Status == expect.Status &&
		result.Priority == expect.Priority)

	if !isEqual {
		t.Errorf("Record not updated")
	}
}

func TestDelete(t *testing.T) {
	task := &Task{
		Name:      "Web service for HR",
		Status:    "Done",
		Priority:  "High",
		CreatedBy: "HR",
		DueDate:   "10-12-2021",
	}

	id, errAdd := cm.Add(task)
	if errAdd != nil {
		fmt.Println("Error while adding task: ", errAdd)
	}

	cm.Delete(id)

	_, result := cm.Get(id)
	expect := sql.ErrNoRows

	if result.Error() != expect.Error() {
		t.Errorf("Record not deleted")
	}

}

func TestGet(t *testing.T) {

	task := &Task{
		Name:      "Web service for HR",
		Status:    "Done",
		Priority:  "High",
		CreatedBy: "HR",
		DueDate:   "10-12-2021",
	}

	id, errAdd := cm.Add(task)
	if errAdd != nil {
		fmt.Println("Error while adding task: ", errAdd)
	}

	task, errGet := cm.Get(id)
	if errGet != nil {
		fmt.Println("Error while adding task: ", errGet)
	}

	result := &task.ID
	expect := id

	if *result != expect {
		t.Errorf("Record not found")
	}
}

func TestList(t *testing.T) {

	task := &Task{
		Name:      "Web service for HR",
		Status:    "Done",
		Priority:  "High",
		CreatedBy: "HR",
		DueDate:   "10-12-2021",
	}

	_, errAdd := cm.Add(task)
	if errAdd != nil {
		fmt.Println("Error while adding task: ", errAdd)
	}

	taskList, errUpdate := cm.List()
	if errUpdate != nil {
		fmt.Println("Error while getting task list: ", errUpdate)
	}

	taskListLen := len(taskList)

	if taskListLen == 0 {
		t.Errorf("Records not found")
	}
}
