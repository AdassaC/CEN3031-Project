package main

import (
	//"encoding/json"
	"fmt"
	"log"
	//"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	//"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
/*
type Post struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
  }
  var posts []Post
  func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
  }
  func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
  }
  func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
	  if item.ID == params["id"] {
		json.NewEncoder(w).Encode(item)
		return
	  }
	}
	json.NewEncoder(w).Encode(&Post{})
  }
  func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
	  if item.ID == params["id"] {
		posts = append(posts[:index], posts[index+1:]...)
		var post Post
		_ = json.NewDecoder(r.Body).Decode(&post)
		post.ID = params["id"]
		posts = append(posts, post)
		json.NewEncoder(w).Encode(&post)
		return
	  }
	}
	json.NewEncoder(w).Encode(posts)
  }
  func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
	  if item.ID == params["id"] {
		posts = append(posts[:index], posts[index+1:]...)
		break
	  }
	}
	json.NewEncoder(w).Encode(posts)
  }

  func handleBookPost(w http.ResponseWriter, r *http.Request) {
	fmt.Print("You are inside of the handleBookPost method");

	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've attempted to post the book: %s on page %s\n", title, page)
   } 

  func main() {
	fmt.Print("inside of main")
	router := mux.NewRouter()
	posts = append(posts, Post{ID: "1", Title: "My first post", Body:      "This is the content of my first post"})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	//router.HandleFunc("/booksPost/{title}/page/{page}", handleBookPost).Methods("POST");
	router.HandleFunc("/books", handleBookPost).Methods("POST");
	// Spotify Handler 
	//router.HandleFunc("/songs", getSongs).Methods("GET")
	//router.HandleFunc("/song/{song}", getSong).Methods("GET")
	//router.HandleFunc("/playlists", getPlaylists).Methods("GET")
	//router.HandleFunc("/playlist{playlist}", getPlaylist).Methods("GET")

	//router.HandleFunc("/playlist{playlist}", createPlaylist).Methods("POST")
	// need more parameters and put info onto google doc sheet 


  http.ListenAndServe(":4201", router)
  }
  */



func main() {

	fmt.Print("inside of main.go")
	host := "127.0.0.1:4201" // may be 4201
	//host := "http://localhost:4201"
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		fmt.Print("Failed to listen to " + host)
		log.Fatalf("Failed to listen on %s: %v", host, err)
	} else {
		fmt.Print("Listening to " + host)
	}

}

// httpHandler creates the backend HTTP router for queries, types,
// and serving the Angular frontend.
func httpHandler() http.Handler {

	fmt.Print("inside of httpHandler in Go")
	router := mux.NewRouter()
	// Your REST API requests go here

	router.HandleFunc("/books/{title}/page/{page}", handleBookGet).Methods("GET")
	router.HandleFunc("/books", handleBooksGet).Methods("GET");
		//router.HandleFunc("/booksPost/{title}/page/{page}", handleBookPost).Methods("POST");
	router.HandleFunc("/books", handleBookPost).Methods("POST");
		//router.HandleFunc("/test", test)
	
	// Add your routes here.
	// WARNING: this route must be the last route defined.

	router.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	
	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
				"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
				"Cache-Control", "Content-Range", "Range"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:4200"}), // maybe should be 4020???
			handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			handlers.MaxAge(86400),
		)(router))
}

type Book struct {
	Name string `json:"name"`
}

func handleBookGet(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You are insideweew of the handleBookGet");
	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func handleBooksGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WE are inside of the handleBooksGet");
}

func handleBookPost(w http.ResponseWriter, r *http.Request) {
	fmt.Print("You are inside of the handleBookPost method");

	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've attempted to post the book: %s on page %s\n", title, page)
} 



func getOrigin() *url.URL {
	origin, _ := url.Parse("http://localhost:4200")
	return origin
}

var origin = getOrigin()

var director = func(req *http.Request) {
	req.Header.Add("X-Forwarded-Host", req.Host)
	req.Header.Add("X-Origin-Host", origin.Host)
	req.URL.Scheme = "http"
	req.URL.Host = origin.Host
}

// AngularHandler loads angular assets
var AngularHandler = &httputil.ReverseProxy{Director: director}