import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class StripeService {
  api: string = 'http://localhost:4201/';
  constructor(private http: HttpClient) {}

  addCustomerToStripe(name: string, phoneNumber: string) {
    //create-customer/name/tedsantiago@gmail.com/phone/603-690-8891
    return this.http.post(this.api + "create-customer/name/" + name + "/phone/" + phoneNumber, name)
    .subscribe((res) => {
      console.log(res);
    });
  }

  createSubscriptionToStripe(paymentID: string, customerID: string, priceID: string) {
    //create-subscription/pay/card_1Mxx78L5cDZvcnZ2NGSoOIi0/customer/cus_NjNDS9m1rbpJsj/price/price_1Mgr7ZL5cDZvcnZ2yI7cvRMH
    return this.http.post(this.api + "create-subscription/pay/" + paymentID + "/customer/" + customerID + "/price/" + priceID, priceID)
    .subscribe((res) => {
      console.log(res);
    });
  }

  cancelSubscriptionToStripe(subscriptionID: string) {
    //cancel-subscription/subscription/sub_1Mxxc4L5cDZvcnZ2QYbCi2lD
    return this.http.post(this.api + "cancel-subscription/subscription/" + subscriptionID, subscriptionID)
    .subscribe((res) => {
      console.log(res);
    });
  }

  updateSubscriptionToStripe(subscriptionID: string, priceID: string) {
    //update-subscription/subscription/sub_1MxyX6L5cDZvcnZ2dyUZAlqU/price/price_1MxyULL5cDZvcnZ2MXsMwWSn
    return this.http.post(this.api + "update-subscription/subscription/" + subscriptionID + "/price/" + priceID, subscriptionID)
    .subscribe((res) => {
      console.log(res);
    });
  }
}
