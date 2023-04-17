package main
/*type Track struct {

}

func addTrackToPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	songName := vars["songName"]
	artist := vars["artist"]
	trackURL := vars["url"]

	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Fprintf("Error connecting to MongoDB: %v", err)
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
	_, err = collection.InsertOne(context.TODO(), track)
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

	fmt.Fprintf(w, "You've added the song: %s by the artist %s from the playlist %s\n", songName, artist, playlistName)
	fmt.Fprintf(w, "Its URL is %s\n", trackURL)
}

func removeTrackFromPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	songName := vars["songName"]
	artist := vars["artist"]

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
}*/