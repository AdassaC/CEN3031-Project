import { Component, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { __param } from 'tslib';

@Component({
  selector: 'app-task-view',
  templateUrl: './task-view.component.html',
  styleUrls: ['../task-styles.scss','./task-view.component.css'],
  encapsulation: ViewEncapsulation.ShadowDom
})
export class TaskViewComponent {
  /*lists: any[];
  //each list should have a title

  tasks: any[];

  constructor(private taskService: TaskService, private route: ActivatedRoute) {}

  ngOnInit() {
    this.route.params.subscribe(
      (params: Params) => {
        console.log(params);
        
        this.taskService.getTasks(params.listId).subscribe((tasks: any[]) => {
          this.tasks = tasks;
        })
      }
    )

    this.taskService.getLists().subscribe((lists: any[]) => {
      this.lists = lists;
    })

  }*/

}


