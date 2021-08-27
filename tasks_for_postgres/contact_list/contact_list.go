package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "javo"
	password  = "445"
	dbname    = "contact_list"
	tablename = "contact_list"
)

// Defining a struct type
type Contact struct {
	ID        string
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  int
}

type ContactManagerI interface {
	Add(contact *Contact) (string, error)
	Update(id string, contact *Contact) error
	Delete(id string) error
	Get(id string) (*Contact, error)
	List() ([]*Contact, error)
}

type ContactManager struct {
	db *sql.DB
}

func NewContactManager(db *sql.DB) ContactManagerI {
	return ContactManager{
		db: db,
	}
}

func (c ContactManager) Add(contact *Contact) (string, error) {
	contact.ID = uuid.NewString()
	insertQuery := fmt.Sprintf("INSERT INTO %s (id, first_name, last_name, phone, email, position) VALUES ('%s','%s', '%s', '%s', '%s', '%d');", tablename, contact.ID, contact.FirstName, contact.LastName, contact.Phone, contact.Email, contact.Position)
	_, e := c.db.Exec(insertQuery)
	if e != nil {
		return "", e
	}

	return contact.ID, nil
}

func (c ContactManager) Update(id string, contact *Contact) error {
	updateQuery := `UPDATE contact_list SET first_name=$1, last_name=$2, phone=$3, email=$4, position=$5 WHERE id=$6`
	_, e := c.db.Exec(updateQuery, contact.FirstName, contact.LastName, contact.Phone, contact.Email, contact.Position, id)
	if e != nil {
		return e
	}

	return nil
}

func (c ContactManager) Delete(id string) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = '%s';", tablename, id)
	_, e := c.db.Exec(deleteQuery)
	if e != nil {
		return e
	}

	return nil
}

func (c ContactManager) Get(id string) (*Contact, error) {
	contact := new(Contact)
	getQuery := "SELECT id, first_name, last_name, phone, email, position FROM " + tablename + " WHERE id=$1"

	row := c.db.QueryRow(getQuery, id)

	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)

	if sql.ErrNoRows == err {
		fmt.Print("There is no retrieved rows")
	}

	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (c ContactManager) List() ([]*Contact, error) {
	var contactList []*Contact
	contact := new(Contact)
	listQuery := fmt.Sprintf("SELECT * FROM %s;", tablename)
	rows, e := c.db.Query(listQuery)
	if e != nil {
		return nil, e
	}

	for rows.Next() {
		e = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
		if e != nil {
			return nil, e
		}
		contactList = append(contactList, contact)

	}

	return contactList, nil
}

func main() {
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

	fmt.Println("Connected!")

	contactManager := NewContactManager(db)
	contact := &Contact{
		FirstName: "Bob",
		LastName:  "Brown",
		Phone:     "123456",
		Email:     "test@mail.com",
		Position:  1,
	}

	contact2 := &Contact{
		FirstName: "Bob",
		LastName:  "Brown",
		Phone:     "6546545",
		Email:     "test@mail.com",
		Position:  5,
	}

	id, err := contactManager.Add(contact)
	if err != nil {
		fmt.Println("Error while adding contact: ", err)
	}

	contactManager.Update(id, contact2)

	row, err := contactManager.Get(id)
	if err != nil {
		fmt.Println("Error while getting contact: ", err)
	}
	fmt.Println(row.FirstName)

	list, err := contactManager.List()
	if err != nil {
		fmt.Println("Error while getting contact list: ", err)
	}
	for _, v := range list {
		fmt.Println(v.FirstName)
	}

	contactManager.Delete(id)

}
