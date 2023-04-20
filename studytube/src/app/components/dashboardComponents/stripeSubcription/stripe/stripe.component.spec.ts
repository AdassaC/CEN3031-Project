import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StripeComponent } from './stripe.component';

describe('StripeComponent', () => {
  let component: StripeComponent;
  let fixture: ComponentFixture<StripeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ StripeComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(StripeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create Stripe', async() => {
    const fixture = TestBed.createComponent(StripeComponent);
    const dashboard = fixture.componentInstance;
    expect(dashboard).toBeTruthy();
  });

  it(`should have as title 'Become Members'`, async() => {
    const fixture = TestBed.createComponent(StripeComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('Become Member');
  });

  it('should render "Become Member" in a title tag', async() => {
    const fixture = TestBed.createComponent(StripeComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('title').textContent).toContain('Become Member');
  });
});
