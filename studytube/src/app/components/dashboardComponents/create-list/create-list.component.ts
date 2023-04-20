import { Component } from '@angular/core';
import { AuthService } from 'src/app/shared/services/auth';

@Component({
  selector: 'app-create-list',
  templateUrl: './create-list.component.html',
  styleUrls: ['./create-list.component.css', '../dashboard.scss']
})
export class CreateListComponent {
  constructor(public authService: AuthService) {}

}
