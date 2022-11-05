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

	employees := []schema.User{
		{"Anurag", "22", "8107175060", "Razorpay", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Striver", "25", "23344333", "Google", schema.Address{"san francisco", "california", "USA", "410013"}},
		{"Govind", "27", "23344333", "Microsoft", schema.Address{"pokhara", "pompom", "Nepal", "410013"}},
		{"Vince", "29", "23344333", "Meta", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Loompy", "31", "23344333", "Zoom", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Albert", "32", "23344333", "Dominate", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Rabin", "22", "8107175060", "Razorpay", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Karp", "25", "23344333", "Google", schema.Address{"san francisco", "california", "USA", "410013"}},
		{"Dijkstra", "27", "23344333", "Microsoft", schema.Address{"pokhara", "pompom", "Nepal", "410013"}},
		{"Bellford", "29", "23344333", "Meta", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Lempel", "31", "23344333", "Zoom", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Ziv", "32", "23344333", "Dominate", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Diogo", "22", "8107175060", "Razorpay", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Maradona", "25", "23344333", "Google", schema.Address{"san francisco", "california", "USA", "410013"}},
		{"Zelinsky", "27", "23344333", "Microsoft", schema.Address{"pokhara", "pompom", "Nepal", "410013"}},
		{"Alex", "29", "23344333", "Meta", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Karp", "31", "23344333", "Zoom", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Jack", "32", "23344333", "Dominate", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Hoop", "22", "8107175060", "Razorpay", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Denzel", "25", "23344333", "Google", schema.Address{"san francisco", "california", "USA", "410013"}},
		{"Manacher", "27", "23344333", "Microsoft", schema.Address{"pokhara", "pompom", "Nepal", "410013"}},
		{"Union", "29", "23344333", "Meta", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Find", "31", "23344333", "Zoom", schema.Address{"bangalore", "karnataka", "india", "410013"}},
		{"Gavin", "32", "23344333", "Dominate", schema.Address{"bangalore", "karnataka", "india", "410013"}},
	}

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
