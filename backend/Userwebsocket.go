package backend

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// function to handle the new user message (when a new user joins the chat
func handleNewUserMessage(message ServerMessage) {
	users = append(users, ServerUser{Name: message.Users[0].Name})
	Broadcast <- ServerMessage{Type: "users", Users: users}
}

func handleGetUsersMessage(conn *websocket.Conn, message ServerMessage) {
	//open db
	db := OpenDatabase()
	defer db.Close()

	var templist []ServerUser
	// add all users in LoggedInUsers to the templist
	for _, v := range LoggedInUsers {
		templist = append(templist, ServerUser{Nickname: v.Nickname})
	}
	Broadcast <- ServerMessage{Type: "users", Users: templist}
}

func handleGetOfflineUsersMessage(conn *websocket.Conn, message ServerMessage) {
	//open db
	db := OpenDatabase()
	defer db.Close()

	// Get all the users from the database
	offusers, err := GetAllUsers(db)
	if err != nil {
		fmt.Println("Error getting all users")
	}

	var templist []ServerUser

	// add all users in offusers to the templist and remove the users that are already in LoggedInUsers
	for _, v := range offusers {
		var found bool
		for _, v2 := range LoggedInUsers {
			if v.Nickname == v2.Nickname {
				found = true
				break
			}
		}
		if !found {
			templist = append(templist, ServerUser{Nickname: v.Nickname})
		}
	}

	Broadcast <- ServerMessage{Type: "offline_users", Users: templist}
}
