package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"context"


	"github.com/adassacoimin/CEN3031-Project/studytube/src/server/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Temporary task struct
type Task struct {
	ID string `bson:"_id,omitempty"`
	Description string `bson:"description"`
	Status string `bson:"status"`
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello-world", helloWorld)

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

func TaskList() error {
	//connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return fmt.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	//get a handle to the tasks collection
	collection := client.Database("taskDB").Collection("tasks")

	//unfinished
}

func AddTask(task Task) error {
	//func for adding task to tasklist
	//connect to mongoDB on localhost port 27017
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return fmt.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())
	
	//get a handle to the tasks collection
	collection := client.Database("taskDB").Collection("tasks")

	//add the task to the task list
	_, err = collection.InsertOne(context.TODO(), task)
	if err != nil {
		return fmt.Errorf("Error adding task to task list: %v", err)
	}

	return nil

}

