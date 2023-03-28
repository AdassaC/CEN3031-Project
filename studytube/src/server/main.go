package main

import (
	"fmt"
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

	// Spotify helper functions in Go 
	r.HandleFunc("/addplaylist/{playlistName}/title/{songName}/artist/{artist}/trackURL/{url}", addTrackToPlaylist)
	r.HandleFunc("/removetrack/{playlistName}/title/{songName}/artist/{artist}", removeTrackFromPlaylist)
	r.HandleFunc("/updatetrack/{playlistName}/title/{songName}/artist/{artist}/newSong/{newSongName}/newArtist/{newArtistName}/newURL/{updatedURL}", updateTrackOnPlaylist)
	r.HandleFunc("/createPlaylist/{playlistName}", createPlaylist)
	r.HandleFunc("/getPlaylist/{playlistName}", getPlaylist)

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

func addTrackToPlaylist(w http.ResponseWriter, r *http.Request) {

}

func removeTrackFromPlaylist(w http.ResponseWriter, r *http.Request) {

}

func updateTrackOnPlaylist(w http.ResponseWriter, r *http.Request) {

}

func createPlaylist(w http.ResponseWriter, r *http.Request) {

}

func getPlaylist(w http.ResponseWriter, r *http.Request) {

}