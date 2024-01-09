# Managing users in a database using Golang and PostgreSQL

## Project description
The project is an example application in the Golang programming language for managing users in a PostgreSQL database. It includes create, read, update, and delete (CRUD) operations, as well as a simple web interface for interacting with the application.

## Customer oriented
The project is aimed at developers who want to learn the basics of working with a database in the Golang programming language, as well as those who are looking for an example of using Gorm (Object-Relational Mapping) to simplify working with a PostgreSQL database.

## Participants
- Zeytkazy Sayat
- Aldebaev Beksultan
- Zhanzhigitov Madiyar

## Screenshot of the first page
![Screenshot](https://github.com/Safyd-Zey/AP1/blob/main/image.png)

## Startup instructions
1. Install PostgreSQL and create a database and user for the project.
2. Run the database initialization script: `go run init_db.go`.
3. Install golang-migrate: `go get -tags 'postgres' -u github.com/golang-migrate/migrate/`.
4. Create migrations to create the `users` table and add the `age` field.
    - Creating a new migration: `migrate create -ext sql -dir db/migrations -seq name_of_migration`.
    - Applying migrations: `migrate -path ./migrations -database "postgres://postgres:123456789@localhost:5432/newDB?sslmode=disable" up`.
5. Install Gorm: `go get -u gorm.io/gorm`.
6. Start the server: `go run main.go`.
7. Open a web browser and go to [http://localhost:8080](http://localhost:8080).

## Tools used
- Golang
- PostgreSQL
- Gorm
- golang-migrate
- HTML/CSS

## Sources
- [Official Golang Documentation](https://golang.org/doc/)
- [Gorm Documentation](https://gorm.io/docs/)
- [GitHub repository golang-migrate](https://github.com/golang-migrate/migrate/)
- [Scoop - command-line installer for Windows](https://scoop.sh/)
