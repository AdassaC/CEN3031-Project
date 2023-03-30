import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TaskViewComponent } from './task-view.component';

describe('TaskViewComponent', () => {
  let component: TaskViewComponent;
  let fixture: ComponentFixture<TaskViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TaskViewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TaskViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create TaskView', async() => {
    const fixture = TestBed.createComponent(TaskViewComponent);
    const register = fixture.componentInstance;
    expect(register).toBeTruthy();
  });

  it(`should have as title 'TaskView'`, async() => {
    const fixture = TestBed.createComponent(TaskViewComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('TaskView');
  });

  it('should render "title" in a a tag', async() => {
    const fixture = TestBed.createComponent(TaskViewComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('a').textContent).toContain('User Profile');
  });
});
