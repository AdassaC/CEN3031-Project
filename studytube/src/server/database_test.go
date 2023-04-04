package main

import (
	"context"
	"fmt"
	"reflect"

	// "fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//UNIT TESTING FUNCTIONS

func TestConnectToDB(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)

	if err != nil {
		t.Fatalf("error pinging MongoDB server: %v", err)
	}
}
func TestAddPlaylist(t *testing.T) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	song := Song{
		Song:   "Forever",
		Artist: "Drake",
		url:    "Spotify test URL",
	}

	var songs [100]Song
	songs[0] = song

	UserID := "12345"

	//add the task to the database
	err = addPlaylist("tempPlayListName", songs, UserID)
	if err != nil {
		t.Errorf("Error adding task to task list: %v", err)
	}
}

func TestAddSong(t *testing.T) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection

	song := Song{
		Song:   "I wonder",
		Artist: "Kanye",
		url:    "Spotify test URL",
	}
	songTwo := Song{
		Song:   "Sick and Tired",
		Artist: "iann dior",
		url:    "Spotify test URL",
	}

	var songs [100]Song
	songs[0] = song
	//songs[0] = songTwo

	//test

	songs[1] = songTwo

	UserID := "12345"

	//add the task to the database
	err = addSongToSpotifyPlaylist("tempPlayListName", songs, UserID)
	if err != nil {
		t.Errorf("Error adding task to task list: %v", err)
	}
}
func TestUpdateSong(t *testing.T) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	song := Song{
		Song:   "Forever",
		Artist: "Drake",
		url:    "Spotify test URL",
	}

	var songs [100]Song
	songs[0] = song

	UserID := "12345"

	newSong := "Memories"
	newArtist := "Thutmose"
	newUrl := "Apple Music Test URL"

	//remove the task from the database
	err = updateSongFromPlaylist("tempPlayListName", song, UserID, newSong, newArtist, newUrl)
	if err != nil {
		t.Errorf("Error adding task to task list: %v", err)
	}
}
func TestRemoveSong(t *testing.T) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	song := Song{
		Song:   "Memories",
		Artist: "Thutmose",
		url:    "Spotify test URL",
	}

	var songs [100]Song
	songs[0] = song

	UserID := "12345"

	//remove the task from the database
	err = removeSongFromPlaylist("tempPlayListName", song, UserID)
	if err != nil {
		t.Errorf("Error adding task to task list: %v", err)
	}
}
func TestGetPlaylist(t *testing.T) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	UserID := "12345"

	getPlaylist("tempPlayListName", UserID)
}
func TestDeletePlaylist(t *testing.T) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	UserID := "12345"

	deletePlaylist("tempPlayListName", UserID)
}

func TestAddTaskToList(t *testing.T) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//create a new test task
	task := Task{
		Description: "Test Task",
		Status:      "Not Started",
		UserID:      "123456",
		ListName:    "Test List",
	}

	//add the task to the database
	err = addTask(task.Description, task.Status, task.UserID, task.ListName)
	if err != nil {
		t.Errorf("Error adding task to task list: %v", err)
	}

	//check that the task was added to the database
	filter := bson.M{
		"description": task.Description,
		"userID":      task.UserID,
		"listName":    task.ListName,
	}
	var result Task
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		t.Errorf("Error finding task in database: %v", err)
	}

	// check that the task was added correctly to the database
	if result.Description != task.Description || result.Status != task.Status || result.UserID != task.UserID || result.ListName != task.ListName {
		t.Errorf("Task not added correctly to database")
	}
}

func TestDeleteTask(t *testing.T) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//create a new test task
	task := Task{
		Description: "Test Task",
		Status:      "Not Started",
		UserID:      "123456",
		ListName:    "Test List",
	}

	//add the task to the database
	// err = addTask(task.Description, task.Status, task.UserID, task.ListName)
	// if err != nil {
	// 	t.Errorf("Error adding task to task list: %v", err)
	// }

	//delete the test task from the database
	err = deleteTask(task.Description, task.ListName)
	if err != nil {
		t.Errorf("Error deleting task from database: %v", err)
	}

	//check that the task was deleted from the database
	filter := bson.M{"description": "Test Task 2", "listName": "Test List"}
	err = collection.FindOne(context.Background(), filter).Decode(&task)
	if err == nil {
		t.Errorf("Task not deleted from database")
	}
}

func TestGetTaskByDescriptionUserIDAndListName(t *testing.T) {
	// Set up test database and collection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//insert test task into database
	testTask := Task{
		ID:          "1",
		Description: "Test task",
		Status:      "Incomplete",
		UserID:      "user123",
		ListName:    "Test list",
	}
	_, err = collection.InsertOne(context.Background(), testTask)
	if err != nil {
		t.Fatalf("Error inserting test task: %v", err)
	}

	//test finding the task by description, user ID, and list name
	foundTask, err := getTaskByDescriptionUserIDAndListName("Test task", "user123", "Test list")
	if err != nil {
		t.Fatalf("Error finding task: %v", err)
	}

	//comparative test
	if foundTask.ID != "1" || foundTask.Status != "Incomplete" {
		t.Errorf("Expected userID 1 and status Incomplete, got ID '%s' and status '%s'", foundTask.ID, foundTask.Status)
	}
	fmt.Printf("Description: %v", foundTask.Description)
}

func TestUpdateTaskStatus(t *testing.T) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//create a new test task
	task := Task{
		Description: "Test Task",
		Status:      "Not Started",
		UserID:      "123456",
		ListName:    "Test List",
	}

	//add the task to the database
	err = addTask(task.Description, task.Status, task.UserID, task.ListName)
	if err != nil {
		t.Errorf("Error adding task to task list: %v", err)
	}

	//update the task status in the database
	err = updateTaskStatus(task.Description, "In Progress", task.ListName)
	if err != nil {
		t.Errorf("Error updating task status in database: %v", err)
	}

	//check that the task status was updated in the database
	filter := bson.M{"description": "Test Task"}
	var updatedTask Task
	err = collection.FindOne(context.Background(), filter).Decode(&updatedTask)
	if err != nil {
		t.Errorf("Error finding task in database: %v", err)
	}
	if updatedTask.Status != "In Progress" {
		t.Errorf("Task status not updated in database")
	}
}

func TestGetListName(t *testing.T) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//test

	//initialize test data
	userID := "test-user"
	expectedListNames := []string{"list1", "list2", "list3"}
	for _, listName := range expectedListNames {
		taskList := TaskList{
			UserID:   userID,
			ListName: listName,
		}
		_, err := client.Database("CEN3031_Test").Collection("TestStructure").InsertOne(context.Background(), taskList)
		if err != nil {
			t.Fatalf("Error inserting test data: %v", err)
		}
	}

	//call getListNames and verify the results
	listNames, err := getListNames(userID)
	if err != nil {
		t.Fatalf("Error calling getListNames: %v", err)
	}
	//reflect.DeepEqual to compare the two lists
	if !reflect.DeepEqual(listNames, expectedListNames) {
		t.Fatalf("Expected list names %v, but got %v", expectedListNames, listNames)
	}
}
