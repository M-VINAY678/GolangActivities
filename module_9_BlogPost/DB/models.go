package DB

import "gorm.io/gorm"

// Post represents the database model for a blog post
// gorm.Model includes fields ID, CreatedAt, UpdatedAt, DeletedAt
type Post struct {
	gorm.Model
	Title   string
	Content string
}

// dBClient is the concrete implementation of the PostRepository interface
// It holds a pointer to the GORM database connection
type dBClient struct {
	db *gorm.DB
}

// PostRepository defines the behavior for interacting with blog posts
// Any struct that implements these methods can be used as a PostRepository
type PostRepository interface {
	Create(title, content string) error
	Read() []Post
	Update(id uint, title, content string) error
	Delete(id uint) error
	Filter(title string) ([]Post, error)
}
