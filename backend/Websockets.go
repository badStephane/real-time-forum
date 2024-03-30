package backend

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Global variables for the WebSocket server
var Broadcast = make(chan ServerMessage)
var users []ServerUser
var categories []ServerCategory
var posts []ServerPost
var clients = make(map[*websocket.Conn]*Session)

// function to start the WebSocket server and go routine for broadcasting messages to all clients
func StartWebSocketServer() {
	upgrader := configureUpgrader()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		cookieValue := r.Header.Get("Cookie")
		if cookieValue != "" {
			fmt.Println("Server >> NEW websocket connection with Cookie value:", cookieValue)
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		handleWebSocketConnection(conn, cookieValue)
	})

	go func() {
		for {
			message := <-Broadcast
			for client := range clients {
				err := client.WriteJSON(message)
				if err != nil {
					log.Println(err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}()
}

// function to configure the upgrader for the WebSocket server
func configureUpgrader() *websocket.Upgrader {
	upgrader := &websocket.Upgrader{}
	upgrader.ReadBufferSize = 1024
	upgrader.WriteBufferSize = 1024
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	return upgrader
}

// function to handle the WebSocket connection and messages
func handleWebSocketConnection(conn *websocket.Conn, cookieValue string) {
	var sessionsMutex sync.Mutex

	if cookieValue != "" {
		for _, v := range LoggedInUsers {
			if v.Cookie == cookieValue[14:] {
				v.WebSocketConn = conn.RemoteAddr().String()
			}
		}
	} else {
		conn.WriteJSON(ServerMessage{Type: "status", Data: map[string]string{"refresh": "true"}})
	}

	// Create a new session for the WebSocket client
	session := Session{
		WebSocketConn: conn.RemoteAddr().String(),
		UserID:        0,
	}

	// Add the session to the clients map
	sessionsMutex.Lock()
	clients[conn] = &session
	sessionsMutex.Unlock()

	// Send initial data to the client
	sendInitialData(conn)

	// Listen for new messages from the client
	for {
		var message ServerMessage
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			delete(clients, conn)
			break
		}

		// Handle the message
		handleWebSocketMessage(conn, message)
	}
}

// function to send the initial data to the client when it connects to the server
func sendInitialData(conn *websocket.Conn) {
	// Send the list of users to the new client
	message := ServerMessage{Type: "users", Users: users}
	conn.WriteJSON(message)

	// Send the list of categories to the new client
	message = ServerMessage{Type: "categories", Categories: categories}
	conn.WriteJSON(message)

	// Send the list of All Posts to the new client
	message = ServerMessage{Type: "posts", Posts: posts}
	conn.WriteJSON(message)
}

// function to handle the messages received from the client
func handleWebSocketMessage(conn *websocket.Conn, message ServerMessage) {

	switch message.Type {
	case "new_user":
		handleNewUserMessage(message)
	case "new_post":
		handleNewPostMessage(message)
	case "get_posts":
		handleGetPostsMessage(conn, message)
	case "get_chat_history":
		handleGetChatHistoryMessage(conn, message)
	case "message":
		handleMessageMessage(conn, message)
	case "login":
		handleLoginMessage(conn, message)
	case "loginResponse":
		handleLoginResponseMessage(conn, message)
	case "logout":
		handleLogoutMessage(conn, message)
	case "register":
		handleRegisterMessage(conn, message)
	case "registerResponse":
		handleRegisterResponseMessage(conn, message)
	case "get_categories":
		handleGetCategoriesMessage(conn, message)
	case "get_comments":
		handleGetCommentsMessage(conn, message)
	case "new_comment":
		handleNewCommentMessage(conn, message)
	case "get_users":
		handleGetUsersMessage(conn, message)
	case "get_offline_users":
		handleGetOfflineUsersMessage(conn, message)
	case "postsByCategory":
		handleGetPostsForCategory(conn, message)
	}
}





