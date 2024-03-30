package backend

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func handleLoginMessage(conn *websocket.Conn, message ServerMessage) {
	//open db
	db := OpenDatabase()
	defer db.Close()
	nickname := message.Data["nickname"]
	password := message.Data["password"]

	if NicknameCheck(nickname) {
		if !CheckIfEmailExist(db, nickname) {
			conn.WriteJSON(ServerMessage{Type: "loginResponse", Data: map[string]string{"status": "error3", "message": "User does not Exist!"}})
			conn.WriteJSON(ServerMessage{Type: "loginResponse", Data: map[string]string{"login": "false"}})
			return
		} else {
			nickname = GetNicknamebyEmail(db, nickname) // transform nickname from email
		}
	}

	// Check if the user exist in the database
	if !CheckIfUserExist(db, nickname) {
		// If the user exist, check if the password is correct
		fmt.Printf("Server >> User %s tried to login but! but does not exist!!!!!\n", nickname)
		conn.WriteJSON(ServerMessage{Type: "loginResponse", Data: map[string]string{"status": "error3", "message": "User does not Exist!"}})
		conn.WriteJSON(ServerMessage{Type: "loginResponse", Data: map[string]string{"login": "false"}})
		return
	}

	if !CheckIfPasswordIsCorrect(db, nickname, password) {
		// If the user exist but the password is incorrect
		fmt.Printf("Server >> User %s tried to login but the password is incorrect!\n", nickname)
		conn.WriteJSON(ServerMessage{Type: "loginResponse", Data: map[string]string{"status": "error4", "message": "Incorrect password"}})
		conn.WriteJSON(ServerMessage{Type: "loginResponse", Data: map[string]string{"login": "false"}})
		return
	}

	if UserLoggedIn(nickname) {
		fmt.Println("Server >> User already logged in logging the user out from other endpoints")

		//loop through the loggedInUsers map and remove the user
		for key, value := range LoggedInUsers {
			if value.Nickname == nickname {
				// alert the user thats get forced logged out
				for client := range clients {
					if client.RemoteAddr().String() == value.WebSocketConn {
						err := client.WriteJSON(ServerMessage{Type: "loginResponse", Data: map[string]string{"status": "error2", "message": "User already logged in"}})
						if err != nil {
							log.Println(err)
							client.Close()
							delete(clients, client)
						}
					}
				}
				delete(LoggedInUsers, key)
			}
		}

		//loop through the clients map and remove the user
		for key, value := range clients {
			if value.Nickname == nickname {
				delete(clients, key)
			}
		}

		//loop through the sessions map and remove the user
		for key, value := range sessions {
			if value.Nickname == nickname {
				delete(sessions, key)
			}
		}

		conn.WriteJSON(ServerMessage{Type: "loginResponse", Data: map[string]string{"status": "error", "message": "User already logged in"}})
		return
	}

	// Set the session and get the session token
	sessionToken := SetClientCookieWithSessionToken(conn, db, nickname)

	// Add the user session to the loggedInUsers map
	session := &Session{
		Nickname:      nickname,
		Cookie:        sessionToken,
		ExpiredTime:   time.Now().Add(1 * time.Hour), 
		WebSocketConn: conn.RemoteAddr().String(),
	}
	LoggedInUsers[sessionToken] = session

	conn.WriteJSON(ServerMessage{Type: "loginResponse", Data: map[string]string{"login": "true", "nickname": nickname}})
	fmt.Printf("Server >> User %s has logged in!\n", nickname)

}

func handleLoginResponseMessage(conn *websocket.Conn, message ServerMessage) {

	fmt.Println("When am I called?")
}
