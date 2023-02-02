package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"net/smtp"
	"bytes"
	"strings"

	"github.com/adassacoimin/CEN3031-Project/studytube/src/server/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

func BugReport(from string, msg []byte, subject string) {
	// send Email
	auth := smtp.PlainAuth("", from, "Password123", "smtp.gmail.com");

	to := []string{"studytubesupport@gmail.com"}

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg);

	if err != nil {
		log.Fatal(err) }
	// POST Email Info
	postBody, _ := json.Marshal(map[string]string{
		"subject": subject,
		"from": from,
		"to": strings.Join(to, " "),
		"msg": string(msg[:]),
	})
	
	responseBody := bytes.NewBuffer(postBody)

	webName := "https://postman-echo.com/post" // temp var for webName -- MUST be changed to REAL website name 

	resp, err := http.Post(webName, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()
}