package backend

import (
	"fmt"
	"strconv"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
)

func handleRegisterMessage(conn *websocket.Conn, message ServerMessage) { // we know this works
	//open db
	db := OpenDatabase()
	defer db.Close()

	nickname := message.Data["nickname"]
	password := message.Data["password"]
	confirmpassword := message.Data["cfpassword"]
	email := message.Data["email"]
	age := message.Data["age"]
	gender := message.Data["gender"]
	firsnName := message.Data["firstname"]
	lastName := message.Data["lastname"]

	//convert age to int
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Error converting age to int")
	}

	// Check if the user exist in the database
	if CheckIfUserExist(db, nickname) {
		// If the user exist, redirect to the login page
		fmt.Printf("Server >> User %s tried to register but already exist!\n", nickname)
		// send to frontend that the user already exist
		if conn != nil {
			conn.WriteJSON(ServerMessage{
				Type: "registerResponse",
				Data: map[string]string{
					"register": "false",
					"status":   "User already exist!",
				},
			})
		}
		return
	}
	if CheckIfEmailExist(db, email) {
		// If the user exist, redirect to the login page
		fmt.Printf("Server >> User %s tried to register but already exist!\n", nickname)
		// send to frontend that the user already exist
		if conn != nil {
			conn.WriteJSON(ServerMessage{
				Type: "registerResponse",
				Data: map[string]string{
					"register": "false",
					"status":   "Email already exist!",
				},
			})
		}
		return
	}
	// Checks if the password has a minimum of 4 characters, contains at least one lowercase letter, and one number.
	// to
	if !CheckPasswordStrength(password) {
		// If the password is not strong enough, redirect to the login page
		if conn != nil {
			conn.WriteJSON(ServerMessage{
				Type: "registerResponse",
				Data: map[string]string{
					"register": "false",
					"status":   "Password needs a minimum of 4 characters, contains at least one lowercase letter, and one number.",
				},
			})
		}
		return
	}
	// Check if the password and the password confirm are the same
	if password != confirmpassword {
		// Check if the password and the password confirm are the same
		// If the password and the password confirm are not the same, redirect to the login page
		if conn != nil {
			conn.WriteJSON(ServerMessage{
				Type: "registerResponse",
				Data: map[string]string{
					"register": "false",
					"status":   "Password mismatched!",
				},
			})
		}
		return
	} else {
		password, err = GenerateHash(password)
		if err != nil {
			fmt.Println(err)
		}
		// Insert the user into the database
		RegisterUser(db, nickname, ageInt, gender, firsnName, lastName, email, password)

	}

	conn.WriteJSON(ServerMessage{Type: "registerResponse", Data: map[string]string{"register": "true"}})
}

func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)

	return string(hash), err
}

func handleRegisterResponseMessage(conn *websocket.Conn, message ServerMessage) {
	//open db
	db := OpenDatabase()
	defer db.Close()
	if message.Data["register"] == "true" {
		fmt.Println("Registration successful")
	} else {
		fmt.Println("Registration failed")
	}
}
