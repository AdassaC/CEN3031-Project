package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adassacoimin/CEN3031-Project/studytube/src/server/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Song struct{
     songName string
     artistName string
     trackURL string
     userID string
    }
type Playlist struct{
	playlistName string
	songs []Song
	userID string
}

func main() {
	r := mux.NewRouter()


	// Spotify helper functions in Go 
	r.HandleFunc("/addsong/{playlistName}/title/{songName}/artist/{artist}/trackURL/{url}", addSongToSpotifyPlaylist)
	r.HandleFunc("/removetrack/{playlistName}/title/{songName}/artist/{artist}", removeSongFromPlaylist)
	r.HandleFunc("/updatetrack/{playlistName}/title/{songName}/artist/{artist}/newSong/{newSongName}/newArtist/{newArtistName}/newURL/{updatedURL}", updateTrackOnPlaylist)
	r.HandleFunc("/createPlaylist/{playlistName}", addPlaylist)
	r.HandleFunc("/getPlaylist/{playlistName}/{userID}", getPlaylist)

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

func addPlaylist(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	
    playlist := Playlist{
        PlaylistName: playlistName,
        Songs:        songs,
        UserID:       userID,
    }
    //
    //connect to MongoDB on localhost port 27017
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        return fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.TODO())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    _, err = collection.InsertOne(context.TODO(), playlist)
    if err != nil {
        return fmt.Errorf("Error adding task to task list: %v", err)
    }

    return nil
}

func addSongToSpotifyPlaylist(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	songName := vars["songName"]
	artist := vars["artist"]
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
        return fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.TODO())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    //find all tasks with the given userID and description
    filter := bson.M{"playlist": playlistName}
    cursor, err := collection.Find(context.Background(), filter)

    if err != nil {
        return fmt.Errorf("Cursor found nothing: %v", err)
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
        if theSongs[i].Artist != "" {
            theSongs[i+1].Artist = artist
            theSongs[i+1].Song = songName
            break
        }
    }
    //fmt.Printf(theSongs[1].Artist)

    //update the task status in the database
    status := bson.M{"playlist": playlistName}
    update := bson.M{"$set": bson.M{"Song": theSongs}}
    result, err := collection.UpdateOne(context.Background(), status, update)

    if err != nil {
        return fmt.Errorf("Error updating task status in database: %v", err)
    }
    if result.ModifiedCount == 0 {
        return fmt.Errorf("Task not found in database")
    }

    return nil
}

func removeSongFromPlaylist(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	songName := vars["songName"]
	artist := vars["artist"]
	//connect to MongoDB on localhost port 27017
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        return fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.TODO())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    //find all tasks with the given userID and description
    filter := bson.M{"playlist": playlistName}
    cursor, err := collection.Find(context.Background(), filter)

    if err != nil {
        return fmt.Errorf("Cursor found nothing: %v", err)
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
        if theSongs[i].Artist == artist && theSongs[i].Song == songName {
            theSongs[i].Artist = ""
            theSongs[i].Song = ""
            break
        }
    }

    //update the task status in the database
    status := bson.M{"playlist": playlistName}
    update := bson.M{"$set": bson.M{"Song": theSongs}}
    result, err := collection.UpdateOne(context.Background(), status, update)

    //result, err :=  collection.UpdateOne(context.Background(), songs, songs[]+1)
    if err != nil {
        return fmt.Errorf("Error updating task status in database: %v", err)
    }
    if result.ModifiedCount == 0 {
        return fmt.Errorf("Task not found in database")
    }

    return nil
}

func updateSongFromPlaylist(w http.ResponseWriter, r *http.Request) error {
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
        return fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.TODO())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    //find all tasks with the given userID and description
    filter := bson.M{"playlist": playlistName}
    cursor, err := collection.Find(context.Background(), filter)

    if err != nil {
        return fmt.Errorf("Cursor found nothing: %v", err)
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
        if theSongs[i].Artist == artist && theSongs[i].Song == songName {
            theSongs[i].Artist = newArtistName
            theSongs[i].Song = newSongName
            theSongs[i].url = updatedURL
            break
        }
    }

    //update the task status in the database
    status := bson.M{"playlist": playlistName}
    update := bson.M{"$set": bson.M{"Song": theSongs}}
    result, err := collection.UpdateOne(context.Background(), status, update)

    //result, err :=  collection.UpdateOne(context.Background(), songs, songs[]+1)
    if err != nil {
        return fmt.Errorf("Error updating task status in database: %v", err)
    }
    if result.ModifiedCount == 0 {
        return fmt.Errorf("Task not found in database")
    }

    return nil
}
func deletePlaylist(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	userID := vars["userID"]
	fmt.Print(userID)
    //connect to MongoDB on localhost port 27017
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        return fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.Background())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    //delete the task from the databasec
    filter := bson.M{"playlist": playlistName}
    result, err := collection.DeleteOne(context.Background(), filter)
    if err != nil {
        return fmt.Errorf("Error deleting playlist from database: %v", err)
    }
    if result.DeletedCount == 0 {
        return fmt.Errorf("Playlist not found in database")
    }

    return nil
}

func getPlaylist(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	userID := vars["userID"]
	fmt.Print(userID)

    //connect to MongoDB on localhost port 27017
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        return fmt.Errorf("Error connecting to MongoDB: %v", err)
    }
    defer client.Disconnect(context.Background())

    //get a handle to the tasks collection
    collection := client.Database("CEN3031_Test").Collection("SongStructure")

    //find all tasks with the given userID and description
    filter := bson.M{"playlist": playlistName}
    cursor, err := collection.Find(context.Background(), filter)

    if err != nil {
        return fmt.Errorf("Cursor found nothing: %v", err)
    }
    
    defer cursor.Close(context.TODO())

    //decode the cursor into a slice of Task objects
    var playlist []Playlist
    err = cursor.All(context.Background(), &playlist)

    //put task into JSON for transfer
    jsonBytes, err := utils.StructToJSON(playlist)
    if err != nil {
    }

    //set content type to json
    fmt.Println(jsonBytes)

    return nil
}