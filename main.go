package main

import "github.com/DmitriiKumancev/photography-website/database"

func main() {
	// Initialize Database
	database.Connect("user:root1234@tcp(172.17.0.1:3306)/user_db?parseTime=true")
	database.Migrate()
}
