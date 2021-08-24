package main

import (
	"testing"
)

func TestCreate(t *testing.T) {
	task := &Task{Name: "a", Status: "b", Priority: "a@b.c", CreatedBy: "Me", DueDate: "2022"}
	Create(task)

	result := TaskList[0]
	expect := *task

	if result != expect {
		t.Errorf("Record not created")
	}
}

func TestUpdate(t *testing.T) {
	task := &Task{ID: 0, Name: "b", Status: "c", Priority: "a@b.c", CreatedBy: "Bob", DueDate: "2025"}
	Update(task)

	result := TaskList[0]
	expect := *task

	if result != expect {
		t.Errorf("Record not updated")
	}
}

func TestDelete(t *testing.T) {
	Delete(0)
	if len(TaskList) != 0 {
		t.Errorf("Record not deleted")
	}
}

func TestGet(t *testing.T) {
	task := &Task{Name: "a", Status: "b", Priority: "a@b.c", CreatedBy: "Me", DueDate: "2022"}
	Create(task)

	result := *task
	expect := *Get(1)

	if result != expect {
		t.Errorf("Record not found")
	}
}

func TestGetAll(t *testing.T) {
	task := &Task{Name: "a", Status: "b", Priority: "a@b.c", CreatedBy: "Me", DueDate: "2022"}
	Create(task)

	if len(GetAll()) != len(TaskList) {
		t.Errorf("Records not found")
	}
}
