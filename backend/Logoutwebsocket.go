package backend

import (
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

func handleLogoutMessage(conn *websocket.Conn, message ServerMessage) {
	db := OpenDatabase()
	defer db.Close()
	modifiedCookie := message.Data["cookie"]
	// ugly fix to get the cookie
	if len(message.Data["cookie"]) > 9 {
		modifiedCookie = modifiedCookie[14:]
	}

	tempUser := ""

	//loop current sesssions
	for _, v := range LoggedInUsers {
		if v.Cookie == modifiedCookie {
			tempUser = v.Nickname
			delete(LoggedInUsers, v.Cookie)

		}
	}

	var templist []ServerUser
	// add all users in LoggedInUsers to the templist
	for _, v := range LoggedInUsers {
		templist = append(templist, ServerUser{Nickname: v.Nickname})
	}

	// Get nickname from session
	var nickname string
	for _, session := range sessions {
		if strings.Compare(session.Cookie, modifiedCookie) == 0 {
			nickname = session.Nickname
		}
	}

	// Remove the user from the LoggedInUsers and sessions map
	delete(LoggedInUsers, nickname)
	delete(sessions, modifiedCookie)
	DeleteSessionByCookie(db, modifiedCookie)

	conn.WriteJSON(ServerMessage{Type: "logoutResponse", Data: map[string]string{"logout": "true"}})
	fmt.Println("Server >> User " + tempUser + " has logged out!")
	Broadcast <- ServerMessage{Type: "users", Users: templist}

}
