package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// indexHandler handles the index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/index.html")
}

// loginHandler handles the login endpoint
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	db := OpenDatabase()
	defer CloseDatabase(db)

	if r.URL.Path != "/login" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
}

// logoutHandler handles the logout endpoint
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	var templist []ServerUser

	for _, v := range LoggedInUsers {
		templist = append(templist, ServerUser{Nickname: v.Nickname})
	}

	Broadcast <- ServerMessage{Type: "users", Users: templist}
}

// signupHandler handles the signup endpoint
func signupHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenDatabase()
	defer CloseDatabase(db)

	if r.URL.Path != "/signup" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Println("Method not allowed")
		return
	}

	nickname := r.FormValue("nickname")
	ageStr := r.FormValue("age")
	
	age, _ := strconv.Atoi(ageStr)
	gender := r.FormValue("gender")
	fname := r.FormValue("firstname")
	lname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("confpassword")

	if CheckIfUserExist(db, nickname) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		// Check if the password and the password confirm are the same
		if password != passwordConfirm {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Insert the user into the database
		RegisterUser(db, nickname, age, gender, fname, lname, email, password)

		// Redirect to the chat page
		http.Redirect(w, r, "/chat", http.StatusSeeOther)
	}
}

func checkLoginHandler(w http.ResponseWriter, r *http.Request) {
	logged_in := false
	db := OpenDatabase()
	defer CloseDatabase(db)

	sessionToken := GetSessionTokenFromCookie(r)

	if CheckIfSessionTokenIsValid(sessionToken) {
		logged_in = true
	} else {
		logged_in = false
	}

	response := map[string]bool{"logged_in": logged_in}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getLoggedInUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []string
	for user := range LoggedInUsers {
		users = append(users, user)
	}
	response := map[string][]string{"users": users}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
