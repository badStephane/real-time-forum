package backend

import (
	"net/http"
)

// StartFileServers starts the file servers
func StartFileServers() {
	//css fileserver
	fs := http.FileServer(http.Dir("frontend/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	//js fileserver
	fs = http.FileServer(http.Dir("frontend/js"))
	http.Handle("/js/", http.StripPrefix("/js/", fs))
}

// StartHandlers starts the handlers
func StartHandlers() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/check_login", checkLoginHandler)
	http.HandleFunc("/get_logged_in_users", getLoggedInUsersHandler)

}

// StartServer starts the server on port 8080
func StartServer() {
	http.ListenAndServe(":8080", nil)
}
