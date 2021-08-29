package main

import (
	"context"
	"database/sql"
	"fmt"
	"go_roadmap/tasks_for_grpc/contact_list/api"
	"log"
	"net"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "javo"
	password  = "445"
	dbname    = "contact_list"
	tablename = "contact_list"
)

type ContactManagerServer struct {
	db *sql.DB
}

func (c *ContactManagerServer) Add(ctx context.Context, in *api.Contact) (resp *api.AddResponse, err error) {

	in.Id = uuid.NewString()
	insertQuery := fmt.Sprintf("INSERT INTO %s (id, first_name, last_name, phone, email, position) VALUES ('%s','%s', '%s', '%s', '%s', '%d');", tablename, in.Id, in.GetFirstName(), in.GetLastName(), in.GetPhone(), in.GetEmail(), in.GetPosition())
	_, err = c.db.Exec(insertQuery)
	if err != nil {
		return
	}
	resp = &api.AddResponse{Id: in.Id}
	//log.Printf(in.Id)

	return

}

func (c *ContactManagerServer) Delete(ctx context.Context, in *api.DeleteRequest) (resp *api.Empty, err error) {

	resp = &api.Empty{}
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = '%s';", tablename, in.GetId())
	_, err = c.db.Exec(deleteQuery)
	if err != nil {
		return
	}

	return
}

func (c *ContactManagerServer) Get(ctx context.Context, in *api.GetRequest) (resp *api.Contact, err error) {

	contact := new(api.Contact)
	getQuery := "SELECT id, first_name, last_name, phone, email, position FROM " + tablename + " WHERE id=$1"

	row := c.db.QueryRow(getQuery, in.Id)

	err = row.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)

	if sql.ErrNoRows == err {
		log.Printf("There is no retrieved rows")
	}
	if err != nil {
		return
	}
	resp = &api.Contact{
		Id:        contact.Id,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
		Position:  contact.Position,
	}

	return
}

func (c *ContactManagerServer) Update(ctx context.Context, in *api.UpdateRequest) (resp *api.Empty, err error) {

	resp = &api.Empty{}
	updateQuery := `UPDATE contact_list SET first_name=$1, last_name=$2, phone=$3, email=$4, position=$5 WHERE id=$6`
	_, err = c.db.Exec(updateQuery, in.Contact.GetFirstName(), in.Contact.GetLastName(), in.Contact.GetPhone(), in.Contact.GetEmail(), in.Contact.GetPosition(), in.Id)
	if err != nil {
		return
	}

	return
}

func (c *ContactManagerServer) List(ctx context.Context, in *api.Empty) (resp *api.ListResponse, err error) {
	var contactList []*api.Contact
	contact := new(api.Contact)
	listQuery := fmt.Sprintf("SELECT * FROM %s;", tablename)
	rows, err := c.db.Query(listQuery)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
		if err != nil {
			return
		}
		contactList = append(contactList, contact)

	}

	resp = &api.ListResponse{
		Contact: contactList,
	}
	return
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9010))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := ContactManagerServer{db}

	grpcServer := grpc.NewServer()

	api.RegisterContactServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
