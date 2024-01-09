package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456789"
	dbname   = "newDB"
)

// User модель для работы с таблицей users
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

var db *gorm.DB

func initDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = database

	// Auto Migrate - автоматическое создание таблицы users
	db.AutoMigrate(&User{})
}

func createUser(name, email, password string) {
	user := User{Name: name, Email: email, Password: password}
	db.Create(&user)
}

func getUserByID(id uint) (User, error) {
	var user User
	result := db.First(&user, id)
	return user, result.Error
}

func updateUserNameByID(id uint, newName string) error {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.Model(&user).Update("Name", newName)
	return result.Error
}

func deleteUserByID(id uint) error {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&user, id)
	return result.Error
}

func getAllUsers() ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

func main() {
	initDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing the form", http.StatusBadRequest)
				return
			}

			name := r.FormValue("name")
			email := r.FormValue("email")
			password := r.FormValue("password")

			fmt.Printf("Получены данные:\nName: %s\nEmail: %s\nPassword: %s\n", name, email, password)

			createUser(name, email, password)

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "The data has been successfully inserted into the database!")
			return
		} else if r.Method == http.MethodGet {
			users, err := getAllUsers()
			if err != nil {
				http.Error(w, "Error getting users from database", http.StatusInternalServerError)
				return
			}

			// Возвращаем список пользователей в виде HTML-страницы или JSON, в зависимости от вашего предпочтения
			// В данном случае просто выводим данные в консоль
			for _, user := range users {
				fmt.Printf("ID: %d, Name: %s, Email: %s, Password: %s\n", user.ID, user.Name, user.Email, user.Password)
			}
		}
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", nil)
}
