package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

// TID is used only for delete request, where data needs to be like this: "data":{"id":1}
type TID struct {
	ID float64 `json:"id"`
}

// These are used for creation of objects

// Teacher contains all information about teacher
type Teacher struct {
	Salary    float64  `json:"salary"`
	Subject   string   `json:"subject"`
	Classroom []string `json:"classroom"`
	Person    struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

// Student contains all information about student
type Student struct {
	Year   float64 `json:"year"`
	Index  string  `json:"index"`
	Person struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

// Staff contains all information about staff
type Staff struct {
	Salary    float64  `json:"salary"`
	Classroom []string `json:"classroom"`
	Person    struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

// These are similar to ones above, but also conatin ID field
// They are used for updating

// TeacherU contains all information about teacher needed to update
type TeacherU struct {
	ID        float64  `json:"id"`
	Salary    float64  `json:"salary"`
	Subject   string   `json:"subject"`
	Classroom []string `json:"classroom"`
	Person    struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

// StudentU contains all information about student needed to update
type StudentU struct {
	ID     float64 `json:"id"`
	Year   float64 `json:"year"`
	Index  string  `json:"index"`
	Person struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

// StaffU contains all information about staff needed to update
type StaffU struct {
	ID        float64  `json:"id"`
	Salary    float64  `json:"salary"`
	Classroom []string `json:"classroom"`
	Person    struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

// Action represents json request
// It can server all 4 requests
type Action struct {
	Action string      `json:"action"`
	Object string      `json:"object"`
	Data   interface{} `json:"data"`
}

func main() {
	// Establish connection
	conn, err := net.Dial("tcp", "127.0.0.1:15395")
	if err != nil {
		panic(err)
	}
	// defer conn.Close() // no need for this

top:
	for {
		// get command
		var inp string
		var msg []byte
		fmt.Print("Please select action (select/create/exit): ")
		fmt.Scan(&inp)

		// parse input
		switch inp {
		case "select":
			// ask for ID
			var ID float64
			fmt.Print("ID: ")
			fmt.Scan(&ID)

			// this function handles everything
			HandleSelected(ID, conn)

			// leave an empty line, for easy understanding
			fmt.Println()
			continue
		case "create":
			// GetJsonCreate handels everything and returns msg in []byte, ready to be sent
			msg = GetJsonCreate()
		case "exit":
			// transmit to server that client is disconecting
			conn.Write([]byte("stop"))
			fmt.Println("Exiting...")

			// close conn and quit
			conn.Close()
			os.Exit(0)
		default: // if command unknown, try again
			continue top
		}
		// Send msg that we got from out commands
		conn.Write(msg)

		// Recieve responce
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			// skip displaying it if we got error
			fmt.Println(err)
			continue
		}

		// Print responce
		fmt.Printf("Response: %s\n", string(buf[:n]))
	}

}

func HandleSelected(ID float64, conn net.Conn) {
	// this is where the fun begins

	var job string // never changes

	// make request of getting job of selected id
	// request will lokk something like this:
	// {"action":3,"object":"Unknown"}
	conn.Write([]byte(fmt.Sprintf("{\"action\":\"%.0f\",\"object\":\"Unknown\"}", ID)))

	// recieve our job string
	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if err != nil {
		// exit HandleSelection if we didn't get job label
		fmt.Println(err)
		return
	}

	// if we got empty string, alert user and exit
	if string(buf[:n]) == "ID not found" {
		fmt.Println("Invalid ID")
		return
	}

	// parse job string
	job = string(buf[:n])

	var end bool // used for exiting

	for {
		var inp string // is user input
		var msg []byte // encoded json msg, that will be sent to server
		var a Action   // not encoded json msg

		// Get command
		fmt.Printf("Selected id:%.0f: Please select action (delete/update/read/exit): ", ID)
		fmt.Scan(&inp)

		// parse command
		switch inp {
		case "delete":
			// ex. {"action":"delete","object":"Student","data":{"id":2}}
			a.Action = "delete"
			a.Data = TID{ID}
			a.Object = job

			// exit this function after sending msg
			end = true
		case "update":
			// ex. {"action":"update","object":"Teacher","data":{"id":3,"subjeect":"oh yeah maaaaath","salary":1234,"classroom":["1", "2", "and 3"],"person":{"surname":"Kisliy","personalCode":"654635464767547"}}}
			a.Action = "update"

			// select certain job and fill a.data acordingly
			switch job {
			case "Teacher":
				a = CreateTeacher()
				var tu TeacherU
				tu.Person = a.Data.(Teacher).Person
				tu.Classroom = a.Data.(Teacher).Classroom
				tu.Salary = a.Data.(Teacher).Salary
				tu.Subject = a.Data.(Teacher).Subject
				tu.ID = ID
				a.Data = tu
			case "Student":
				a = CreateStudent()
				var tu StudentU
				tu.Person = a.Data.(Student).Person
				tu.Index = a.Data.(Student).Index
				tu.Year = a.Data.(Student).Year
				tu.ID = ID
				a.Data = tu
			case "Staff":
				a = CreateStaff()
				var tu StaffU
				tu.Person = a.Data.(Staff).Person
				tu.Classroom = a.Data.(Staff).Classroom
				tu.Salary = a.Data.(Staff).Salary
				tu.ID = ID
				a.Data = tu
			}
			a.Action = "update"
		case "read":
			// ex. {"action":"read","object":"Teacher","data":{"id":3}}
			a.Action = "read"
			a.Object = job
			a.Data = TID{ID}
		case "exit":
			// alert user and exit this function
			// we will return to main for loop
			fmt.Println("Deselected")
			return
		default:
			// if command is invalid, try again
			continue
		}

		// encode msg
		msg, _ = json.Marshal(a)

		// Send msg
		conn.Write(msg)

		// Recieve responce
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			// if something went wrong, dont output msg and continue
			fmt.Println(err)
			continue
		}

		// output responce for the user
		fmt.Printf("Response: %s\n", string(buf[:n]))

		// exit if command was "delete"
		if end {
			break
		}
	}
	fmt.Println("Deselected")
}

// SelectJob prompts user with choice of all job, until user decides
func SelectJob() string {
	var inp string // here input will be stored
	for {
		fmt.Print("Please select job (Teacher/Student/Staff): ") // prompt

		// getting input
		fmt.Scan(&inp)

		// making sure, that capitalization wont matter
		inp = strings.ToLower(inp)

		// parse input
		switch inp {
		case "teacher":
			return "Teacher"
		case "student":
			return "Student"
		case "staff":
			return "Staff"
		}

		// if user entered something bad, prompt hom again
	}
}

// GetJsonCreate gets json request of creation. User selects who he will create
func GetJsonCreate() []byte {
	var msg []byte // enoced json msg, that will be sent

	// getting job label from user
	job := SelectJob()

	// parsing job label, filling msg acordingly
	switch job {
	case "Teacher":
		msg, _ = json.Marshal(CreateTeacher())
	case "Student":
		msg, _ = json.Marshal(CreateStudent())
	case "Staff":
		msg, _ = json.Marshal(CreateStaff())
	}

	// returning completed msg
	return msg
}

// CreateTeacher simply asks a bunch of questoins and fills Teacher struct, returns as part of Action struct
func CreateTeacher() Action {
	fmt.Println("There are no undos for this action!")

	// filling basic fields
	var t Teacher
	fmt.Print("Please enter teacher's name, surname and personal code: ")
	fmt.Scan(&t.Person.Name)
	fmt.Scan(&t.Person.Surname)
	fmt.Scan(&t.Person.PersonalCode)
	fmt.Print("Subject: ")
	fmt.Scan(&t.Subject)
	fmt.Print("Salary: ")
	fmt.Scan(&t.Salary)

	// filling classes slice
	var a []string
	var inp string
	for {
		fmt.Print("Please enter one or more classrooms, \"exit\" to stop: ")
		fmt.Scan(&inp)
		if inp == "exit" {
			break
		}
		a = append(a, inp)
	}

	//assigning everything

	t.Classroom = a
	fmt.Println("Done!")
	var A Action
	A.Data = t
	A.Action = "create"
	A.Object = "Teacher"
	return A
}

// CreateStaff simply asks a bunch of questoins and fills Staff struct, returns as part of Action struct
func CreateStaff() Action {
	fmt.Println("There are no undos for this action!")

	// filling basic fields
	var t Staff
	fmt.Print("Please enter staff member's name, surname and personal code: ")
	fmt.Scan(&t.Person.Name)
	fmt.Scan(&t.Person.Surname)
	fmt.Scan(&t.Person.PersonalCode)
	fmt.Print("Salary: ")
	fmt.Scan(&t.Salary)

	// filling classes slice
	var a []string
	var inp string
	for {
		fmt.Print("Please enter one or more classrooms, \"exit\" to stop: ")
		fmt.Scan(&inp)
		if inp == "exit" {
			break
		}
		a = append(a, inp)
	}

	//asigning everything
	t.Classroom = a
	fmt.Println("Done!")
	var A Action
	A.Data = t
	A.Action = "create"
	A.Object = "Staff"
	return A
}

// CreateTeacher simply asks a bunch of questoins and fills Student struct, returns as part of Action struct
func CreateStudent() Action {
	fmt.Println("There are no undos for this action!")

	// filling all the fields
	var t Student
	fmt.Print("Please enter student's name, surname and personal code: ")
	fmt.Scan(&t.Person.Name)
	fmt.Scan(&t.Person.Surname)
	fmt.Scan(&t.Person.PersonalCode)
	fmt.Print("Year: ")
	fmt.Scan(&t.Year)
	fmt.Print("Index: ")
	fmt.Scan(&t.Index)
	fmt.Println("Done!")

	// asigning everything
	var A Action
	A.Data = t
	A.Action = "create"
	A.Object = "Student"
	return A
}
