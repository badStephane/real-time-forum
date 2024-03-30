package main

import (
	. "realtime-forum/backend"
)

func main() {
	// Init message when the server starts
	InitMessage()

	// Start the file servers
	StartFileServers()

	// Start the handlers
	StartHandlers()

	// Start the websocket
	StartWebSocketServer()

	// Start the server
	StartServer()
}
