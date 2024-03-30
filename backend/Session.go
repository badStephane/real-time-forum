package backend

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

var sessions = make(map[string]Session)

func SetClientCookieWithSessionToken(conn *websocket.Conn, db *sql.DB, nickname string) string {
	u2, err := uuid.NewV4()
	if err != nil {
		return "500 INTERNAL SERVER ERROR: GENERATING SESSION TOKEN FAILED"
	}
	sessionToken := u2.String()
	expiredTime := time.Now().Add(3600 * 24 * 3 * time.Second)

	for oldSessionToken, oldSession := range sessions {
		if strings.Compare(oldSession.Nickname, nickname) == 0 {
			delete(sessions, oldSessionToken)
		}
	}

	sessions[sessionToken] = Session{Nickname: nickname, Cookie: sessionToken, ExpiredTime: expiredTime}

	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiredTime,
		Path:    "/",
	}

	err = InsertSessionIntoDB(db, nickname, sessionToken, expiredTime)
	if err != nil {
		fmt.Println("Error inserting session into database:", err)
	}

	if conn != nil {
		conn.WriteJSON(ServerMessage{
			Type: "loginResponse",
			Data: map[string]string{
				"login":  "true",
				"cookie": cookie.String(),
			},
		})
	}

	return sessionToken
}

func sessionExpired(session Session) bool {
	return session.ExpiredTime.Before(time.Now())
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("session_token")
	if err == http.ErrNoCookie {
		return "401 UNAUTHORIZED: CLIENT COOKIE NOT SET OR SESSION EXPIRED"
	}
	if err != nil {
		return "400 BAD REQUEST: REQUEST NOT ALLOWED"
	}
	sessionToken := cookie.Value

	session, status := sessions[sessionToken]
	if !status {
		http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Expires: time.Now(), Path: "/"})
		return "401 UNAUTHORIZED: INVALID SESSION TOKEN"
	}

	if sessionExpired(session) {
		DeleteSessionAndCookie(w, sessionToken)
		return "401 UNAUTHORIZED: SESSION EXPIRED"
	}

	return session.Nickname
}

func DeleteSessionAndCookie(w http.ResponseWriter, sessionToken string) {
	delete(sessions, sessionToken)
	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Expires: time.Now(), Path: "/"})
}

func UserLoggedIn(nickname string) bool {
	for _, session := range sessions {
		if strings.Compare(session.Nickname, nickname) == 0 {
			fmt.Println("UserLoggedIn: ", session.Nickname, " ", nickname)
			return true
		}
	}
	return false
}

func LogUserOut(conn *websocket.Conn, r *http.Request) string {
	fmt.Println("LogUserOut called")

	if r == nil {
		fmt.Println("DEBUG: r == nil")
	}

	cookie, err := r.Cookie("session_token")
	if err == http.ErrNoCookie {
		return "401 UNAUTHORIZED: CLIENT COOKIE NOT SET OR SESSION EXPIRED"
	}
	if err != nil {
		return "400 BAD REQUEST: REQUEST NOT ALLOWED"
	}
	sessionToken := cookie.Value

	var nickname string
	for _, session := range sessions {
		if strings.Compare(session.Cookie, sessionToken) == 0 {
			nickname = session.Nickname
			fmt.Println("Server >> User " + nickname + " has logged out!")
		}
	}

	delete(LoggedInUsers, nickname)

	var templist []ServerUser
	for _, v := range LoggedInUsers {
		templist = append(templist, ServerUser{Nickname: v.Nickname})
	}

	Broadcast <- ServerMessage{Type: "users", Users: templist}

	delete(sessions, sessionToken)

	db := OpenDatabase()
	defer CloseDatabase(db)
	err = RemoveSessionFromDatabase(db, sessionToken)
	if err != nil {
		fmt.Println("Error deleting session from database:", err)
	}

	conn.WriteJSON(ServerMessage{Type: "logoutResponse", Data: map[string]string{"logout": "true"}})

	return "200 OK"
}

func RefreshSession(w http.ResponseWriter, r *http.Request) {
	mess := AuthenticateUser(w, r)
	if strings.Compare(mess[:4], "400 ") == 0 ||
		strings.Compare(mess[:4], "401 ") == 0 || !UserLoggedIn(mess) {
		return
	}

	newExpiredTime := time.Now().Add(300 * time.Second)

	for sessionToken, session := range sessions {
		if strings.Compare(session.Nickname, mess) == 0 {
			session.ExpiredTime = newExpiredTime
			http.SetCookie(w, &http.Cookie{Name: "session_token", Value: sessionToken, Expires: newExpiredTime, Path: "/"})
			break
		}
	}
}

func GetSessionTokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("session_token")
	if err == http.ErrNoCookie {
		return ""
	}
	if err != nil {
		return ""
	}
	sessionToken := cookie.Value

	return sessionToken
}

func CheckIfSessionTokenIsValid(sessionToken string) bool {
	_, status := sessions[sessionToken]
	return status
}
