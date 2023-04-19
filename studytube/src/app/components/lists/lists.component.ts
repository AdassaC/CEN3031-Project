import { Component, OnInit } from '@angular/core';
import { NgModule } from '@angular/core';
import { AuthService } from 'src/app/shared/services/auth';

@Component({
  selector: 'app-lists',
  templateUrl: './lists.component.html',
  styleUrls: ['./lists.component.css']
})
export class ListsComponent {
  constructor(public authService: AuthService) {}

}
