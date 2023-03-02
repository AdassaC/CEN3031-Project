import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService } from 'src/app/shared/services/auth';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  constructor(private route: ActivatedRoute, private router: Router, public authService: AuthService) { }
  title = 'register-user';
  goToSignIn(): void {
    this.router.navigate(['./login']);
  }
}
