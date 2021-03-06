package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

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
type Action struct {
	Action string      `json:"actoin"`
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

	// Create json request
	msg := []byte("{\"action\":\"create\",\"object\":\"Teacher\",\"data\":{\"subjeect\":\"Math\",\"salary\":2345,\"classroom\":[\"CL-001\",\"CL-002\",\"CL-005\"],\"person\":{\"name\":\"Ivan\",\"surname\":\"Popov\",\"personalCode\":\"123422-43235\"}}}")

	for {
		var inp string
		var msg []byte
		fmt.Print("Please select action (select/create/exit): ")
		fmt.Scan(&inp)
		switch inp {
		case "select":
		case "create":
			msg = GetJsonCreate()
		case "exit":
			fmt.Println("Exiting...")
			conn.Close()
			os.Exit(0)
		}
		// Send msg
		conn.Write(msg)

		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Response: %s\n", string(buf[:n]))
	}

}

func GetJsonCreate() []byte {
	var inp string
	var a byte
	var msg []byte
	for a == 0 { // for will run 1 time, unless default case is trigered
		a++
		fmt.Print("Please select job (Teacher/Student/Staff): ")
		fmt.Scan(&inp)
		inp = strings.ToLower(inp)
		switch inp {
		case "teacher":
			msg, _ = json.Marshal(CreateTeacher())
		case "student":
			msg, _ = json.Marshal(CreateStudent())
		case "staff":
			msg, _ = json.Marshal(CreateStaff())
		default:
			continue
		}
	}
	return msg
}

func CreateTeacher() Action {
	fmt.Println("Creating teacher")
	fmt.Println("There are no undos for this action!")
	var t Teacher
	fmt.Print("Please enter teacher's name, surname and personal code: ")
	fmt.Scan(&t.Person.Name)
	fmt.Scan(&t.Person.Surname)
	fmt.Scan(&t.Person.PersonalCode)
	fmt.Print("Subject: ")
	fmt.Scan(&t.Subject)
	fmt.Print("Salary: ")
	fmt.Scan(&t.Salary)

	var a []string
	var inp string
	for {
		fmt.Print("Please enter one or more classroom, \"exit\" to stop")
		fmt.Scan(&inp)
		if inp == "exit" {
			break
		}
		a = append(a, inp)
	}
	t.Classroom = a
	fmt.Println("Done!")
	var A Action
	A.Data = t
	A.Action = "create"
	A.Object = "Teacher"
	return A
}

func CreateStaff() Action {
	fmt.Println("Creating staff")
	fmt.Println("There are no undos for this action!")
	var t Staff
	fmt.Print("Please enter staff member's name, surname and personal code: ")
	fmt.Scan(&t.Person.Name)
	fmt.Scan(&t.Person.Surname)
	fmt.Scan(&t.Person.PersonalCode)
	fmt.Print("Salary: ")
	fmt.Scan(&t.Salary)

	var a []string
	var inp string
	for {
		fmt.Print("Please enter one or more classroom, \"exit\" to stop")
		fmt.Scan(&inp)
		if inp == "exit" {
			break
		}
		a = append(a, inp)
	}
	t.Classroom = a
	fmt.Println("Done!")
	var A Action
	A.Data = t
	A.Action = "create"
	A.Object = "Staff"
	return A
}

func CreateStudent() Action {
	fmt.Println("Creating student")
	fmt.Println("There are no undos for this action!")
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
	var A Action
	A.Data = t
	A.Action = "create"
	A.Object = "Student"
	return A
}
