import { Component, OnInit } from '@angular/core';
import { SpotifyPlayerComponent } from '../spotify-player/spotify-player.component';
import { NgModule } from '@angular/core';
import { AuthService } from 'src/app/shared/services/auth';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css'],
})

export class DashboardComponent implements OnInit {
  title = 'Dashboard';
  constructor(public authService: AuthService) {}
  ngOnInit(): void {}
};