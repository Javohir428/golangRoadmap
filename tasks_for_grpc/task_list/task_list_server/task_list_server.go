package main

import (
	"context"
	"database/sql"
	"fmt"
	"go_roadmap/tasks_for_grpc/task_list/api"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "javo"
	password  = "445"
	dbname    = "task_list"
	tablename = "task_list"
)

type TaskManagerServer struct {
	db *sql.DB
}

func (t *TaskManagerServer) Add(ctx context.Context, in *api.Task) (resp *api.AddResponse, err error) {

	in.Id = uuid.NewString()
	insertQuery := fmt.Sprintf("INSERT INTO %s (id, name, status, priority, created_at, created_by, due_date) VALUES ('%s','%s', '%s', '%s', '%s', '%s', '%s');", tablename, in.GetId(), in.GetName(), in.GetStatus(), in.GetPriority(), time.Now().Format(time.RFC3339), in.GetCreatedBy(), in.GetDueDate())
	_, err = t.db.Exec(insertQuery)
	if err != nil {
		return
	}
	resp = &api.AddResponse{Id: in.Id}

	return

}

func (t *TaskManagerServer) Delete(ctx context.Context, in *api.DeleteRequest) (resp *api.Empty, err error) {

	resp = &api.Empty{}
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = '%s';", tablename, in.GetId())
	_, err = t.db.Exec(deleteQuery)
	if err != nil {
		return
	}

	return
}

func (t *TaskManagerServer) Get(ctx context.Context, in *api.GetRequest) (resp *api.Task, err error) {

	task := new(api.Task)
	getQuery := "SELECT id, name, status, priority, created_at, created_by, due_date FROM " + tablename + " WHERE id=$1"

	row := t.db.QueryRow(getQuery, in.Id)

	err = row.Scan(&task.Id, &task.Name, &task.Status, &task.Priority, &task.CreatedAt, &task.CreatedBy, &task.DueDate)

	if sql.ErrNoRows == err {
		log.Printf("There is no retrieved rows")
	}
	if err != nil {
		return
	}
	resp = &api.Task{
		Id:        task.Id,
		Name:      task.Name,
		Status:    task.Status,
		Priority:  task.Priority,
		CreatedAt: task.CreatedAt,
		CreatedBy: task.CreatedBy,
		DueDate:   task.DueDate,
	}

	return
}

func (c *TaskManagerServer) Update(ctx context.Context, in *api.UpdateRequest) (resp *api.Empty, err error) {

	resp = &api.Empty{}
	updateQuery := `UPDATE task_list SET status=$1, priority=$2, due_date=$3 WHERE id=$4`
	_, err = c.db.Exec(updateQuery, in.Task.GetStatus(), in.Task.GetPriority(), in.Task.GetDueDate(), in.Id)
	if err != nil {
		return
	}

	return
}

func (c *TaskManagerServer) List(ctx context.Context, in *api.Empty) (resp *api.ListResponse, err error) {
	var taskList []*api.Task
	task := new(api.Task)
	listQuery := fmt.Sprintf("SELECT * FROM %s;", tablename)
	rows, err := c.db.Query(listQuery)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&task.Id, &task.Name, &task.Status, &task.Priority, &task.CreatedAt, &task.CreatedBy, &task.DueDate)
		if err != nil {
			return
		}
		taskList = append(taskList, task)

	}

	resp = &api.ListResponse{
		Task: taskList,
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := TaskManagerServer{db}

	grpcServer := grpc.NewServer()

	api.RegisterTaskServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
