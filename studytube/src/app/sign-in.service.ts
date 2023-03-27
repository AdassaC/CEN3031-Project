import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

class Book { 
  constructor(
    public title: string,
    public author: string, 
  ) {}
}

@Injectable({
  providedIn: 'root'
})
export class SignInService {
  api: string = 'https://jsonplaceholder.typicode.com/posts';
  constructor(private http: HttpClient) {}

  getBooks() {  
    return this.http.get<Book[]>(this.api);
  }
  
  postBook(book : Book) {
    return this.http.post<Book>(this.api, book)
  }


}



