import { Component } from '@angular/core';
import { AuthService } from 'src/app/shared/services/auth';

@Component({
  selector: 'app-lists',
  templateUrl: './lists.component.html',
  styleUrls: ['./lists.component.css','../dashboard.scss']
})
export class ListsComponent {
  constructor(public authService: AuthService) {}
}
