package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewPost acts like a constructor for dBClient.
// It loads environment variables, reads the database name from CLI flags or .env,
// establishes a GORM database connection, and performs auto-migration for the Post model.
func NewPost() (*dBClient, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbNameFlag := flag.String("dbname", "", "Database name to connect to") //returns a pointer to a string that will hold value provided by CLI
	flag.Parse()
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := *dbNameFlag
	if dbName == "" {
		dbName = os.Getenv("DB_NAME")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database : %v", err)
		return nil, err
	}
	fmt.Println("Database Succesfully connected")
	if err := db.AutoMigrate(&Post{}); err != nil {
		fmt.Println("Migration Failed", err)
		return nil, err
	}
	return &dBClient{db: db}, nil
}

// Create inserts a new Post record into the database and returns error if any occurs during insertion
func (dBClient *dBClient) Create(title, content string) error {
	post := Post{Title: title, Content: content}
	if err := dBClient.db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

// Read retrieves all Post records from the database.
func (dBClient *dBClient) Read() []Post {
	var posts []Post
	dBClient.db.Find(&posts)
	return posts
}

// Update modifies an existing Post record identified by its ID.
func (dBClient *dBClient) Update(id uint, title, content string) error {
	var post Post
	if err := dBClient.db.First(&post, id).Error; err != nil {
		return err
	}
	post.Title = title
	post.Content = content
	return dBClient.db.Save(&post).Error
}

// Delete removes a Post record from the database by ID.
func (dBClient *dBClient) Delete(id uint) error {
	return dBClient.db.Delete(&Post{}, id).Error
}

// Filter retrieves Post records that match the given title.
func (dBClient *dBClient) Filter(title string) ([]Post, error) {
	var posts []Post
	if err := dBClient.db.Where("title=?", title).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
