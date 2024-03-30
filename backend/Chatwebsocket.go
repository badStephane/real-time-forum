package backend

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)



// function to handle the get chat history message (when a user wants to see the chat history)
func handleGetChatHistoryMessage(conn *websocket.Conn, message ServerMessage) {

	db := OpenDatabase()
	defer db.Close()
	jsonStr, _ := json.Marshal(message.User)
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println(err)
	}
	nickname := data["nickname"].(string)
	userID := GetUserID(db, nickname)
	from := GetUserID(db, message.From)

	chatHistory := GetChatHistory(userID, from, message.Start)

	// Send the conversation history to the client
	conn.WriteJSON(ServerMessage{Type: "chat_history", ChatHistory: chatHistory})
}

// function to handle the message message (when a user sends a message to another user)
func handleMessageMessage(conn *websocket.Conn, message ServerMessage) {

	message.Nickname = message.From
	db := OpenDatabase()
	defer db.Close()

	historyTo := GetUserID(db, message.To)
	historyFrom := GetUserID(db, message.From)
	AddMessageToHistory(historyFrom, historyTo, message.Text)

	for _, value := range LoggedInUsers {
		if value.Nickname == message.To {
			message.To = value.WebSocketConn
		}
		if value.Nickname == message.From {
			message.From = value.WebSocketConn
		}
	}

	for client := range clients {
		if client.RemoteAddr().String() == message.To {
			err := client.WriteJSON(message)
			if err != nil {
				log.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func handleTypingMessage(conn *websocket.Conn, message ServerMessage) {
	//open db
	db := OpenDatabase()
	defer db.Close()
	// Get the user that sent the message
	user := message.Data["to"]

	// send typing signal to the client
	Broadcast <- ServerMessage{Type: "typing", Data: map[string]string{"to": user, "from": message.Data["from"]}}
}

func handleStopTypingMessage(conn *websocket.Conn, message ServerMessage) {
	//open db
	db := OpenDatabase()
	defer db.Close()
	// Get the user that sent the message
	user := message.Data["to"]

	// send typing signal to the client
	Broadcast <- ServerMessage{Type: "stopTyping", Data: map[string]string{"to": user, "from": message.Data["from"]}}
}
