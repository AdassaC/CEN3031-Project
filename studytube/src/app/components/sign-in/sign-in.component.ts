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
  
  constructor(
    public authService: AuthService,
    private bookService: SignInService, 
  ) {
    //this.books = bookService.getBooks();
    this.bookService.getBooks().subscribe(res => {
      this.books = res; 
    })
   }

  addBook(title: string, author: string) {
    this.bookService.addBook(title, author);
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