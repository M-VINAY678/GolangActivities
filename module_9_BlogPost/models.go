package main

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string
	Content string
}
type dBClient struct {
	db *gorm.DB
}
type PostRepository interface {
	Create(title, content string) error
	Read() []Post
	Update(id uint, title, content string) error
	Delete(id uint) error
	Filter(title string) ([]Post, error)
}
