package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/adassacoimin/CEN3031-Project/studytube/src/server/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// temporary task struct
type Task struct {
	ID          string `bson:"_id,omitempty"`
	Description string `bson:"description"`
	Status      string `bson:"status"`
	UserID      string `bson:"userID"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello-world", helloWorld)
	//routing for retrieving taskList and addTask
	r.HandleFunc("/tasks/{userID}", handleGetTasksByUserID).Methods("GET")
	r.HandleFunc("/tasks", handleAddTask).Methods("POST")
	r.HandleFunc("/tasks/status", handleUpdateStatus).Methods("PUT")
	r.HandleFunc("/tasks/delete", handleDeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/description", handleGetTaskByDescriptionAndUserID).Methods("GET")

	// Solves Cross Origin Access Issue
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})
	handler := c.Handler(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + os.Getenv("PORT"),
	}

	//DB testing section
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Disconnect(context.TODO())

	// collection := client.Database("CEN3031_Test").Collection("TestStructure")

	err = addTask("Random Task Wrong User", "Not Started", "123456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Task added successfully!")
	}

	err2 := addTask("Random Task 2", "Not Started", "1234567")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Task added successfully!")
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	//test getTaskByUserID
	tasks, err := getTasksByUserID("1234567")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Tasks for user 1234567:")
		for _, task := range tasks {
			fmt.Println(task.Description)
		}
	}

	//test updateStatus
	err = updateTaskStatus("Random Task 2", "Started")
	if err != nil {
		fmt.Println(err)
	}

	//test deleteTask
	err = deleteTask("Random Task Wrong User")
	if err != nil {
		fmt.Println(err)
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

func getTasksByUserID(userID string) ([]Task, error) {
	//connect to the MongoDB server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.TODO())

	//get the collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//define a filter that matches tasks with the specified userID
	filter := bson.M{"userID": userID}

	//query the collection for tasks that match the filter
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	//decode the results into a list of Task structs
	var tasks []Task
	for cursor.Next(context.TODO()) {
		var task Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	//return the list of tasks
	return tasks, nil
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

func addTask(description, status, userID string) error {
	//create new Task struct with string arguments
	task := Task{
		Description: description,
		Status:      status,
		UserID:      userID,
	}

	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return fmt.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//add the task to the task list
	_, err = collection.InsertOne(context.TODO(), task)
	if err != nil {
		return fmt.Errorf("Error adding task to task list: %v", err)
	}

	return nil
}

// similar functionality to addTask, query the DB for correct task and delete it
func deleteTask(description string) error {
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

	return nil
}

// similar functionality to addTask, query the DB for correct task and alter the status
func updateTaskStatus(description string, newStatus string) error {
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
	update := bson.M{"$set": bson.M{"status": newStatus}}
	result, err := collection.UpdateOne(context.Background(), status, update)
	if err != nil {
		return fmt.Errorf("Error updating task status in database: %v", err)
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("Task not found in database")
	}

	return nil
}

// Handler Funcs

func handleGetTasksByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	//retrieve list of tasks to be written to json with error checks
	tasks, err := getTasksByUserID(userID)
	if err != nil {
		//error handling WIP
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//put task into JSON for transfer
	jsonBytes, err := utils.StructToJSON(tasks)
	if err != nil {
		//error handling WIP
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//set content type to json
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func handleAddTask(w http.ResponseWriter, r *http.Request) {
	//parse the request body into a new Task struct, assuming incoming as json format
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		//error handling WIP
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//add the task to the database
	err = addTask(task.Description, task.Status, task.UserID)
	if err != nil {
		//error handling WIP
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//return a success message
	fmt.Fprint(w, "Task added successfully")
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	//get the description of the task to delete from the URL query parameter
	description := r.URL.Query().Get("description")
	if description == "" {
		http.Error(w, "Missing description parameter", http.StatusBadRequest)
		return
	}

	//call the deleteTask function to delete the task from the database
	err := deleteTask(description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//return a success message to the client
	fmt.Fprintf(w, "Task deleted successfully")
}

func handleUpdateStatus(w http.ResponseWriter, r *http.Request) {
	//parse the request body to get the description and new status values
	var requestBody struct {
		Description string `json:"description"`
		NewStatus   string `json:"new_status"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//call the updateTaskStatus function to update the task status in the database
	err = updateTaskStatus(requestBody.Description, requestBody.NewStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//return a success message to the client
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

func addPlaylist(playlistName string, songs [100]Song, userID string) error {
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

func addSongToSpotifyPlaylist(playlistName string, songs [100]Song, userID string) error {
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