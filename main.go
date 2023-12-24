package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handlePostData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var receivedData map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&receivedData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("Received data: %#v\n", receivedData)
		response := map[string]string{"message": "Data received successfully"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/postdata", handlePostData)
	http.HandleFunc("/", serveHTML)

	fmt.Println("Server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
