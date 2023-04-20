import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SettingsComponent } from './settings.component';

describe('SettingsComponent', () => {
  let component: SettingsComponent;
  let fixture: ComponentFixture<SettingsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SettingsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SettingsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create SettingsComponent', async() => {
    expect(component).toBeTruthy();
  });

  it(`should have as title 'Settings'`, async() => {
    const fixture = TestBed.createComponent(SettingsComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('Settings');
  });

  it('should render "Create New List in a title tag', async() => {
    const fixture = TestBed.createComponent(SettingsComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('title').textContent).toContain('Settings');
  });
});
