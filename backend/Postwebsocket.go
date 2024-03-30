package backend

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Send the updated list of categories to all clients
func handleGetCategoriesMessage(conn *websocket.Conn, message ServerMessage) {

	// open the database
	db := OpenDatabase()
	defer CloseDatabase(db)

	// Get all the categories from the database and send them to the client
	tempUserlist, _ := GetAllCategories(db)
	tempUserlist2 := make([]ServerCategory, len(tempUserlist))
	for i, v := range tempUserlist {
		tempUserlist2[i] = ServerCategory{ID: v.ID, CategoryName: v.CategoryName, Description: v.Description, CreatedAt: v.CreatedAt}
	}

	Broadcast <- ServerMessage{Type: "categories", Categories: tempUserlist2}
}

// function to handle the new post message (when a new post is created)
func handleNewPostMessage(message ServerMessage) {
	// open the database
	db := OpenDatabase()
	defer CloseDatabase(db)

	// Add the new post to the database
	AddPost(db, message.Data["title"], message.Data["content"], message.Data["category"], message.Data["nickname"])

	//Add the post category relation to the database
	AddPostCategoryRelation(db, message.Data["title"], message.Data["category"])

	// Get all the posts from the database and send them to the client
	posts, _ := GetLatestPosts(db)
	postList := make([]ServerPost, len(posts))
	for i, v := range posts {
		postList[i] = ServerPost{ID: v.ID, Title: v.Title, Content: v.Content, CreatedAt: v.CreatedAt, UserID: v.UserID, NickName: v.NickName, CategoryName: v.CategoryName}
	}

	// send the posts to the client
	Broadcast <- ServerMessage{Type: "posts", Posts: postList}
}

func handleGetPostsMessage(conn *websocket.Conn, message ServerMessage) {
	//open db
	db := OpenDatabase()
	defer db.Close()
	// Get all the posts from the database and send them to the client
	posts, _ := GetLatestPosts(db)
	postList := make([]ServerPost, len(posts))
	for i, v := range posts {
		postList[i] = ServerPost{ID: v.ID, Title: v.Title, Content: v.Content, CreatedAt: v.CreatedAt, UserID: v.UserID, NickName: v.NickName, CategoryName: v.CategoryName} // Add the CategoryName field
	}

	// send the posts to the client
	Broadcast <- ServerMessage{Type: "posts", Posts: postList}
}

func handleGetCommentsMessage(conn *websocket.Conn, message ServerMessage) {
	//open db
	db := OpenDatabase()
	defer db.Close()

	comments, err := GetCommentsByPostTitle(db, message.Data["content"])
	if err != nil {
		fmt.Println("Error getting comments")
	}

	commentList := make([]ServerComment, len(comments))
	for i, v := range comments {
		commentList[i] = ServerComment{ID: v.ID, Content: v.Content, CreatedAt: v.CreatedAt, UserID: v.UserID, NickName: v.Nickname}
	}

	// send the comments to the requesting client
	response := ServerMessage{Type: "comments", Comment: commentList}
	conn.WriteJSON(response)
}

func handleNewCommentMessage(conn *websocket.Conn, message ServerMessage) {
	//open db
	db := OpenDatabase()
	defer db.Close()
	// Get the user that sent the message
	user := message.Data["nickname"]

	// Insert the comment into the database
	InsertComment(db, message.Data["content"], user, message.Data["postid"])
	// Get all the comments from the database and send them to the client
	comments, _ := GetCommentsByPostTitle(db, message.Data["postid"])
	commentList := make([]ServerComment, len(comments))
	for i, v := range comments {
		commentList[i] = ServerComment{ID: v.ID, Content: v.Content, CreatedAt: v.CreatedAt, PostID: v.PostID, UserID: v.UserID, NickName: v.NickName}
	}
	// send the comments to the client
	//Broadcast <- ServerMessage{Type: "comments", Comments: commentList}
}


// get posts for categories
func handleGetPostsForCategory(conn *websocket.Conn, message ServerMessage) {
	var category = message.Data["Text"]
	//open db
	db := OpenDatabase()
	defer db.Close()
	// Get all the posts from the database and send them to the client
	posts, _ := GetPostsByCategory(db, category)
	postList := make([]ServerPost, len(posts))
	for i, v := range posts {
		postList[i] = ServerPost{ID: v.ID, Title: v.Title, Content: v.Content, CreatedAt: v.CreatedAt, UserID: v.UserID, NickName: v.NickName, CategoryName: v.CategoryName} // Add the CategoryName field
	}

	// send the posts to the client
	conn.WriteJSON(ServerMessage{Type: "postsbyCategory", Posts: postList})
}
