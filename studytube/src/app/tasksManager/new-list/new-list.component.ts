import { Component } from '@angular/core';

@Component({
  selector: 'app-new-list',
  templateUrl: './new-list.component.html',
  styleUrls: ['../task-styles.scss','./new-list.component.css']
})
export class NewListComponent {
  title = 'Create a new list';
  /*constructor(private taskService: TaskService) {}

  ngOnInit() {}

  createList(title: string) {
    this.taskService.createList(title).subscribe((response: any) => {
      console.log(response);
    });
  
  }*/
}
