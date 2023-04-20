import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListsComponent } from './lists.component';

describe('ListsComponent', () => {
  let component: ListsComponent;
  let fixture: ComponentFixture<ListsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ListsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create Lists', async() => {
    const fixture = TestBed.createComponent(ListsComponent);
    const dashboard = fixture.componentInstance;
    expect(dashboard).toBeTruthy();
  });

  it(`should have as title 'View Task Lists'`, async() => {
    const fixture = TestBed.createComponent(ListsComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('View Task Lists');
  });

  it('should render "View Task Lists" in a title tag', async() => {
    const fixture = TestBed.createComponent(ListsComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('title').textContent).toContain('View Task Lists');
  });
});
