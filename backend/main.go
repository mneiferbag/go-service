package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Marshaling (Go struct to JSON)
	user := User{Name: "Alice", Email: "alice@example.com"}
	jsonData, _ := json.Marshal(user)

	// Unmarshaling (JSON to Go struct)
	var newUser User
	json.Unmarshal(jsonData, &newUser)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(
		w, "%s", jsonData,
	)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handler)

	fmt.Println("Go backend started!")
	log.Fatal(http.ListenAndServe(":8080", r))
}
