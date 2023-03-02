package main

import (
	"context"
	"encoding/json"
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
		return
	}

	//put task into format for transfer
	jsonBytes, err := utils.StructToJSON(tasks)
	if err != nil {
		//error handling WIP
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
		return
	}

	//add the task to the database
	err = addTask(task.Description, task.Status, task.UserID)
	if err != nil {
		//error handling WIP
		return
	}

	//return a success message
	fmt.Fprint(w, "Task added successfully")
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {

}

func handleUpdateStatus(w http.ResponseWriter, r *http.Request) {

}
