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

    getTaskLists(userId: string): Observable<TaskList[]> {
        const url = `${this.api}/tasklists?userId=${userId}`;
        return this.http.get<TaskList[]>(url);
      }
    
      getTaskList(userId: string, listName: string): Observable<TaskList> {
        const url = `${this.api}/tasklists/${userId}/${listName}`;
        return this.http.get<TaskList>(url);
      }
    
      addTaskList(taskList: TaskList): Observable<TaskList> {
        const url = `${this.api}/tasklists`;
        return this.http.post<TaskList>(url, taskList);
      }
    
      deleteTaskList(userId: string, listName: string): Observable<any> {
        const url = `${this.api}/tasklists/${userId}/${listName}`;
        return this.http.delete(url);
      }
    
      updateTaskList(taskList: TaskList): Observable<TaskList> {
        const url = `${this.api}/tasklists/${taskList.userId}/${taskList.name}`;
        return this.http.put<TaskList>(url, taskList);
      }
    
      getTasks(userId: string, listName: string): Observable<any> {
        const url = `${this.api}/tasks?userId=${userId}&listName=${listName}`;
        return this.http.get(url);
      }
    
      getTask(userId: string, listName: string, taskId: string): Observable<any> {
        const url = `${this.api}/tasks/${userId}/${listName}/${taskId}`;
        return this.http.get(url);
      }
    
      addTask(userId: string, listName: string, task: any): Observable<any> {
        const url = `${this.api}/tasks`;
        return this.http.post(url, task);
      }
    
      deleteTask(userId: string, listName: string, taskId: string): Observable<any> {
        const url = `${this.api}/tasks/${userId}/${listName}/${taskId}`;
        return this.http.delete(url);
      }
    
      updateTask(userId: string, listName: string, taskId: string, task: any): Observable<any> {
        const url = `${this.api}/tasks/${userId}/${listName}/${taskId}`;
        return this.http.put(url, task);
      }
}