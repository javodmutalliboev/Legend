package developer

import (
	"Legend/database"
	"Legend/password"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CreateAdminCLI function
func CreateAdminCLI() {
	// Check if user is root
	if os.Getenv("ENVIRONMENT") != "development" {
		fmt.Println("You must run this function in development environment!")
		return
	}

	// open database connection
	database := database.DB()
	defer database.Close()

	fmt.Println("Enter the following details to create an admin:")

	fmt.Print("Enter Name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter Surname: ")
	surname, _ := reader.ReadString('\n')

	fmt.Print("Enter Email: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	pass, _ := reader.ReadString('\n')

	// Trim the newline character from each input
	name = strings.TrimSpace(name)
	surname = strings.TrimSpace(surname)
	email = strings.TrimSpace(email)
	pass = strings.TrimSpace(pass)

	pass, err := password.HashPassword(pass)
	if err != nil {
		fmt.Println("Error hashing password!")
		return
	}

	// Insert the admin into the database
	_, err = database.Exec("INSERT INTO admin (name, surname, email, password) VALUES ($1, $2, $3, $4)", name, surname, email, pass)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error inserting admin into database!")
		return
	}

	// Your admin creation logic goes here
	fmt.Println("Admin created successfully!")
}
