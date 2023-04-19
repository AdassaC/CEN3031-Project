package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	//"github.com/adassacoimin/CEN3031-Project/studytube/src/server/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"google.golang.org/genproto/googleapis/iam/credentials/v1"
)

type Song struct{
     songName string
     artistName string
     trackURL string
     userID string
    }
type Playlist struct{
	PlaylistName string
	Songs [100]Song
	UserID string
}

func main() {

	fmt.Print("inside of main.go")
	host := "127.0.0.1:4201" 

	// if err := http.ListenAndServe(host, httpHandler()); err != nil {
	// 	fmt.Print("Failed to listen to " + host)
	// 	log.Fatalf("Failed to listen on %s: %v", host, err)
	// } else {
	// 	fmt.Print("Listening to " + host)
	//}
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		fmt.Print("Failed to listen to " + host)
		log.Fatalf("Failed to listen on %s: %v", host, err)
	} else {
		fmt.Print("Listening to " + host)
	}
	router := mux.NewRouter()

	
	// Spotify helper functions in Go 
	router.HandleFunc("/createPlaylist/{playlistName}/{userID}", addPlaylist).Methods("POST")
	router.HandleFunc("/addsong/{playlistName}/title/{songName}/artist/{artist}/trackURL/{url}", addSongToSpotifyPlaylist).Methods("POST")
	router.HandleFunc("/removetrack/{playlistName}/title/{songName}/artist/{artist}", removeSongFromPlaylist).Methods("POST")
	router.HandleFunc("/updatetrack/{playlistName}/title/{songName}/artist/{artist}/newSong/{newSongName}/newArtist/{newArtistName}/newURL/{updatedURL}", updateSongFromPlaylist).Methods("POST")
	router.HandleFunc("/getPlaylist/{playlistName}/{userID}", getPlaylist).Methods("POST")
	
	
	// WARNING: this route must be the last route defined.

	router.PathPrefix("/").Handler(AngularHandler).Methods("GET")


	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"})
    //ttl := handlers.MaxAge(3600)
    origins := handlers.AllowedOrigins([]string{"http://localhost:4200"})
    log.Fatal(http.ListenAndServe(":4201", handlers.CORS(credentials, methods, origins)(router)))

	//DB testing section
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Disconnect(context.TODO())

}

// httpHandler creates the backend HTTP router for queries, types,
// and serving the Angular frontend.
func httpHandler() http.Handler {

	fmt.Print("inside of httpHandler in Go")
	router := mux.NewRouter()

	
	// Spotify helper functions in Go 
	router.HandleFunc("/createPlaylist/{playlistName}/{userID}", addPlaylist).Methods("POST")
	router.HandleFunc("/addsong/{playlistName}/title/{songName}/artist/{artist}/trackURL/{url}", addSongToSpotifyPlaylist).Methods("POST")
	router.HandleFunc("/removetrack/{playlistName}/title/{songName}/artist/{artist}", removeSongFromPlaylist).Methods("POST")
	router.HandleFunc("/updatetrack/{playlistName}/title/{songName}/artist/{artist}/newSong/{newSongName}/newArtist/{newArtistName}/newURL/{updatedURL}", updateSongFromPlaylist).Methods("POST")
	router.HandleFunc("/getPlaylist/{playlistName}/{userID}", getPlaylist).Methods("GET")
	
	
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

func addPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	userID := vars["userID"]

    var songs [100]Song

    playlist := Playlist{
        PlaylistName: playlistName,
        Songs:        songs,
        UserID:       userID,
    }
    //
    //connect to MongoDB on localhost port 27017
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.TODO())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    _, err = collection.InsertOne(context.TODO(), playlist)
    if err != nil {
        fmt.Errorf("Error adding task to task list: %v", err)
    }

}

func addSongToSpotifyPlaylist(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	songName := vars["songName"]
	artist := vars["artist"]

	// need to get songs from playlist 


	//create new Song  struct with string arguments
    // song := Song{
    //  Song:   songName,
    //  Artist: artistName,
    //  url:    trackURL,
    //  UserID: userID,
    // }
    // playlist := Playlist{
    //  PlaylistName: playlistName,
    //  Songs:        songs,
    //  UserID:       userID,
    // }

    //connect to MongoDB on localhost port 27017
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.TODO())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    //find all tasks with the given userID and description
    filter := bson.M{"playlist": playlistName}
    cursor, err := collection.Find(context.Background(), filter)

    if err != nil {
        fmt.Errorf("Cursor found nothing: %v", err)
    }

    // fmt.Printf(cursor.Current.String())

    defer cursor.Close(context.TODO())

    //decode the cursor into a slice of Task objects
    var playlist []Playlist
    err = cursor.All(context.Background(), &playlist)

    theSongs := playlist[0].Songs

    if len(playlist) == 0 {
        fmt.Printf("0 Size")
    } else {
        fmt.Printf("Not of size 0\n,")
    }
    for i := 0; i < len(theSongs); i++ {
        if theSongs[i].artistName != "" {
            theSongs[i+1].artistName = artist
            theSongs[i+1].songName = songName
            break
        }
    }
    //fmt.Printf(theSongs[1].Artist)

    //update the task status in the database
    status := bson.M{"playlist": playlistName}
    update := bson.M{"$set": bson.M{"Song": theSongs}}
    result, err := collection.UpdateOne(context.Background(), status, update)

    if err != nil {
        fmt.Errorf("Error updating task status in database: %v", err)
    }
    if result.ModifiedCount == 0 {
        fmt.Errorf("Task not found in database")
    }

}

func removeSongFromPlaylist(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	songName := vars["songName"]
	artist := vars["artist"]
	//connect to MongoDB on localhost port 27017
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.TODO())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    //find all tasks with the given userID and description
    filter := bson.M{"playlist": playlistName}
    cursor, err := collection.Find(context.Background(), filter)

    if err != nil {
		fmt.Errorf("Cursor found nothing: %v", err)
    }

    defer cursor.Close(context.TODO())

    //decode the cursor into a slice of Task objects
    var playlist []Playlist
    err = cursor.All(context.Background(), &playlist)

    theSongs := playlist[0].Songs

    if len(playlist) == 0 {
        fmt.Printf("0 Size")
    } else {
        fmt.Printf("Not of size 0\n,")
    }
    for i := 0; i < len(theSongs); i++ {
        if theSongs[i].artistName == artist && theSongs[i].songName == songName {
            theSongs[i].artistName = ""
            theSongs[i].songName = ""
            break
        }
    }

    //update the task status in the database
    status := bson.M{"playlist": playlistName}
    update := bson.M{"$set": bson.M{"Song": theSongs}}
    result, err := collection.UpdateOne(context.Background(), status, update)

    //result, err :=  collection.UpdateOne(context.Background(), songs, songs[]+1)
    if err != nil {
        fmt.Errorf("Error updating task status in database: %v", err)
    }
    if result.ModifiedCount == 0 {
        fmt.Errorf("Task not found in database")
    }

}

func updateSongFromPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	songName := vars["songName"]
	artist := vars["artist"]
	newSongName := vars["newSongName"]
	newArtistName := vars["newArtistName"]
	updatedURL := vars["updatedURL"]
    //connect to MongoDB on localhost port 27017
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.TODO())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    //find all tasks with the given userID and description
    filter := bson.M{"playlist": playlistName}
    cursor, err := collection.Find(context.Background(), filter)

    if err != nil {
        fmt.Errorf("Cursor found nothing: %v", err)
    }

    defer cursor.Close(context.TODO())

    //decode the cursor into a slice of Task objects
    var playlist []Playlist
    err = cursor.All(context.Background(), &playlist)

    theSongs := playlist[0].Songs

    if len(playlist) == 0 {
        fmt.Printf("0 Size")
    } else {
        fmt.Printf("Not of size 0\n,")
    }
    for i := 0; i < len(theSongs); i++ {
        if theSongs[i].artistName == artist && theSongs[i].songName == songName {
            theSongs[i].artistName = newArtistName
            theSongs[i].songName = newSongName
            theSongs[i].trackURL = updatedURL
            break
        }
    }

    //update the task status in the database
    status := bson.M{"playlist": playlistName}
    update := bson.M{"$set": bson.M{"Song": theSongs}}
    result, err := collection.UpdateOne(context.Background(), status, update)

    //result, err :=  collection.UpdateOne(context.Background(), songs, songs[]+1)
    if err != nil {
        fmt.Errorf("Error updating task status in database: %v", err)
    }
    if result.ModifiedCount == 0 {
        fmt.Errorf("Task not found in database")
    }

}
func deletePlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	userID := vars["userID"]
	fmt.Print(userID)
    //connect to MongoDB on localhost port 27017
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.Background())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    //delete the task from the databasec
    filter := bson.M{"playlist": playlistName}
    result, err := collection.DeleteOne(context.Background(), filter)
    if err != nil {
        fmt.Errorf("Error deleting playlist from database: %v", err)
    }
    if result.DeletedCount == 0 {
		fmt.Errorf("Playlist not found in database")
    }

}

func getPlaylist(w http.ResponseWriter, r *http.Request) {

	
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	userID := vars["userID"]
	fmt.Print(userID)

    //connect to MongoDB on localhost port 27017
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.Background())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    //find all tasks with the given userID and description
    filter := bson.M{"playlist": playlistName}
    cursor, err := collection.Find(context.Background(), filter)

    if err != nil {
        fmt.Errorf("Cursor found nothing: %v", err)
    }
    
    defer cursor.Close(context.TODO())

    //decode the cursor into a slice of Task objects
    var playlist []Playlist
    err = cursor.All(context.Background(), &playlist)

    err = json.NewEncoder(w).Encode(playlist)
	if err != nil {
		
	}
    
    

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
