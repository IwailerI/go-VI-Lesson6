package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Person contains Teacher/Student/Staff
type Person interface {
	GetID() float64
}

// Action describes request
type Action struct {
	Action  string `json:"action"`
	ObjName string `json:"object"`
}

// Teacher contains all information about teacher
type Teacher struct {
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

// GetID returns ID
func (t Teacher) GetID() float64 {
	return t.ID
}

// Student contains all information about student
type Student struct {
	ID     float64 `json:"id"`
	Year   float64 `json:"year"`
	Index  string  `json:"index"`
	Person struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

// GetID returns ID
func (t Student) GetID() float64 {
	return t.ID
}

// Staff contains all information about staff
type Staff struct {
	ID        float64  `json:"id"`
	Salary    float64  `json:"salary"`
	Classroom []string `json:"classroom"`
	Person    struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

// GetID returns ID
func (t Staff) GetID() float64 {
	return t.ID
}

// DefinedAction is one of 4 action: CRUD
type DefinedAction interface {
	GetFromJSON([]byte)
	Process()
}

// GetCreate a
func (t Teacher) GetCreate() DefinedAction {
	return &CreateTeacher{}
}

// GetRead a
func (t Teacher) GetRead() DefinedAction {
	return &ReadTeacher{}
}

// GetUpdate a
func (t Teacher) GetUpdate() DefinedAction {
	return &UpdateTeacher{}
}

// GetDelete a
func (t Teacher) GetDelete() DefinedAction {
	return &DeleteTeacher{}
}

// GetCreate a
func (t Student) GetCreate() DefinedAction {
	return &CreateStudent{}
}

// GetRead a
func (t Student) GetRead() DefinedAction {
	return &ReadStudent{}
}

// GetUpdate a
func (t Student) GetUpdate() DefinedAction {
	return &UpdateStudent{}
}

// GetDelete a
func (t Student) GetDelete() DefinedAction {
	return &DeleteStudent{}
}

// GetCreate a
func (t Staff) GetCreate() DefinedAction {
	return &CreateStaff{}
}

// GetRead a
func (t Staff) GetRead() DefinedAction {
	return &ReadStaff{}
}

// GetUpdate a
func (t Staff) GetUpdate() DefinedAction {
	return &UpdateStaff{}
}

// GetDelete a
func (t Staff) GetDelete() DefinedAction {
	return &DeleteStaff{}
}

// CreateTeacher a
type CreateTeacher struct {
	T Teacher `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *CreateTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// ReadTeacher a
type ReadTeacher struct {
	T Teacher `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *ReadTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// UpdateTeacher a
type UpdateTeacher struct {
	T Teacher `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *UpdateTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// DeleteTeacher a
type DeleteTeacher struct {
	T Teacher `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *DeleteTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// CreateStudent a
type CreateStudent struct {
	T Student `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *CreateStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// ReadStudent a
type ReadStudent struct {
	T Student `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *ReadStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// UpdateStudent a
type UpdateStudent struct {
	T Student `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *UpdateStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// DeleteStudent a
type DeleteStudent struct {
	T Student `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *DeleteStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// CreateStaff a
type CreateStaff struct {
	T Staff `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *CreateStaff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// ReadStaff a
type ReadStaff struct {
	T Staff `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *ReadStaff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// UpdateStaff a
type UpdateStaff struct {
	T Staff `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *UpdateStaff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// DeleteStaff a
type DeleteStaff struct {
	T Staff `json:"data"`
}

// GetFromJSON unmarshals data from rawData into itself
func (action *DeleteStaff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		panic(err)
	}
}

// Process a
func (action *CreateTeacher) Process() {
	IDCOUNTER++
	fmt.Printf("Creating teacher, id:%.0f\n", IDCOUNTER)
	action.T.ID = IDCOUNTER
	DATABASE = append(DATABASE, action.T)
}

// Process a
func (action *ReadTeacher) Process() {
	fmt.Printf("Reading teacher, id:%0.f\n", action.T.ID)
	for _, n := range DATABASE {
		if n.GetID() == action.T.ID {
			fmt.Println(n)
			return
		}
	}
	fmt.Println("Teacher not found")
}

// Process a
func (action *UpdateTeacher) Process() {
	fmt.Printf("Updating teacher, id:%0.f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			DATABASE[i] = action.T
			return
		}
	}
	fmt.Println("Teacher not found")
}

// Process a
func (action *DeleteTeacher) Process() {
	fmt.Printf("Deleting teacher, id:%0.f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			DATABASE[i], DATABASE[len(DATABASE)-1] = DATABASE[len(DATABASE)-1], DATABASE[i]
			DATABASE = DATABASE[:len(DATABASE)-1]
			return
		}
	}
}

// Process a
func (action *CreateStudent) Process() {
	IDCOUNTER++
	fmt.Printf("Creating student, id:%.0f\n", IDCOUNTER)
	action.T.ID = IDCOUNTER
	DATABASE = append(DATABASE, action.T)
}

// Process a
func (action *ReadStudent) Process() {
	fmt.Printf("Reading student, id:%0.f\n", action.T.ID)
	for _, n := range DATABASE {
		if n.GetID() == action.T.ID {
			fmt.Println(n)
			return
		}
	}
	fmt.Println("Student not found")
}

// Process a
func (action *UpdateStudent) Process() {
	fmt.Printf("Updating student, id:%0.f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			DATABASE[i] = action.T
			return
		}
	}
	fmt.Println("Student not found")
}

// Process a
func (action *DeleteStudent) Process() {
	fmt.Printf("Deleting student, id:%0.f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			DATABASE[i], DATABASE[len(DATABASE)-1] = DATABASE[len(DATABASE)-1], DATABASE[i]
			DATABASE = DATABASE[:len(DATABASE)-1]
			return
		}
	}
}

// Process a
func (action *CreateStaff) Process() {
	IDCOUNTER++
	fmt.Printf("Creating staff, id:%.0f\n", IDCOUNTER)
	action.T.ID = IDCOUNTER
	DATABASE = append(DATABASE, action.T)
}

// Process a
func (action *ReadStaff) Process() {
	fmt.Printf("Reading staff, id:%0.f\n", action.T.ID)
	for _, n := range DATABASE {
		if n.GetID() == action.T.ID {
			fmt.Println(n)
			return
		}
	}
	fmt.Println("Staff not found")
}

// Process a
func (action *UpdateStaff) Process() {
	fmt.Printf("Updating staff, id:%0.f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			DATABASE[i] = action.T
			return
		}
	}
	fmt.Println("Staff not found")
}

// Process a
func (action *DeleteStaff) Process() {
	fmt.Printf("Deleting staff, id:%0.f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			DATABASE[i], DATABASE[len(DATABASE)-1] = DATABASE[len(DATABASE)-1], DATABASE[i]
			DATABASE = DATABASE[:len(DATABASE)-1]
			return
		}
	}
}

// GeneralObject a
type GeneralObject interface {
	GetCreate() DefinedAction
	GetRead() DefinedAction
	GetUpdate() DefinedAction
	GetDelete() DefinedAction
}

// DATABASE is main data sotrage here
var DATABASE []Person

// IDCOUNTER stores current id, increment each time object is created
var IDCOUNTER float64

func main() {
	fin, err := os.Open("data2.dat")
	if err != nil {
		panic(err)
	}
	defer fin.Close()
	in := bufio.NewScanner(fin)
	for in.Scan() {
		dat := []byte(in.Text())

		var act Action
		err = json.Unmarshal(dat, &act)
		if err != nil {
			panic(err)
		}

		var obj GeneralObject
		switch act.ObjName {
		case "Teacher":
			obj = &Teacher{}
		case "Student":
			obj = &Student{}
		case "Staff":
			obj = &Staff{}
		}

		var task DefinedAction
		switch act.Action {
		case "create":
			task = obj.GetCreate()
		case "read":
			task = obj.GetRead()
		case "update":
			task = obj.GetUpdate()
		case "delete":
			task = obj.GetDelete()
		}

		task.GetFromJSON(dat)
		task.Process()
	}
}
