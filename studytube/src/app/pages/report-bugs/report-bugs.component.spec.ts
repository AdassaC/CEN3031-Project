import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReportBugsComponent } from './report-bugs.component';

describe('ReportBugsComponent', () => {
  let component: ReportBugsComponent;
  let fixture: ComponentFixture<ReportBugsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReportBugsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReportBugsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create ReportBugs', async() => {
    const fixture = TestBed.createComponent(ReportBugsComponent);
    const rbugs = fixture.componentInstance;
    expect(rbugs).toBeTruthy();
  });
});
