import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/shared/services/auth';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  //title: 'Sign In';
  constructor(
    public authService: AuthService
  ) { }
  ngOnInit() { }
}
