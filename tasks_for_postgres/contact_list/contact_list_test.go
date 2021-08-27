package main

import (
	"os"
	"testing"

	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var cm ContactManagerI

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

	cm = NewContactManager(db)

	os.Exit(m.Run())
}

func TestAdd(t *testing.T) {

	contact := &Contact{
		FirstName: "Bob",
		LastName:  "Brown",
		Phone:     "123456",
		Email:     "test@mail.com",
		Position:  1,
	}
	id, err := cm.Add(contact)
	if err != nil {
		fmt.Println("Error while adding contact: ", err)
	}

	result, _ := cm.Get(id)
	expect := *contact

	if *result != expect {
		t.Errorf("Record not added")
	}
}

func TestUpdate(t *testing.T) {

	contact := &Contact{
		FirstName: "Bob",
		LastName:  "Brown",
		Phone:     "654321",
		Email:     "test@mail.com",
		Position:  5,
	}

	id, errAdd := cm.Add(contact)
	if errAdd != nil {
		fmt.Println("Error while adding contact: ", errAdd)
	}
	errUpdate := cm.Update(id, contact)
	if errUpdate != nil {
		fmt.Println("Error while adding contact: ", errUpdate)
	}

	result, _ := cm.Get(id)
	expect := *contact

	if *result != expect {
		t.Errorf("Record not updated")
	}
}

func TestDelete(t *testing.T) {
	contact := &Contact{
		FirstName: "James",
		LastName:  "Brown",
		Phone:     "654321",
		Email:     "test@mail.com",
		Position:  4,
	}

	id, errAdd := cm.Add(contact)
	if errAdd != nil {
		fmt.Println("Error while adding contact: ", errAdd)
	}

	cm.Delete(id)

	_, result := cm.Get(id)
	expect := sql.ErrNoRows

	fmt.Println(" Result ", result.Error())
	fmt.Println(" Expect ", expect.Error())

	if result.Error() != expect.Error() {
		t.Errorf("Record not deleted")
	}

}

func TestGet(t *testing.T) {

	newContact := &Contact{
		FirstName: "Bob",
		LastName:  "Brown",
		Phone:     "654321",
		Email:     "test@mail.com",
		Position:  5,
	}

	id, errAdd := cm.Add(newContact)
	if errAdd != nil {
		fmt.Println("Error while adding contact: ", errAdd)
	}

	contact, errGet := cm.Get(id)
	if errGet != nil {
		fmt.Println("Error while adding contact: ", errGet)
	}

	result := &contact.ID
	expect := id

	if *result != expect {
		t.Errorf("Record not found")
	}
}

func TestList(t *testing.T) {

	contact := &Contact{
		FirstName: "Charlie",
		LastName:  "Brown",
		Phone:     "789798",
		Email:     "test2@mail.com",
		Position:  3,
	}

	_, errAdd := cm.Add(contact)
	if errAdd != nil {
		fmt.Println("Error while adding contact: ", errAdd)
	}

	contactList, errUpdate := cm.List()
	if errUpdate != nil {
		fmt.Println("Error while getting contact list: ", errUpdate)
	}

	contactListLen := len(contactList)

	if contactListLen == 0 {
		t.Errorf("Records not found")
	}
}
