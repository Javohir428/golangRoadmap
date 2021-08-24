package main

import (
	"fmt"
)

// Defining a struct type
type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  int
}

func Create(contact *Contact) {
	ID++
	contact.ID = ID
	ContactList = append(ContactList, *contact)
}

func Update(contact *Contact) {
	for i := 0; i < len(ContactList); i++ {
		attr := &ContactList[i]
		if attr.ID == contact.ID {
			if contact.FirstName != "" {
				attr.FirstName = contact.FirstName
			}
			if contact.LastName != "" {
				attr.LastName = contact.LastName
			}
			if contact.Phone != "" {
				attr.Phone = contact.Phone
			}
			if contact.Email != "" {
				attr.Email = contact.Email
			}
			if contact.Position != 0 {
				attr.Position = contact.Position
			}
			break
		}

		if i >= len(ContactList)-1 {
			fmt.Println("Record not found")
		}
	}
}

func Delete(id int) {
	for i := 0; i < len(ContactList); i++ {
		attr := ContactList[i]
		if attr.ID == id {
			for k, v := range ContactList {
				if id == v.ID {
					ContactList = append(ContactList[:k], ContactList[k+1:]...)
				}
			}
			break
		}
		if i >= len(ContactList)-1 {
			fmt.Println("Record not found")
		}
	}

}

func Get(id int) *Contact {
	for i := 0; i < len(ContactList); i++ {
		attr := ContactList[i]
		if attr.ID == id {
			for k, v := range ContactList {
				if id == v.ID {
					return &ContactList[k]
				}
			}
			break
		}
		if i >= len(ContactList)-1 {
			fmt.Println("Record not found")
		}
	}
	return nil
}

func GetAll() []Contact {
	return ContactList
}

var ID int = -1
var ContactList []Contact

func main() {
}
