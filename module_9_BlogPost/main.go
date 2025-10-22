package main

import (
	"fmt"
	"module_9_BlogPost/DB"
)

func main() {
	// Initialize the database connection and get a PostRepository interface
	// NewPost() connects to the DB, performs migration, and returns an implementation of PostRepository
	db, err := DB.NewPost()
	if err != nil {
		fmt.Println("Error while connecting to DB:", err)
		return
	}

	// Create a new post in the database
	if err := db.Create("fida", "It is an love movie"); err != nil {
		fmt.Println("Error creating post:", err)
	} else {
		fmt.Println(" Post created successfully!")
	}

	// Read all posts from the database
	/*posts := db.Read()
	fmt.Println("All Posts:")
	for _, post := range posts {
		fmt.Printf("ID: %d | Title: %s | Content: %s\n", post.ID, post.Title, post.Content)
	}*/

	// Update an existing post by ID
	/*updateID := uint(2)
	if err := db.Update(updateID, "Mirchi", "A Romantic Action movie"); err != nil {
		fmt.Println("Error updating post:", err)
	} else {
		fmt.Println("Post updated successfully!")
	}*/

	// Filter posts by title
	/*filtered, err := db.Filter("Mirchi")
	if err != nil {
		fmt.Println("fileterd error ", err)
	}
	fmt.Println(filtered)*/

	// Delete a post by ID
	/*deleteID := uint(2)
	if err := db.Delete(deleteID); err != nil {
		fmt.Println("Error deleting post:", err)
	} else {
		fmt.Println("Post deleted successfully!")
	}*/

}
