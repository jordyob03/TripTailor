package DBtests

import (
	db "backend/db/models"
	"bufio"
	"fmt"
	"os"
)

func DBtest() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1: Add User")
		fmt.Println("2: Get User by Email")
		fmt.Println("3: Get All Users")
		fmt.Println("4: Update User")
		fmt.Println("5: Delete User")
		fmt.Println("6: Exit")

		if !scanner.Scan() {
			break
		}
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter email: ")
			scanner.Scan()
			email := scanner.Text()
			fmt.Print("Enter password: ")
			scanner.Scan()
			password := scanner.Text()
			err := db.AddUser(email, password)
			if err != nil {
				fmt.Println("Error adding user:", err)
			} else {
				fmt.Println("User added successfully!")
			}

		case "2":
			fmt.Print("Enter user email: ")
			scanner.Scan()
			email := scanner.Text()
			user, err := db.GetUser(email) // Update to use GetUserByEmail
			if err != nil {
				fmt.Println("Error retrieving user:", err)
			} else {
				fmt.Printf("User: %+v\n", user)
			}

		case "3":
			users, err := db.GetAllUsers()
			if err != nil {
				fmt.Println("Error retrieving users:", err)
			} else {
				fmt.Println("All Users:")
				for _, user := range users {
					fmt.Printf("Email: %s\n", user.Email) // Updated to display email
				}
			}

		case "4":
			fmt.Print("Enter user email to update: ")
			scanner.Scan()
			email := scanner.Text()
			fmt.Print("Enter new email: ")
			scanner.Scan()
			newEmail := scanner.Text()
			fmt.Print("Enter new password: ")
			scanner.Scan()
			password := scanner.Text()
			err := db.UpdateUser(email, map[string]interface{}{
				"email":    newEmail,
				"password": password,
			}) // Update function to use email
			if err != nil {
				fmt.Println("Error updating user:", err)
			} else {
				fmt.Println("User updated successfully!")
			}

		case "5":
			fmt.Print("Enter user email to delete: ")
			scanner.Scan()
			email := scanner.Text()
			err := db.DeleteUser(email) // Update to use email
			if err != nil {
				fmt.Println("Error deleting user:", err)
			} else {
				fmt.Println("User deleted successfully!")
			}

		case "6":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
