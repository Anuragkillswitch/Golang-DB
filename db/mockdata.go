package main

import "db/schema"

func GetMockData() []schema.User {
	employees := []schema.User{
		{Name: "Anurag", Age: "22", Contact: "8107175060", Company: "Razorpay", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Striver", Age: "25", Contact: "23344333", Company: "Google", Address: schema.Address{City: "san francisco", State: "california", Country: "USA", Pincode: "410013"}},
		{Name: "Govind", Age: "27", Contact: "23344333", Company: "Microsoft", Address: schema.Address{City: "pokhara", State: "pompom", Country: "Nepal", Pincode: "410013"}},
		{Name: "Vince", Age: "29", Contact: "23344333", Company: "Meta", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Loompy", Age: "31", Contact: "23344333", Company: "Zoom", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Albert", Age: "32", Contact: "23344333", Company: "Dominate", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Rabin", Age: "22", Contact: "8107175060", Company: "Razorpay", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Karp", Age: "25", Contact: "23344333", Company: "Google", Address: schema.Address{City: "san francisco", State: "california", Country: "USA", Pincode: "410013"}},
		{Name: "Dijkstra", Age: "27", Contact: "23344333", Company: "Microsoft", Address: schema.Address{City: "pokhara", State: "pompom", Country: "Nepal", Pincode: "410013"}},
		{Name: "Bellford", Age: "29", Contact: "23344333", Company: "Meta", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Lempel", Age: "31", Contact: "23344333", Company: "Zoom", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Ziv", Age: "32", Contact: "23344333", Company: "Dominate", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Diogo", Age: "22", Contact: "8107175060", Company: "Razorpay", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Maradona", Age: "25", Contact: "23344333", Company: "Google", Address: schema.Address{City: "san francisco", State: "california", Country: "USA", Pincode: "410013"}},
		{Name: "Zelinsky", Age: "27", Contact: "23344333", Company: "Microsoft", Address: schema.Address{City: "pokhara", State: "pompom", Country: "Nepal", Pincode: "410013"}},
		{Name: "Alex", Age: "29", Contact: "23344333", Company: "Meta", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Karp", Age: "31", Contact: "23344333", Company: "Zoom", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Jack", Age: "32", Contact: "23344333", Company: "Dominate", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Hoop", Age: "22", Contact: "8107175060", Company: "Razorpay", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Denzel", Age: "25", Contact: "23344333", Company: "Google", Address: schema.Address{City: "san francisco", State: "california", Country: "USA", Pincode: "410013"}},
		{Name: "Manacher", Age: "27", Contact: "23344333", Company: "Microsoft", Address: schema.Address{City: "pokhara", State: "pompom", Country: "Nepal", Pincode: "410013"}},
		{Name: "Union", Age: "29", Contact: "23344333", Company: "Meta", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Find", Age: "31", Contact: "23344333", Company: "Zoom", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
		{Name: "Gavin", Age: "32", Contact: "23344333", Company: "Dominate", Address: schema.Address{City: "bangalore", State: "karnataka", Country: "india", Pincode: "410013"}},
	}
	return employees
}
