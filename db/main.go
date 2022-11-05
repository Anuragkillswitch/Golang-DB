package main

import (
	"db/schema"
	"encoding/json"
	"fmt"
)

func main() {

	dir := "./"

	db, err := New(dir, nil)

	if err != nil {
		fmt.Println("Error", err)
	}

	employees := GetMockData()

	for _, value := range employees {
		db.Write("users", value.Name, schema.User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(records)

	allUsers := []schema.User{}

	for _, record := range records {
		employeeFound := schema.User{}
		if err := json.Unmarshal([]byte(record), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allUsers = append(allUsers, employeeFound)
	}

	fmt.Println(allUsers)

	if err := db.Delete("user", "Anitej"); err != nil {
		fmt.Println("Error", err)
	}

	if err := db.Delete("user", ""); err != nil {
		fmt.Println("Error", err)
	}

}
