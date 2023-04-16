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

}
