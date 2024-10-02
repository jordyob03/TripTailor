package DBTests

import (
	"fmt"
	"strconv"
	"time"

	db "backend/db/models"
)

func test() {

	err := db.CreateUserTable()
	if err != nil {
		fmt.Println("Error creating users table:", err)
	} else {
		fmt.Println("Users table created successfully.")
	}

	user1Id, err := db.AddUser("johndoe", "johndoe@example.com", "password123", time.Date(1990, time.March, 15, 0, 0, 0, 0, time.UTC))
	if err != nil {
		fmt.Println("Error adding user 1:", err)
	} else {
		fmt.Printf("User 1 added successfully with ID: %d\n", user1Id)
	}

	user2Id, err := db.AddUser("janedoe", "janedoe@example.com", "securepass456", time.Date(1985, time.December, 20, 0, 0, 0, 0, time.UTC))
	if err != nil {
		fmt.Println("Error adding user 2:", err)
	} else {
		fmt.Printf("User 2 added successfully with ID: %d\n", user2Id)
	}

	_, err = db.AddUser("johndoe", "john.different@example.com", "differentpass789", time.Date(1992, time.January, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		fmt.Println("Expected error for duplicate username:", err)
	}

	_, err = db.AddUser("johnnew", "johndoe@example.com", "differentpass789", time.Date(1992, time.January, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		fmt.Println("Expected error for duplicate email:", err)
	}

	user, err := db.GetUser(strconv.Itoa(user1Id))
	if err != nil {
		fmt.Println("Error retrieving user 1:", err)
	} else {
		fmt.Printf("User 1 retrieved successfully: %+v\n", user)
	}

	user, err = db.GetUser("9999") // Assuming this user ID doesn't exist.
	if err != nil {
		fmt.Println("Expected error for non-existent user:", err)
	} else {
		fmt.Printf("Unexpectedly retrieved user: %+v\n", user)
	}

	users, err := db.GetAllUsers()
	if err != nil {
		fmt.Println("Error retrieving users:", err)
	} else {
		fmt.Printf("Users retrieved successfully: %+v\n", users)
	}

	updateData := map[string]interface{}{
		"username":    "johnupdated",
		"password":    "newpassword123",
		"dateOfBirth": time.Date(1991, time.June, 10, 0, 0, 0, 0, time.UTC),
	}
	err = db.UpdateUser(strconv.Itoa(user1Id), updateData)
	if err != nil {
		fmt.Println("Error updating user 1:", err)
	} else {
		fmt.Println("User 1 updated successfully.")
	}

	updatedUser, err := db.GetUser(strconv.Itoa(user1Id))
	if err != nil {
		fmt.Println("Error retrieving updated user 1:", err)
	} else {
		fmt.Printf("Updated User 1: %+v\n", updatedUser)
	}

	err = db.DeleteUser(strconv.Itoa(user2Id))
	if err != nil {
		fmt.Println("Error deleting user 2:", err)
	} else {
		fmt.Println("User 2 deleted successfully.")
	}

	_, err = db.GetUser(strconv.Itoa(user2Id))
	if err != nil {
		fmt.Println("Expected error for deleted user:", err)
	} else {
		fmt.Println("Unexpectedly retrieved deleted user 2.")
	}

}
