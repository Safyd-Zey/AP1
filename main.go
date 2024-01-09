package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456789"
	dbname   = "newDB"
)

func main() {
	// Открываем соединение с базой данных
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Проверяем соединение с базой данных
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing the form", http.StatusBadRequest)
				return
			}

			name := r.FormValue("name")
			email := r.FormValue("email")
			password := r.FormValue("password")

			fmt.Printf("Получены данные:\nName: %s\nEmail: %s\nPassword: %s\n", name, email, password)

			// Вставляем данные в базу данных
			_, err = db.Exec("INSERT INTO users(name, email, password) VALUES ($1, $2, $3)", name, email, password)
			if err != nil {
				http.Error(w, "Error when inserting data into database", http.StatusInternalServerError)
				return
			}

			// Ответ клиенту
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "The data has been successfully inserted into the database!")
			return
		}

		// Возвращаем HTML-страницу с формой для GET-запроса
		http.ServeFile(w, r, "index.html")

	})

	fmt.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", nil)
}

