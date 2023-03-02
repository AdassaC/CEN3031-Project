package main

import (
	"context"
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

func TestAddTask(t *testing.T) {
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
	}

	//add the task to the database
	err = addTask(task.Description, task.Status, task.UserID)
	if err != nil {
		t.Errorf("Error adding task to task list: %v", err)
	}

	//check that the task was added to the database
	filter := bson.M{"description": "Test Task"}
	var result Task
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		t.Errorf("Error finding task in database: %v", err)
	}
	//check
	if result.Description != task.Description || result.Status != task.Status || result.UserID != task.UserID {
		t.Errorf("Task not added correctly to database")
	}
}

func TestGetTasksByUserID(t *testing.T) {
	//connect to MongoDB on localhost port 27017
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	//get a handle to the tasks collection
	collection := client.Database("CEN3031_Test").Collection("TestStructure")

	//add some test tasks to the database
	task1 := Task{
		Description: "Test Task 1",
		Status:      "Not Started",
		UserID:      "1234567",
	}
	task2 := Task{
		Description: "Test Task 2",
		Status:      "Not Started",
		UserID:      "12345678",
	}
	_, err = collection.InsertMany(context.Background(), []interface{}{task1, task2})
	if err != nil {
		t.Errorf("Error adding test tasks to database: %v", err)
	}

	//get tasks for user "123456"
	tasks, err := getTasksByUserID("1234567")
	if err != nil {
		t.Errorf("Error getting tasks from database: %v", err)
	}

	//check that the correct tasks were retrieved
	if len(tasks) != 1 {
		t.Errorf("Incorrect number of tasks retrieved from database")
	}
	if tasks[0].Description != task1.Description || tasks[0].Status != task1.Status || tasks[0].UserID != task1.UserID {
		t.Errorf("Incorrect task retrieved from database")
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
	}

	//add the task to the database
	// err = addTask(task.Description, task.Status, task.UserID)
	// if err != nil {
	// 	t.Errorf("Error adding task to task list: %v", err)
	// }

	//delete the test task from the database
	err = deleteTask(task.Description)
	if err != nil {
		t.Errorf("Error deleting task from database: %v", err)
	}

	//check that the task was deleted from the database
	filter := bson.M{"description": "Test Task"}
	err = collection.FindOne(context.Background(), filter).Decode(&task)
	if err == nil {
		t.Errorf("Task not deleted from database")
	}
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
	}

	//add the task to the database
	err = addTask(task.Description, task.Status, task.UserID)
	if err != nil {
		t.Errorf("Error adding task to task list: %v", err)
	}

	//update the task's status
	newStatus := "In Progress"
	err = updateTaskStatus(task.Description, newStatus)
	if err != nil {
		t.Errorf("Error updating task status: %v", err)
	}

	//check that the task's status was updated in the database
	status := bson.M{"description": "Test Task"}
	var result Task
	err = collection.FindOne(context.Background(), status).Decode(&result)
	if err != nil {
		t.Errorf("Error finding task in database: %v", err)
	}
	if result.Status != newStatus {
		t.Errorf("Task status not updated correctly in database")
	}

	//delete the test task from the database
	//decided to add this since most test tasks were getting too similar, will add to other tests eventually or a command to clear out the test DB
	// err = deleteTask(task.Description)
	// if err != nil {
	// 	t.Errorf("Error deleting task from database: %v", err)
	// }
}
