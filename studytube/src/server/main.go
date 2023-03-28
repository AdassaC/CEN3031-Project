package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/adassacoimin/CEN3031-Project/studytube/src/server/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello-world", helloWorld)

	r.HandleFunc("/userID/set", setUserIDCookieHandler)
	r.HandleFunc("/userID/get", getUserIDCookieHandler)
	// eventually get parameters for userID 
	// "/userID/{userIDVal}/get"

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

func helloWorld(w http.ResponseWriter, r *http.Request) {
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

func setUserIDCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name: "userID",
		Value: "101",
		Path: "/",
		HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteLaxMode,
	}
	// Use the http.SetCookie() function to send the cookie to the client.
    // Behind the scenes this adds a `Set-Cookie` header to the response
    // containing the necessary cookie data.
	http.SetCookie(w, &cookie)

	io.WriteString(w, cookie.String())
}

func getUserIDCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("userID")
	if err != nil {
		switch {
        case errors.Is(err, http.ErrNoCookie):
            http.Error(w, "Cookie not found", http.StatusBadRequest)
        default:
            log.Println(err)
            http.Error(w, "Server error", http.StatusInternalServerError)
        }
        return
	}
	io.WriteString(w, cookie.String())
}