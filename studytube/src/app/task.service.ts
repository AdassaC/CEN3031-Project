import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class TaskService {

  //lists: any[];

  //constructor(private webReqService: WebResuestService) { }
  createLists(title: string) {
    //send web request to create list

    //return this.webReqService.post('lists', {title});
  }

  getLists() {
    //return this.webReqService.get('lists');
    //this.lists = lists;
  }
}
