import { Component } from '@angular/core';
import { AuthService } from 'src/app/shared/services/auth';
import { Book, SignInService } from 'src/app/sign-in.service';



@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css']
})
export class SignInComponent {
  public books: Book[] = [];
  
private count = 0;  

  constructor(
    public authService: AuthService,
    private bookService: SignInService, 
  ) {
    // I think the get is messing everything up 
    this.books = bookService.get();
      //this.bookService.getBooks().subscribe(res => {
      //  this.books = res; 
      //})
   }

   getAllBooks() {
      this.bookService.getBooks().subscribe(res => {
        this.books = res; 
      });
   }

   // need to make sure this works with paramaters
   // so the front end can input information and add it to the database 
  addBookTwo(TitleName: string, AuthorName: string) { 
    console.log('inside of addBook with parameters')
    this.books = this.bookService.add({
      title: TitleName,
      author: AuthorName
    })
  }
  
   addBook() {    //(title: string, author: string) {
    console.log('inside of addBook')
    //this.bookService.addBook(title, author);
        /*this.books = this.bookService.add({
          title:  `Harry Potter ${this.count}`,
          author: 'Ana Maria'
        }) */
    this.bookService.addBook({
        title:  `Harry Potter ${this.count}`,
        author: 'Ana Maria'
    })
    this.count++; 
    /*
      .then(
      this.books => {
        this.books = books;
      }, err => {
        
      }
    )
    */
  }
}