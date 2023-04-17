import { Component } from '@angular/core';
import { StripeService } from 'src/app/stripe.service';

@Component({
  selector: 'app-stripe',
  templateUrl: './stripe.component.html',
  styleUrls: ['./stripe.component.css']
})
export class StripeComponent {
    constructor(
      private stripeService: StripeService,
    ) {

    }
    addCustomer(customerName: string) {
      console.log("This is the name of the Customer: " + customerName);
      this.stripeService.addCustomerToStripe(
        customerName,
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
