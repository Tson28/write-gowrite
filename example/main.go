#author: Nguyễn Thái Sơn
package main

import (
	"fmt"
	"log"

	"github.com/writeas/go-writeas"
)

func main() {
	fmt.Println("=== Write.as Go Client Library Example ===")
	
	// Create a new client
	c := writeas.NewClient()
	
	// Example 1: Get a public post
	fmt.Println("\n1. Getting a public post...")
	post, err := c.GetPost("3psnxyhqxy3hq")
	if err != nil {
		log.Printf("Error getting post: %v", err)
	} else {
		fmt.Printf("Post Title: %s\n", post.Title)
		fmt.Printf("Post Content: %s\n", post.Content[:100] + "...")
	}
	
	// Example 2: Get a collection
	fmt.Println("\n2. Getting a collection...")
	collection, err := c.GetCollection("blog")
	if err != nil {
		log.Printf("Error getting collection: %v", err)
	} else {
		fmt.Printf("Collection Title: %s\n", collection.Title)
		fmt.Printf("Collection Description: %s\n", collection.Description)
	}
	
	// Example 3: Create a new post (anonymous)
	fmt.Println("\n3. Creating a new post...")
	newPost, err := c.CreatePost(&writeas.PostParams{
		Title:   "Test Post from Go Client",
		Content: "This is a test post created using the Go client library.",
		Font:    "sans",
	})
	if err != nil {
		log.Printf("Error creating post: %v", err)
	} else {
		fmt.Printf("New post created with ID: %s\n", newPost.ID)
		fmt.Printf("Post URL: %s\n", newPost.URL)
		
		// Save the token for later use
		token := newPost.Token
		fmt.Printf("Post token: %s\n", token)
		
		// Example 4: Update the post
		fmt.Println("\n4. Updating the post...")
		updatedPost, err := c.UpdatePost(&writeas.PostParams{
			ID:      newPost.ID,
			Token:   token,
			Content: "This post has been updated using the Go client library!",
		})
		if err != nil {
			log.Printf("Error updating post: %v", err)
		} else {
			fmt.Printf("Post updated successfully: %s\n", updatedPost.Content)
		}
		
		// Example 5: Delete the post
		fmt.Println("\n5. Deleting the post...")
		err = c.DeletePost(&writeas.PostParams{
			ID:    newPost.ID,
			Token: token,
		})
		if err != nil {
			log.Printf("Error deleting post: %v", err)
		} else {
			fmt.Println("Post deleted successfully!")
		}
	}
	
	fmt.Println("\n=== Example completed ===")
}
