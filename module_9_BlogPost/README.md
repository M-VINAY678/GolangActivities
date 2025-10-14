# Go Gorm CRUD Project : Blog Post
This project demonstrates a simple CRUD(Create,Read,Update,Delete) applications uisng Go and Gorm With MySQL.
It allows managing Post records in a database through a clean interface with command-line database configuration.

##Features
* Connect to MySQL database using credential from .env or command-line flag.
* CRUD operations on Post Records:
 - Create a new post
 - Read all posts
 - Update a posts
 - Delete a post
 - Filter posts by title
- Uses GORM for database interactions and auto-migration.
- Interface-driven design with PostRepository for abstraction

## To run it in Locally 
go run . --dbname=blogPostDB

