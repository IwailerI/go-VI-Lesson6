package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Person contains Teacher/Student/Staff
type Person interface {
	GetID() float64
	Lock()
	Unlock()
	GetJob() string
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
		ch           chan bool
	} `json:"person"`
}

// GetID returns ID
func (t Teacher) GetID() float64 {
	return t.ID
}

// Lock is makeshift mutex
func (t Teacher) Lock() {
	<-t.Person.ch
}

// Unlock is makeshift mutex
func (t Teacher) Unlock() {
	t.Person.ch <- true
}

// GetJob returns "Teacher"
func (t Teacher) GetJob() string {
	return "Teacher"
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
		ch           chan bool
	} `json:"person"`
}

// GetID returns ID
func (t Student) GetID() float64 {
	return t.ID
}

// Lock is makeshift mutex
func (t Student) Lock() {
	<-t.Person.ch
}

// Unlock is makeshift mutex
func (t Student) Unlock() {
	t.Person.ch <- true
}

// GetJob returns "Student"
func (t Student) GetJob() string {
	return "Student"
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
		ch           chan bool
	} `json:"person"`
}

// GetID returns ID
func (t Staff) GetID() float64 {
	return t.ID
}

// Lock is makeshift mutex
func (t Staff) Lock() {
	<-t.Person.ch
}

// Unlock is makeshift mutex
func (t Staff) Unlock() {
	t.Person.ch <- true
}

// GetJob returns "Staff"
func (t Staff) GetJob() string {
	return "Staff"
}

// DefinedAction is one of 4 action: CRUD
type DefinedAction interface {
	GetFromJSON([]byte)
	Process() []byte
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

// In these methods all the magic hapens, my sanity won't let me comment everytinf in there
// they are pretty easy to understand, a lot of copypasta

// Process a
func (action *CreateTeacher) Process() []byte {
	IDCOUNTER++
	fmt.Printf("Creating teacher, id:%.0f\n", IDCOUNTER)
	action.T.ID = IDCOUNTER
	resp := fmt.Sprint(IDCOUNTER)
	action.T.Person.ch = make(chan bool, 1)
	action.T.Unlock()
	DATABASE = append(DATABASE, action.T)
	return []byte(resp)
}

// Process a
func (action *ReadTeacher) Process() []byte {
	fmt.Printf("Reading teacher, id:%.0f\n", action.T.ID)
	for _, n := range DATABASE {
		if n.GetID() == action.T.ID {
			resp, _ := json.Marshal(n)
			return resp
		}
	}
	fmt.Println("Teacher not found")
	return []byte("Teacher not found")
}

// Process a
func (action *UpdateTeacher) Process() []byte {
	fmt.Printf("Updating teacher, id:%.0f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			DATABASE[i].Lock()
			action.T.Person.ch = DATABASE[i].(Teacher).Person.ch // transmit mutex channel to new actoin.t
			DATABASE[i] = action.T
			DATABASE[i].Unlock()
			return []byte("Succes")
		}
	}
	fmt.Println("Teacher not found")
	return []byte("Teacher not found")
}

// Process a
func (action *DeleteTeacher) Process() []byte {
	fmt.Printf("Deleting teacher, id:%.0f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			lock()
			DATABASE[i].Lock()
			DATABASE[i], DATABASE[len(DATABASE)-1] = DATABASE[len(DATABASE)-1], DATABASE[i]
			DATABASE = DATABASE[:len(DATABASE)-1]
			unlock()
			return []byte("Succes")
		}
	}
	fmt.Println("Teacher not found")
	return []byte("Teacher not found")
}

// Process a
func (action *CreateStudent) Process() []byte {
	IDCOUNTER++
	fmt.Printf("Creating student, id:%.0f\n", IDCOUNTER)
	action.T.ID = IDCOUNTER
	resp := fmt.Sprint(IDCOUNTER)
	action.T.Person.ch = make(chan bool, 1)
	action.T.Unlock()
	DATABASE = append(DATABASE, action.T)
	return []byte(resp)
}

// Process a
func (action *ReadStudent) Process() []byte {
	fmt.Printf("Reading student, id:%.0f\n", action.T.ID)
	for _, n := range DATABASE {
		if n.GetID() == action.T.ID {
			resp, _ := json.Marshal(n)
			return resp
		}
	}
	fmt.Println("Student not found")
	return []byte("Student not found")
}

// Process a
func (action *UpdateStudent) Process() []byte {
	fmt.Printf("Updating student, id:%.0f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			DATABASE[i].Lock()
			action.T.Person.ch = DATABASE[i].(Student).Person.ch // transmit mutex channel to new action.t
			DATABASE[i] = action.T
			DATABASE[i].Unlock()
			return []byte("Succes")
		}
	}
	fmt.Println("Student not found")
	return []byte("Student not found")
}

// Process a
func (action *DeleteStudent) Process() []byte {
	fmt.Printf("Deleting student, id:%.0f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			lock()
			DATABASE[i].Lock()
			DATABASE[i], DATABASE[len(DATABASE)-1] = DATABASE[len(DATABASE)-1], DATABASE[i]
			DATABASE = DATABASE[:len(DATABASE)-1]
			unlock()
			return []byte("Succes")
		}
	}
	fmt.Println("Student not found")
	return []byte("Student not found")
}

// Process a
func (action *CreateStaff) Process() []byte {
	IDCOUNTER++
	fmt.Printf("Creating staff, id:%.0f\n", IDCOUNTER)
	action.T.ID = IDCOUNTER
	resp := fmt.Sprint(IDCOUNTER)
	action.T.Person.ch = make(chan bool, 1)
	action.T.Unlock()
	DATABASE = append(DATABASE, action.T)
	return []byte(resp)
}

// Process a
func (action *ReadStaff) Process() []byte {
	fmt.Printf("Reading staff, id:%.0f\n", action.T.ID)
	for _, n := range DATABASE {
		if n.GetID() == action.T.ID {
			resp, _ := json.Marshal(n)
			return resp
		}
	}
	fmt.Println("Staff not found")
	return []byte("Staff not found")
}

// Process a
func (action *UpdateStaff) Process() []byte {
	fmt.Printf("Updating staff, id:%.0f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			DATABASE[i].Lock()
			action.T.Person.ch = DATABASE[i].(Staff).Person.ch // transmit mutex channel to new action.t
			DATABASE[i] = action.T
			DATABASE[i].Unlock()
			return []byte("Succes")
		}
	}
	fmt.Println("Staff not found")
	return []byte("Staff not found")
}

// Process a
func (action *DeleteStaff) Process() []byte {
	fmt.Printf("Deleting staff, id:%.0f\n", action.T.ID)
	for i, n := range DATABASE {
		if n.GetID() == action.T.ID {
			lock()
			DATABASE[i].Lock()
			DATABASE[i], DATABASE[len(DATABASE)-1] = DATABASE[len(DATABASE)-1], DATABASE[i]
			DATABASE = DATABASE[:len(DATABASE)-1]
			unlock()
			return []byte("Succes")
		}
	}
	fmt.Println("Staff not found")
	return []byte("Staff not found")
}

// GeneralObject a
type GeneralObject interface {
	GetCreate() DefinedAction
	GetRead() DefinedAction
	GetUpdate() DefinedAction
	GetDelete() DefinedAction
}

func getJob(ID float64) []byte {
	fmt.Printf("Searching for job of %.0f\n", ID)
	for _, n := range DATABASE {
		if n.GetID() == ID {
			return []byte(n.GetJob())
		}
	}
	fmt.Println("ID not found")
	return []byte("ID not found")
}

// DATABASE is main data sotrage here
var DATABASE []Person

// IDCOUNTER stores current id, increment each time object is created
var IDCOUNTER float64

// global mutex is used when deleting elemnts from database, so noonone can acces it while elemnts are shifting
var mutex chan bool

func init() {
	mutex = make(chan bool, 1) // asigning mutex
}

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		fmt.Println("Listing IDs")
		resp := GetIDS()
		io.WriteString(w, resp)
	} else if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return
		}
		resp := HandleConn(data)
		io.WriteString(w, resp)
	} else {
		w.WriteHeader(405)
	}

}

func main() {
	http.HandleFunc("/", Handler)
	unlock()
	err := http.ListenAndServe(":5555", nil)
	panic(err)
}

func GetIDS() string {
	fmt.Println("Getting IDS")
	var resp string
	resp = "\n"
	for _, n := range DATABASE {
		resp += strconv.Itoa(int(n.GetID())) + "\n"
	}
	return resp
}

// HandleConn is self explanatory
func HandleConn(data []byte) string {
	// Decode json request
	var act Action
	err := json.Unmarshal(data, &act)
	if err != nil {
		// if request is invalid, allert client and wait continue
		return "Invalid request"
	}

	// Get job label or object that we will be working with
	var obj GeneralObject
	switch act.ObjName {
	case "Teacher":
		obj = &Teacher{}
	case "Student":
		obj = &Student{}
	case "Staff":
		obj = &Staff{}
	case "Unknown":
		// That means that act.Actions contains ID
	}

	var resp string        // we will send back to client
	var task DefinedAction //task we will be executing

	// deciding wich action we will nedd to execute
	switch act.Action {
	case "create":
		task = obj.GetCreate()
	case "read":
		task = obj.GetRead()
	case "update":
		task = obj.GetUpdate()
	case "delete":
		task = obj.GetDelete()
	default: // act.Action contains id
		// client wants to know job of certain id
		var id float64

		// parsing id, that client sent
		id, _ = strconv.ParseFloat(act.Action, 64)

		// filling resp
		resp = string(getJob(id))
	}

	if len(resp) == 0 {
		// Execute json request, if we didn't fill resp yet

		// decode json request
		task.GetFromJSON(data)

		// execute request
		resp = string(task.Process())
	}

	if len(resp) == 0 {
		// failsave
		// if resp was nil, client on the other side would be waiting for eterniry
		// because of this he will recieve "null"
		resp = "null"
	}

	// Respond
	fmt.Printf("Responce: %s\n", string(resp))
	return resp

}

// global mutex lock
func lock() {
	<-mutex
}

// global mutex unlock
func unlock() {
	mutex <- true
}
