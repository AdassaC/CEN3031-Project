import { Component, OnInit } from '@angular/core';
import { NgModule } from '@angular/core';
import { AuthService } from 'src/app/shared/services/auth';
import { StripeService } from 'src/app/stripe.service';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent {
  constructor(public authService: AuthService, private stripeService : StripeService) {}
  ngOnInit(): void {}

  cancelSubscription(subscriptionID: string) {
    this.stripeService.cancelSubscriptionToStripe(
      subscriptionID
    )
  }
} 

