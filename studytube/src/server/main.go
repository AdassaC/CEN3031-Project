package main

import (
	"context"
	"encoding/json"
	"fmt"

	// "log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/adassacoimin/CEN3031-Project/studytube/src/server/utils"
	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// task struct
type Task struct {
	ID          string `bson:"_id,omitempty"`
	Description string `bson:"description"`
	Status      string `bson:"status"`
	UserID      string `bson:"userID"`
	ListName    string `bson:"listName"`
}

type TaskList struct {
	ListName string `bson:"listName"`
	UserID   string `bson:"userID"`
}

type Playlist struct {
	ID           string    `bson:"_id,omitempty"`
	PlaylistName string    `bson:"playlist"`
	Songs        [100]Song `bson:"Song"`
	UserID       string    `bson:"userID"`
}

type Song struct {
	ID     string `bson:"_id,omitempty"`
	Song   string `bson:"song"`
	Artist string `bson:"artist"`
	url    string `bson:"url"`
	// UserID string `bson:"userID"`
}

func main() {
	fmt.Print("inside of main.go")
	// host := "127.0.0.1:4201"
	// if err := http.ListenAndServe(host, httpHandler()); err != nil {
	// 	fmt.Print("Failed to listen to " + host)
	// 	log.Fatalf("Failed to listen on %s: %v", host, err)
	// } else {
	// 	fmt.Print("Listening to " + host)
	// }

	//DB testing section
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Disconnect(context.TODO())

	// collection := client.Database("CEN3031_Test").Collection("TestStructure")
}

// httpHandler creates the backend HTTP router for queries, types,
// and serving the Angular frontend.
// func httpHandler() http.Handler {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/hello-world", helloWorld)
// 	//routing for retrieving taskList and addTask
// 	router.HandleFunc("/tasks/{userID}", handleGetTasksByUserID).Methods("GET")
// 	router.HandleFunc("/tasks", handleAddTask).Methods("POST")
// 	router.HandleFunc("/tasks/status", handleUpdateStatus).Methods("PUT")
// 	router.HandleFunc("/tasks/delete", handleDeleteTask).Methods("DELETE")
// 	router.HandleFunc("/tasks/description", handleGetTaskByDescriptionAndUserID).Methods("GET")

// 	// WARNING: this route must be the last route defined.

// 	router.PathPrefix("/").Handler(AngularHandler).Methods("GET")

// 	return handlers.LoggingHandler(os.Stdout,
// 		handlers.CORS(
// 			handlers.AllowCredentials(),
// 			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
// 				"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
// 				"Cache-Control", "Content-Range", "Range"}),
// 			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
// 			handlers.AllowedOrigins([]string{"http://localhost:4200"}), // maybe should be 4020???
// 			handlers.ExposedHeaders([]string{"DNT", "Kxeep-Alive", "User-Agent",
// 				"X-Requested-With", "If-Modified-Since", "Cache-Control",
// 				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
// 			handlers.MaxAge(86400),
// 		)(router))
// }

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

// new getTasksByUserID also takes in a listname to account for seperate tasklists
func getTasksByUserID(userID string, listName string) ([]Task, error) {
	//connect to the MongoDB server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.TODO())

	//get the collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//define a filter that matches tasks with the specified userID and list name
	filter := bson.M{"userID": userID, "taskLists.name": listName}

	//query the collection for tasks that match the filter
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	//decode the results into a list of task structs
	var tasks []Task
	for cursor.Next(context.TODO()) {
		//creating a taskList to match and return
		var result struct {
			TaskLists []struct {
				Name  string `bson:"name"`
				Tasks []Task `bson:"tasks"`
			} `bson:"taskLists"`
		}
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		for _, taskList := range result.TaskLists {
			if taskList.Name == listName {
				tasks = taskList.Tasks
				break
			}
		}
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	//return the list of tasks
	return tasks, nil
}

func getTaskByDescriptionUserIDAndListName(description string, userID string, listName string) (Task, error) {
	//connect to the MongoDB server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return Task{}, err
	}
	defer client.Disconnect(context.TODO())

	//get the collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//define a filter that matches tasks with the specified description, userID, and listName
	filter := bson.M{"description": description, "userID": userID, "listName": listName}

	//query the collection for a task that matches the filter
	result := collection.FindOne(context.Background(), filter)

	//decode the result into a Task struct
	var task Task
	err = result.Decode(&task)
	if err != nil {
		return Task{}, err
	}

	//return the task
	return task, nil
}

// find by description and userID
func getTaskByDescriptionAndUserID(description string, userID int) ([]Task, error) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, fmt.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//find all tasks with the given userID and description
	filter := bson.M{"userid": userID, "description": description}
	cursor, err := collection.Find(context.Background(), filter)

	if err != nil {
		return nil, fmt.Errorf("Error finding tasks in database: %v", err)
	}
	defer cursor.Close(context.Background())

	//decode the cursor into a slice of Task objects
	var tasks []Task
	err = cursor.All(context.Background(), &tasks)
	if err != nil {
		return nil, fmt.Errorf("Error decoding tasks from database: %v", err)
	}

	return tasks, nil
}

func addTask(description, status, userID, listName string) error {
	//create a new Task struct with string arguments
	task := Task{
		Description: description,
		Status:      status,
		UserID:      userID,
		ListName:    listName,
	}

	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//add the task to the task list
	_, err = collection.InsertOne(context.Background(), task)
	if err != nil {
		return fmt.Errorf("error adding task to task list: %v", err)
	}

	return nil
}

func deleteTask(description, ListName string) error {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return fmt.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//delete the task from the database
	filter := bson.M{"description": description, "listName": ListName}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("Error deleting task from database: %v", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("Task not found in database")
	}

	return nil
}

func updateTaskStatus(description string, newStatus string, listName string) error {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return fmt.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//update the task status in the database
	status := bson.M{"description": description}
	filter := bson.M{"$and": []bson.M{status, {"listName": listName}}}
	update := bson.M{"$set": bson.M{"status": newStatus}}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return fmt.Errorf("Error updating task status in database: %v", err)
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("Task not found in database")
	}

	return nil
}

func getListNames(userID string) ([]string, error) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, fmt.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the task lists collection
	collection := client.Database("CEN3031_Test").Collection("TaskLists")

	//create a filter to find the task lists for the specified user
	filter := bson.M{"userID": userID}

	//query the database for the task lists
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("Error finding task lists in database: %v", err)
	}
	defer cursor.Close(context.Background())

	//create an empty map to store unique list names
	listNamesMap := make(map[string]bool)

	//iterate through the results of the database query
	for cursor.Next(context.Background()) {
		//create a TaskList object to store the decoded result
		var taskList TaskList
		//decode the current result and store it in the TaskList object
		if err := cursor.Decode(&taskList); err != nil {
			return nil, fmt.Errorf("Error decoding task list: %v", err)
		}
		//add the current list name to the map if it hasn't already been added
		listNamesMap[taskList.ListName] = true
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("Error reading task lists from cursor: %v", err)
	}

	//convert the map to a slice by iterating over the keys of the map and appending each key (which is a list name) to the slice
	var listNames []string
	for listName := range listNamesMap {
		listNames = append(listNames, listName)
	}

	return listNames, nil
}

// Handler Funcs

//FIX FOR LISTNAME

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get task details
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
		return
	}

	//get task details from form
	description := r.FormValue("description")
	status := r.FormValue("status")
	userID := r.FormValue("userID")
	listName := r.FormValue("listName")

	//add task to database
	err = addTask(description, status, userID, listName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error adding task to database: %v", err), http.StatusInternalServerError)
		return
	}

	//return success message
	fmt.Fprintf(w, "Task added successfully")
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	//get the description and list name from the query parameters
	description := r.URL.Query().Get("description")
	listName := r.URL.Query().Get("listName")

	//delete the task from the database
	err := deleteTask(description, listName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//write a success message to the response
	fmt.Fprintf(w, "Task deleted successfully")
}

func updateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	//extracts variables from mux package rather than from URL itself for testing purposes
	params := mux.Vars(r)
	description := params["description"]
	newStatus := params["newStatus"]
	listName := params["listName"]

	err := updateTaskStatus(description, newStatus, listName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Task status updated successfully")
}

func handleGetTaskByDescriptionAndUserID(w http.ResponseWriter, r *http.Request) {
	//get the description and userID from the request parameters
	description := r.URL.Query().Get("description")
	userIDStr := r.URL.Query().Get("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	//call the getTaskByDescriptionAndUserID function
	tasks, err := getTaskByDescriptionAndUserID(description, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//encode the tasks as JSON and write to the response
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "Error encoding tasks to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getTasksByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	//get the userID and listName from the query parameters
	userID := r.URL.Query().Get("userID")
	listName := r.URL.Query().Get("listName")

	//call the getTasksByUserID function to retrieve the tasks
	tasks, err := getTasksByUserID(userID, listName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//convert the tasks to JSON
	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//set the content type and write the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTasks)
}

// SPOTIFY FUNCTIONS
func addPlaylist(playlistName string, songs [100]Song, userID string) error {
	playlist := Playlist{
		PlaylistName: playlistName,
		Songs:        songs,
		UserID:       userID,
	}
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

func addSongToSpotifyPlaylist(playlistName string, songs [100]Song, userID string) error {
	//create new Song  struct with string arguments
	// song := Song{
	// 	Song:   songName,
	// 	Artist: artistName,
	// 	url:    trackURL,
	// 	UserID: userID,
	// }
	// playlist := Playlist{
	// 	PlaylistName: playlistName,
	// 	Songs:        songs,
	// 	UserID:       userID,
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
			theSongs[i+1].Artist = songs[0].Artist
			theSongs[i+1].Song = songs[0].Song
			break
		}
	}
	//fmt.Printf(theSongs[1].Artist)

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

func removeSongFromPlaylist(playlistName string, song Song, userID string) error {
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
		if theSongs[i].Artist == song.Artist && theSongs[i].Song == song.Song {
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

func updateSongFromPlaylist(playlistName string, song Song, UserID string, newSong string, newArtist string, newUrl string) error {
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
		if theSongs[i].Artist == song.Artist && theSongs[i].Song == song.Song {
			theSongs[i].Artist = newArtist
			theSongs[i].Song = newSong
			theSongs[i].url = newUrl
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
func deletePlaylist(playlistName string, userID string) error {
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

func getPlaylist(playlistName string, userID string) error {
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
