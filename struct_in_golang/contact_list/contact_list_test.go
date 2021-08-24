package main

import (
	"testing"
)

func TestCreate(t *testing.T) {
	contact := &Contact{FirstName: "a", LastName: "d", Email: "a@d.c", Position: 1}
	Create(contact)

	result := ContactList[0]
	expect := *contact

	if result != expect {
		t.Errorf("Record not created")
	}
}

func TestUpdate(t *testing.T) {
	contact := &Contact{ID: 0, FirstName: "b", LastName: "e", Email: "a@d.c", Position: 2}
	Update(contact)

	result := ContactList[0]
	expect := *contact

	if result != expect {
		t.Errorf("Record not updated")
	}
}

func TestDelete(t *testing.T) {
	Delete(0)
	if len(ContactList) != 0 {
		t.Errorf("Record not deleted")
	}
}

func TestGet(t *testing.T) {
	contact := &Contact{FirstName: "a", LastName: "d", Email: "a@d.c", Position: 1}
	Create(contact)

	result := *contact
	expect := *Get(1)

	if result != expect {
		t.Errorf("Record not found")
	}
}

func TestGetAll(t *testing.T) {
	contact := &Contact{FirstName: "z", LastName: "l", Email: "a@d.c", Position: 2}
	Create(contact)

	if len(GetAll()) != len(ContactList) {
		t.Errorf("Records not found")
	}
}
