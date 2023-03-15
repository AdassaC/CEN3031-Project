import { Component } from '@angular/core';

export class ReportBugs {
  constructor(
    public id: number,
    public title: string, 
    public description: string
  ) {}
}

@Component({
  selector: 'app-report-bugs',
  templateUrl: './report-bugs.component.html',
  styleUrls: ['./report-bugs.component.css']
})
export class ReportBugsComponent {
  submitted = false;

  onSubmit() { this.submitted = true; }

}
