import { Component } from '@angular/core';
import { StripeService } from 'src/app/stripe.service';
import { AuthService } from 'src/app/shared/services/auth';

@Component({
  selector: 'app-subscribe',
  templateUrl: './subscribe.component.html',
  styleUrls: ['./subscribe.component.css', '../../dashboard.scss']
})
export class SubscribeComponent {
  constructor(
    private stripeService: StripeService,
    public authService: AuthService
  ) {}
}
