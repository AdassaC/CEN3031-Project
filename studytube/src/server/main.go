package main

import (
	//"encoding/json"
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	//"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	//"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


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

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	// For sample support and debugging, not required for production:
	stripe.SetAppInfo(&stripe.AppInfo{
		Name:    "stripe-samples/subscription-use-cases/usage-based-subscriptions",
		Version: "0.0.1",
		URL:     "https://github.com/stripe-samples/subscription-use-cases/usage-based-subscriptions",
	})
 

	// Cookie API Handlers
	router.HandleFunc("/userID/{userID}/set", setUserIDCookieHandler).Methods("POST")
	router.HandleFunc("/userID/get", getUserIDCookieHandler).Methods("GET")
	
	// Spotify API Handlers 
	router.HandleFunc("/addplaylist/{playlistName}/title/{songName}/artist/{artist}/trackURL/{url}", addTrackToPlaylist).Methods("POST")
	router.HandleFunc("/removetrack/{playlistName}/title/{songName}/artist/{artist}", removeTrackFromPlaylist).Methods("DELETE") // not sure about DELETE
	router.HandleFunc("/updatetrack/{playlistName}/title/{songName}/artist/{artist}/newSong/{newSongName}/newArtist/{newArtistName}/newURL/{updatedURL}", updateTrackOnPlaylist).Methods("PUT")
	router.HandleFunc("/createPlaylist/{playlistName}", createPlaylist).Methods("POST")
	router.HandleFunc("/getPlaylist{playlistName}", getPlaylist).Methods("GET")
	
	// Stripe API Handlers
	http.HandleFunc("/config", handleConfig)
	http.HandleFunc("/create-customer", handleCreateCustomer)
	http.HandleFunc("/retrieve-customer-payment-method", handleRetrieveCustomerPaymentMethod)
	http.HandleFunc("/create-subscription", handleCreateSubscription)
	http.HandleFunc("/cancel-subscription", handleCancelSubscription)
	http.HandleFunc("/update-subscription", handleUpdateSubscription)
	http.HandleFunc("/retry-invoice", handleRetryInvoice)
	http.HandleFunc("/retrieve-upcoming-invoice", handleRetrieveUpcomingInvoice)
	http.HandleFunc("/webhook", handleWebhook)
	
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
			handlers.ExposedHeaders([]string{"DNT", "Kxeep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			handlers.MaxAge(86400),
		)(router))
}

type Track struct {

}

func addTrackToPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	songName := vars["songName"]
	artist := vars["artist"]
	trackURL := vars["url"]



	/*
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("SpotifyStructure")

	track := Track{
		SongTitle: songName,
		ArtistName: artist,
		UrlName: trackURL,
	}

				////add the task to the database
				//err = addTrack(track.SongTitle, track.ArtistName, track.UrlName)
				//if err != nil {
				//	fmt.Fprintf("Error adding task to task list: %v", err)
				//}
	//add the task to the task list
	_, err = collection.InsertOne(context.TODO(), task)
		if err != nil {
			return fmt.Errorf("Error adding task to task list: %v", err)
		}


	//check that the task was added to the database
	filter := bson.M{"description": "Test Task"}
	var result Track
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		fmt.Fprintf("Error finding task in database: %v", err)
	}
	//check
	if result.SongTitle != track.SongTitle || result.ArtistName != track.ArtistName || result.UrlName != track.UrlName {
		fmt.Print("Task not added correctly to database")
	}

	*/

	fmt.Fprintf(w, "You've added the song: %s by the artist %s from the playlist %s\n", songName, artist, playlistName)
	fmt.Fprintf(w, "Its URL is %s\n", trackURL)
}

func removeTrackFromPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	songName := vars["songName"]
	artist := vars["artist"]

	/*

	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return fmt.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//delete the task from the database
	status := bson.M{"description": description}
	result, err := collection.DeleteOne(context.Background(), status)
	if err != nil {
		return fmt.Errorf("Error deleting task from database: %v", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("Task not found in database")
	}

	*/
	

	fmt.Fprintf(w, "You've removed the song: %s by the artist %s from the playlist %s\n", songName, artist, playlistName)
}

func updateTrackOnPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	songName := vars["songName"]
	artist := vars["artist"]
	newSongName := vars["newSongName"]
	newArtistName := vars["newArtistName"]
	updatedURL := vars["updatedURL"]

	/*
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return fmt.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//update the task status in the database
	newSong := bson.M{"songName": newSongName}
	newArtist := bson.M{"artist": newArtistName}	
	// may still have to update the url

	result, err := collection.UpdateOne(context.Background(), newSong, newArtist)
	if err != nil {
		return fmt.Errorf("Error updating task status in database: %v", err)
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("Task not found in database")
	}

	*/
	fmt.Fprintf(w, "You've updated the song: %s by the artist %s from the playlist %s\n", songName, artist, playlistName)
	fmt.Fprintf(w, "You've changed it to song: %s by the artist %s with the url %s\n", newSongName, newArtistName, updatedURL)
}

func createPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]

	fmt.Fprintf(w, "You've created the playlist: %s\n", playlistName)
}

func getPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]

	fmt.Fprintf(w, "You've retrieved the playlist: %s\n", playlistName)
}

type Book struct {
	Name string `json:"name"`
}

func handleBookGet(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	cookie := http.Cookie{
		Name: "userID",
		Value: title, // need to put value here 
		Path: "/",
		HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteLaxMode,
	}
	// Use the http.SetCookie() function to send the cookie to the client.
    // Behind the scenes this adds a `Set-Cookie` header to the response
    // containing the necessary cookie data.
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "You've attempted to create the cookie: %s\n", title)

	io.WriteString(w, cookie.String())


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

func setUserIDCookieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := vars["userID"]

	cookie := http.Cookie{
		Name: "userID",
		Value: value, // need to put value here 
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