package main

import (
	"go_roadmap/tasks_for_grpc/contact_list/api"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9010", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewContactServiceClient(conn)

	responseAdd, err := c.Add(context.Background(), &api.Contact{FirstName: "Asus", LastName: "Acer", Phone: "123", Email: "a@a.a", Position: 5})
	if err != nil {
		log.Fatalf("Error while adding contact: %s", err)
	}
	responseGet, err := c.Get(context.Background(), &api.GetRequest{Id: responseAdd.GetId()})
	if err != nil {
		log.Fatalf("Error while getting contact: %s", err)
	}
	log.Println("Add response", responseGet.GetId(), responseGet.GetFirstName(), responseGet.GetLastName())
	contact := new(api.Contact)
	contact.FirstName = "Apple"
	contact.LastName = "Macbook"
	_, err = c.Update(context.Background(), &api.UpdateRequest{Id: responseAdd.Id, Contact: contact})
	if err != nil {
		log.Fatalf("Error while updating contact: %s", err)
	}

	responseGet2, err := c.Get(context.Background(), &api.GetRequest{Id: responseAdd.GetId()})
	if err != nil {
		log.Fatalf("Error while getting contact: %s", err)
	}

	c.Delete(context.Background(), &api.DeleteRequest{Id: responseGet2.GetId()})

	responseGet2, err = c.Get(context.Background(), &api.GetRequest{Id: responseAdd.GetId()})
	if err != nil {
		log.Fatalf("Error while getting contact: %s", err)
	}
	log.Println("Update response", responseGet2.GetId(), responseGet2.GetFirstName(), responseGet2.GetLastName(), responseGet2.GetPhone())
	log.Println("Get response", responseGet2.GetId(), responseGet2.GetFirstName(), responseGet2.GetLastName(), responseGet2.GetPhone())

	responseList, err := c.List(context.Background(), &api.Empty{})
	if err != nil {
		log.Println("Delete response", err)
	}

	log.Println("List response", responseList.Contact)

}
