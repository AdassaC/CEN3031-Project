import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateListComponent } from './create-list.component';

describe('CreateListComponent', () => {
  let component: CreateListComponent;
  let fixture: ComponentFixture<CreateListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create CreateList', async() => {
    const fixture = TestBed.createComponent(CreateListComponent);
    const dashboard = fixture.componentInstance;
    expect(dashboard).toBeTruthy();
  });

  it(`should have as title 'Create New List'`, async() => {
    const fixture = TestBed.createComponent(CreateListComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('Dashboard');
  });

  it('should render "Create New List in a title tag', async() => {
    const fixture = TestBed.createComponent(CreateListComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('title').textContent).toContain('Create New List');
  });
});
