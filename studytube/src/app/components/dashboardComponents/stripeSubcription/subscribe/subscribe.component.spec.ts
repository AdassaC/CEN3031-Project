import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SubscribeComponent } from './subscribe.component';

describe('SubscribeComponent', () => {
  let component: SubscribeComponent;
  let fixture: ComponentFixture<SubscribeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SubscribeComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SubscribeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create Subscribe', async() => {
    const fixture = TestBed.createComponent(SubscribeComponent);
    const dashboard = fixture.componentInstance;
    expect(dashboard).toBeTruthy();
  });

  it(`should have as title 'Add Subscription'`, async() => {
    const fixture = TestBed.createComponent(SubscribeComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('Add Subscription');
  });

  it('should render "Add Subscription" in a title tag', async() => {
    const fixture = TestBed.createComponent(SubscribeComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('title').textContent).toContain('Add Subscription');
  });
});
