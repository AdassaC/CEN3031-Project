import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/shared/services/auth';


@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css']
})
export class SignInComponent implements OnInit {
  title = 'Sign In';
  constructor(
    public authService: AuthService
  ) { }
  ngOnInit() { }
}