import { Component } from '@angular/core';
import { StripeService } from 'src/app/stripe.service';
import { AuthService } from 'src/app/shared/services/auth';

@Component({
  selector: 'app-stripe',
  templateUrl: './stripe.component.html',
  styleUrls: ['./stripe.component.css', '../../dashboard.scss']
})
export class StripeComponent {

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
