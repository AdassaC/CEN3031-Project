import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { TaskList } from './tasklist.model';

  export interface Task {
    Description: string; 
    Status: string; 
    UserID: string; 
    ListName: string;
  }
  
  export class TaskItem {
    Description: string | undefined;
    Status: string | undefined;
    UserID: string | undefined;
    ListName: string | undefined;
  }

@Injectable({
    providedIn: 'root'
  })
export class Database{

    api: string = 'http://localhost:4201/';
    constructor(private http: HttpClient) {}

    //getTaskList by userID
    getTaskLists(userId: string): Observable<TaskList[]> {
        const url = `${this.api}/tasklists?userId=${userId}`;
        return this.http.get<TaskList[]>(url);
      }
    
      //getTaskList by userID and listName
      getTaskList(userId: string, listName: string): Observable<TaskList> {
        const url = `${this.api}/tasklists/${userId}/${listName}`;
        return this.http.get<TaskList>(url);
      }

      //adds a taskList
      addTaskList(taskList: TaskList): Observable<TaskList> {
        const url = `${this.api}/tasklists`;
        return this.http.post<TaskList>(url, taskList);
      }

      //delete an entire list by userID and a listName
      deleteTaskList(userId: string, listName: string): Observable<any> {
        const url = `${this.api}/tasklists/${userId}/${listName}`;
        return this.http.delete(url);
      }

      //update a tasklist by replacing it
      updateTaskList(taskList: TaskList): Observable<TaskList> {
        const url = `${this.api}/tasklists/${taskList.userId}/${taskList.name}`;
        return this.http.put<TaskList>(url, taskList);
      }

      //returns a group of tasks by userID and listName
      getTasks(userId: string, listName: string): Observable<any> {
        const url = `${this.api}/tasks?userId=${userId}&listName=${listName}`;
        return this.http.get(url);
      }

      //retrieves a single task by the userID and listName with taskID
      getTask(userId: string, listName: string, taskId: string): Observable<any> {
        const url = `${this.api}/tasks/${userId}/${listName}/${taskId}`;
        return this.http.get(url);
      }
      
      //adds a task by userID and listname 
      addTask(userId: string, listName: string, task: any): Observable<any> {
        const url = `${this.api}/tasks`;
        return this.http.post(url, task);
      }
      
      //deletes a specific task by userID, listname and taskID
      deleteTask(userId: string, listName: string, taskId: string): Observable<any> {
        const url = `${this.api}/tasks/${userId}/${listName}/${taskId}`;
        return this.http.delete(url);
      }

      //updates a singular task by taking in a userID, listName, taskID and a task
      updateTask(userId: string, listName: string, taskId: string, task: any): Observable<any> {
        const url = `${this.api}/tasks/${userId}/${listName}/${taskId}`;
        return this.http.put(url, task);
      }
}