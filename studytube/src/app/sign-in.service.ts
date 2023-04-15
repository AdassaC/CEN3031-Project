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
  private _books: Book[] = [{
    title: "Test",
    author: "Testeee"
  }]

  api: string = 'http://localhost:4201/';
  constructor(private http: HttpClient) {}

  addBook(book : Book) {
    console.log("inside of the addBook in service");
    console.log(book);
    return this.http.post(this.api + "books", book)  //<Book>(this.api + "books", book)
    .subscribe((res) => {
      console.log(res);
    });
  }

  getBooks():  Observable<Book[]> {  
    console.log("inside service getBooks");
    //return this.http.get<Book[]>(this.api + "books");
    return this.http.get<Book[]>(this.api + "books/{HarryPotter}/page/{20}");
  }

  add(book : Book) { //(title: string, author: string) { //addBook(book : Book) {
    this._books.push(book);
    return this._books;
    //return this.http.post<Book[]>("/books", book)
  }

  get(): Book[]{
    return this._books;
  }

}



