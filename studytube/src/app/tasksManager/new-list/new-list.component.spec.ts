import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewListComponent } from './new-list.component';

describe('NewListComponent', () => {
  let component: NewListComponent;
  let fixture: ComponentFixture<NewListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NewListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create TaskView', async() => {
    const fixture = TestBed.createComponent(NewListComponent);
    const register = fixture.componentInstance;
    expect(register).toBeTruthy();
  });

  it(`should have as title 'TaskView'`, async() => {
    const fixture = TestBed.createComponent(NewListComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('TaskView');
  });

  it('should render "title" in a a tag', async() => {
    const fixture = TestBed.createComponent(NewListComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).toContain('User Profile');
  });
});
