import { ComponentFixture, TestBed } from '@angular/core/testing';


import { PlaylistGeneratorComponent } from './playlist-generator.component';

describe('PlaylistGeneratorComponent', () => {
  let component: PlaylistGeneratorComponent;
  let fixture: ComponentFixture<PlaylistGeneratorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PlaylistGeneratorComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PlaylistGeneratorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
