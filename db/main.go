package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"github.com/jcelliott/lumber" // Logger package
)

const Version = "1.0.0"

type(
    Logger interface{
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}
	Driver struct{
		mutex sync.Mutex
		mutexes map[string]*sync.Mutex
		dir string
		log Logger
	}
)

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

type Options struct{
	Logger 
}

func New()(){

}

func (d *Driver) Write() error {
   
}

func (d *Driver) Read() error {
  
}

func (d *Driver) ReadAll()() {

}

func (d *Driver) Delete() error{

}

func (d *Driver) getOrCreateMutex() *sync.Mutex{

}

func main() {

	dir := "./"

	db, err := New(dir, nil)

	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"Anurag", "22", "32334", "Palantir", Address{"Bangalore", "Karnataka", "India", "5600001"}},
		{"Anitej", "32", "2234", "Razorpay", Address{"Kolkata", "West Bengal", "India", "700025"}},
		{"Murali", "42", "2234", "Amazon", Address{"Hardoi", "Uttar Pradesh", "India", "241001"}},
		{"JSON", "52", "6234", "Statham", Address{"Varkala", "Kerala", "India", "630001"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.RealAll("users")

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(records)
	}

	allUsers := []User{}

	for _, record := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(record), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allUsers = append(allUsers, employeeFound)
	}

	fmt.Println((allUsers))

	// if err := db.Delete("user", "Anitej"); err != nil {
	// 	fmt.Println("Error", err)
	// }

	// if err := db.Delete("user", ""); err != nil {
	// 	fmt.Println("Error", err)
	// }

}
