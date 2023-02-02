package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adassacoimin/CEN3031-Project/studytube/src/server/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/smtp"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello-world", HelloWorld)

	// Solves Cross Origin Access Issue
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})
	handler := c.Handler(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + os.Getenv("PORT"),
	}

	log.Fatal(srv.ListenAndServe())
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Golang + Angular Starter Kit",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return
}

func BugReport(from string, msg []byte) {

	auth := smtp.PlainAuth("", from, "Password123", "smtp.gmail.com");

	to := []string{"studytubesupport@gmail.com"}

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg);

	if err != nil {
		log.Fatal(err) }
}