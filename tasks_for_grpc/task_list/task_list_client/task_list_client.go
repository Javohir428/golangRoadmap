package main

import (
	"go_roadmap/tasks_for_grpc/task_list/api"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewTaskServiceClient(conn)

	responseAdd, err := c.Add(context.Background(), &api.Task{
		Name:      "Web service for HR",
		Status:    "Created",
		Priority:  "High",
		CreatedBy: "HR",
		DueDate:   "10-12-2021",
	})
	if err != nil {
		log.Fatalf("Error while adding task: %s", err)
	}
	responseGet, err := c.Get(context.Background(), &api.GetRequest{Id: responseAdd.GetId()})
	if err != nil {
		log.Fatalf("Error while getting task: %s", err)
	}
	log.Println("Add response", responseGet.GetId(), responseGet.GetName(), responseGet.GetStatus())
	task := new(api.Task)
	task.Status = "Done"
	task.Priority = "Low"
	task.DueDate = "10-12-2021"
	_, err = c.Update(context.Background(), &api.UpdateRequest{Id: responseAdd.Id, Task: task})
	if err != nil {
		log.Fatalf("Error while updating task: %s", err)
	}

	responseGet2, err := c.Get(context.Background(), &api.GetRequest{Id: responseAdd.GetId()})
	if err != nil {
		log.Fatalf("Error while getting task: %s", err)
	}

	c.Delete(context.Background(), &api.DeleteRequest{Id: responseGet2.GetId()})

	responseGet2, err = c.Get(context.Background(), &api.GetRequest{Id: responseAdd.GetId()})
	if err != nil {
		log.Fatalf("Error while getting task: %s", err)
	}
	log.Println("Update response", responseGet2.GetId(), responseGet2.GetName(), responseGet2.GetStatus())
	log.Println("Get response", responseGet2.GetId(), responseGet2.GetName(), responseGet2.GetStatus())

	responseList, err := c.List(context.Background(), &api.Empty{})
	if err != nil {
		log.Println("Delete response", err)
	}

	log.Println("List response", responseList.Task)

}
