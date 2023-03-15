import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SignInComponent } from './sign-in.component';

describe('SignInComponent', () => {
  let component: SignInComponent;
  let fixture: ComponentFixture<SignInComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SignInComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SignInComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create Sign-In', async() => {
    const fixture = TestBed.createComponent(SignInComponent);
    const register = fixture.componentInstance;
    expect(register).toBeTruthy();
  });

  it(`should have as title 'User Profile'`, async() => {
    const fixture = TestBed.createComponent(SignInComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('User Profile');
  });

  it('should render "title" in a a tag', async() => {
    const fixture = TestBed.createComponent(SignInComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('a').textContent).toContain('User Profile');
  });
});
