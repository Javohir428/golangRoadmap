package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

type Director struct {
	Departmant string
	Person
}

type TeamLead struct {
	TeamName string
	Person
}

type Programmer struct {
	TypeName string
	Person
}

type Task struct {
	ID     int
	Title  string
	Status string
}

func NewDirector(departmant string, p Person) *Director {
	d := new(Director)
	d.FirstName = p.FirstName
	d.LastName = p.LastName
	d.Phone = p.Phone
	d.Departmant = departmant
	return d
}

func NewTeamLead(teamName string, p Person) *TeamLead {
	t := new(TeamLead)
	t.FirstName = p.FirstName
	t.LastName = p.LastName
	t.Phone = p.Phone
	t.TeamName = teamName
	return t
}

func NewProgrammer(typeName string, p Person) *Programmer {
	programmer := new(Programmer)
	programmer.FirstName = p.FirstName
	programmer.LastName = p.LastName
	programmer.Phone = p.Phone
	programmer.TypeName = typeName
	return programmer
}

func NewTask(title string) *Task {
	t := new(Task)
	t.ID = ID
	t.Title = title
	t.Status = "Created"
	return t
}

func (director Director) GiveTask(title string) {
	newTask := NewTask(title)
	newTask.ID++
	TaskList = append(TaskList, *newTask)
}

func ChangeStatus(id int, str string) {
	for i := 0; i < len(TaskList); i++ {
		attr := &TaskList[i]
		if attr.ID == id {
			attr.Status = str
			break
		}
		if i >= len(TaskList)-1 {
			fmt.Println("Task not found")
		}
	}
}

func FindTask(id int) *Task {
	for i := 0; i < len(TaskList); i++ {
		attr := TaskList[i]
		if attr.ID == id {
			for k, v := range TaskList {
				if id == v.ID {
					return &TaskList[k]
				}
			}
			break
		}
		if i >= len(TaskList)-1 {
			fmt.Println("Task not found")
		}
	}
	return nil
}

func (t TeamLead) DelegateTask(id int) {
	task := FindTask(id)
	if task.Status == "Created" {
		ChangeStatus(id, "Dev")
	}
	if task.Status == "Test" {
		ChangeStatus(id, "Done")
	}
}

func (programmer Programmer) Develop(id int) {
	ChangeStatus(id, "Test")
}

var ID int = -1
var TaskList []Task

func main() {

	pDirector := new(Person)
	pDirector.FirstName = "Bob"
	pDirector.LastName = "Brown"
	pDirector.Phone = "654654"
	director := NewDirector("HR", *pDirector)
	director.GiveTask("HR Web")

	fmt.Println(director, FindTask(0))

	pTeamLead := new(Person)
	pTeamLead.FirstName = "Charlie"
	pTeamLead.LastName = "Wolf"
	pTeamLead.Phone = "123654"
	teamLead := NewTeamLead("Web", *pTeamLead)
	teamLead.DelegateTask(0)

	fmt.Println(teamLead, FindTask(0))

	pProgrammer := new(Person)
	pProgrammer.FirstName = "Nelly"
	pProgrammer.LastName = "Crouch"
	pProgrammer.Phone = "456789"
	programmer := NewProgrammer("Web programmer", *pProgrammer)
	programmer.Develop(0)

	fmt.Println(programmer, FindTask(0))

	teamLead.DelegateTask(0)

	fmt.Println(teamLead, FindTask(0))

}
