package main

import "fmt"

func main() {
	dbstruct, err := NewPost() //returns db struct
	var db PostRepository = dbstruct

	if err != nil {
		fmt.Println("Error while connecting to DB:", err)
		return
	}
	if err := db.Create("fida", "It is an love movie"); err != nil {
		fmt.Println("Error creating post:", err)
	} else {
		fmt.Println(" Post created successfully!")
	}

	/*posts := db.Read()
	fmt.Println("All Posts:")
	for _, post := range posts {
		fmt.Printf("ID: %d | Title: %s | Content: %s\n", post.ID, post.Title, post.Content)
	}*/

	/*updateID := uint(2)
	if err := db.Update(updateID, "Mirchi", "A Romantic Action movie"); err != nil {
		fmt.Println("Error updating post:", err)
	} else {
		fmt.Println("Post updated successfully!")
	}*/

	/*filtered, err := db.Filter("Mirchi")
	if err != nil {
		fmt.Println("fileterd error ", err)
	}
	fmt.Println(filtered)*/

	/*deleteID := uint(2)
	if err := db.Delete(deleteID); err != nil {
		fmt.Println("Error deleting post:", err)
	} else {
		fmt.Println("Post deleted successfully!")
	}*/

}
