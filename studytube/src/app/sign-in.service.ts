import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

/*export class Book { // what is the difference between interface and a class?
  constructor(
    public title: string,
    public author: string, 
  ) {}
} */
export interface Book {
  title: string; 
  author: string; 
}

@Injectable({
  providedIn: 'root'
})
export class SignInService {
  api: string = 'https://jsonplaceholder.typicode.com/posts';
  constructor(private http: HttpClient) {}

  getBooks():  Observable<Book[]> {  
    return this.http.get<Book[]>(this.api);
  }
  
  addBook(title: string, author: string) { //addBook(book : Book) {
    //return this.http.post<Book>(this.api, book)
    return this.http.post<Book[]>("/books", {
      title, author
    })
  }


}



