import { Component } from '@angular/core';
import { StripeService } from 'src/app/stripe.service';
import { AuthService } from 'src/app/shared/services/auth';

@Component({
  selector: 'app-subscribe',
  templateUrl: './subscribe.component.html',
  styleUrls: ['./subscribe.component.css', '../../dashboard.scss']
})
export class SubscribeComponent {
  ngOnInit(): void {}
  constructor(
    private stripeService: StripeService,
    public authService: AuthService
  ) {

  }
  addCustomer(customerName: string, phoneNumber: string) {
    console.log("This is the name of the Customer: " + customerName);
    this.stripeService.addCustomerToStripe(
      customerName,
      phoneNumber,
    )
  }
  createSubscription(paymentID: string, customerID: string, priceID: string) {
    this.stripeService.createSubscriptionToStripe(
      paymentID,
      customerID, 
      priceID
    )
    alert("Subscription has been added to your account.");
  }
  cancelSubscription(subscriptionID: string) {
      this.stripeService.cancelSubscriptionToStripe(
        subscriptionID
      )
  } 

  updateSubscription(subscriptionID: string, priceID: string) {
    this.stripeService.updateSubscriptionToStripe(
      subscriptionID,
      priceID
    )
  }
}
