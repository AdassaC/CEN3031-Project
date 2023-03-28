import { Component, OnInit } from '@angular/core';
import { NgModule } from '@angular/core';
import { AuthService } from 'src/app/shared/services/auth';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent {
  constructor(public authService: AuthService) {}
  ngOnInit(): void {}
}
