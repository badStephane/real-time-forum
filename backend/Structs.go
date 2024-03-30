package backend

import (
	"time"
)

// Global maps to store all users, posts, comments, categories, and sessions
var LoggedInUsers = make(map[string]*Session)

// Struct to define a session
type Session struct {
	Nickname      string
	UserID        int
	Cookie        string
	ExpiredTime   time.Time
	WebSocketConn string
}

// Struct to define a user account
type User struct {
	ID        int
	Nickname  string
	Fname     string
	Lname     string
	Age       int
	Gender    string
	Password  string
	Email     string
	CreatedAt string
	Send      chan []byte
}

// Struct to define a post
type Category struct {
	ID           int
	CategoryName string
	Description  string
	CreatedAt    string
}

type Comment struct {
	ID        int
	UserID    int
	NickName  string
	PostID    int
	Nickname  string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Date      string
	IsEdited  bool
}

type Post struct {
	ID           int
	UserID       int
	NickName     string
	Title        string
	Content      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Date         string
	IsEdited     bool
	CategoryName string
}

type Reaction struct {
	ID        int
	UserID    int
	PostID    int
	IsLiked   int
	CreatedAt string
}

type Relation struct {
	ID         int
	CategoryID int
	PostID     int
}

// Struct to define a message
type Message struct {
	ID       int
	Nickname string
	// The nickname of the recipient of the message (if it is a private message)
	RecipientNickname string
	Content           string
	CreatedAt         string    `json:"createdat"`
	From              string    `json:"from"`
	Text              string    `json:"text"`
	ChatHistory       []Message `json:"chathistory"`
	To                string    `json:"to,omitempty"`
	Read              int       `json:"isread"`
}

// Structs to define a hub (for websockets)
type ServerUser struct {
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
}

type ServerMessage struct {
	Type        string            `json:"type"`
	Users       []ServerUser      `json:"users"`
	Categories  []ServerCategory  `json:"categories"`
	Posts       []ServerPost      `json:"posts"`
	User        ServerUser        `json:"user"`
	Post        ServerPost        `json:"post"`
	Category    ServerCategory    `json:"category"`
	Comment     []ServerComment   `json:"comment"`
	Message     string            `json:"message"`
	To          string            `json:"to"`
	From        string            `json:"from"`
	Text        string            `json:"text"`
	Nickname    string            `json:"nickname"`
	ChatHistory []Message         `json:"chathistory"`
	Password    string            `json:"password"`
	Data        map[string]string `json:"data"`
	Start       int               `json:"start"`
	Limit       int               `json:"limit"`
	Offset      int               `json:"offset"`
}

type ServerCategory struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	ID           int    `json:"id"`
	CategoryName string `json:"categoryname"`
	Description  string `json:"description"`
	CreatedAt    string `json:"createdat"`
}
type ServerPost struct {
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CategoryName string    `json:"category"`
	Author       string    `json:"author"`
	Date         string    `json:"date"`
	ID           int       `json:"id"`
	UserID       int       `json:"userid"`
	NickName     string    `json:"nickname"`
	CreatedAt    time.Time `json:"createdat"`
	UpdatedAt    time.Time `json:"updatedat"`
	IsEdited     bool      `json:"isedited"`
}

type ServerComment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userid"`
	NickName  string    `json:"nickname"`
	PostID    int       `json:"postid"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	Date      string    `json:"date"`
	IsEdited  bool      `json:"isedited"`
}
