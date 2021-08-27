package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "javo"
	password  = "445"
	dbname    = "task_list"
	tablename = "task_list"
)

// Defining a struct type
type Task struct {
	ID        string
	Name      string
	Status    string
	Priority  string
	CreatedAt string
	CreatedBy string
	DueDate   string
}

type TaskManagerI interface {
	Add(task *Task) (string, error)
	Update(id string, task *Task) error
	Delete(id string) error
	Get(id string) (*Task, error)
	List() ([]*Task, error)
}

type TaskManager struct {
	db *sql.DB
}

func NewTaskManager(db *sql.DB) TaskManagerI {
	return TaskManager{
		db: db,
	}
}

func (c TaskManager) Add(task *Task) (string, error) {
	task.ID = uuid.NewString()
	insertQuery := fmt.Sprintf("INSERT INTO %s (id, name, status, priority, created_at, created_by, due_date) VALUES ('%s','%s', '%s', '%s', '%s', '%s', '%s');", tablename, task.ID, task.Name, task.Status, task.Priority, time.Now().Format(time.RFC3339), task.CreatedBy, task.DueDate)
	_, e := c.db.Exec(insertQuery)
	if e != nil {
		return "", e
	}
	return task.ID, nil
}

func (c TaskManager) Update(id string, task *Task) error {
	updateQuery := `UPDATE task_list SET status=$1, priority=$2, due_date=$3 WHERE id=$4`
	_, e := c.db.Exec(updateQuery, task.Status, task.Priority, task.DueDate, id)
	if e != nil {
		return e
	}
	return nil
}

func (c TaskManager) Delete(id string) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = '%s';", tablename, id)
	_, e := c.db.Exec(deleteQuery)
	if e != nil {
		return e
	}

	return nil
}

func (c TaskManager) Get(id string) (*Task, error) {
	task := new(Task)
	getQuery := "SELECT id, name, status, priority, created_at, created_by, due_date FROM " + tablename + " WHERE id=$1"

	row := c.db.QueryRow(getQuery, id)

	err := row.Scan(&task.ID, &task.Name, &task.Status, &task.Priority, &task.CreatedAt, &task.CreatedBy, &task.DueDate)

	if sql.ErrNoRows == err {
		fmt.Print("There is no retrieved rows")
	}

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (c TaskManager) List() ([]*Task, error) {
	var taskList []*Task
	task := new(Task)
	listQuery := fmt.Sprintf("SELECT * FROM %s;", tablename)
	rows, e := c.db.Query(listQuery)
	if e != nil {
		return nil, e
	}

	for rows.Next() {
		e = rows.Scan(&task.ID, &task.Name, &task.Status, &task.Priority, &task.CreatedAt, &task.CreatedBy, &task.DueDate)
		if e != nil {
			return nil, e
		}
		taskList = append(taskList, task)

	}

	return taskList, nil
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

	taskManager := NewTaskManager(db)
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
		Priority:  "High",
		CreatedBy: "HR",
		DueDate:   "10-12-2021",
	}

	id, err := taskManager.Add(task)
	if err != nil {
		fmt.Println("Error while adding task: ", err)
	}

	taskManager.Update(id, task2)

	row, err := taskManager.Get(id)
	if err != nil {
		fmt.Println("Error while getting task: ", err)
	}
	fmt.Println(row.Name)

	list, err := taskManager.List()
	if err != nil {
		fmt.Println("Error while getting task list: ", err)
	}
	for _, v := range list {
		fmt.Println(v.Name)
	}

	//taskManager.Delete(id)

}
