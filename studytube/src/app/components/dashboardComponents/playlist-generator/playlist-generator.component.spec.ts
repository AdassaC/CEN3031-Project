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

  it('should create PlayistGenerator', async() => {
    const fixture = TestBed.createComponent(PlaylistGeneratorComponent);
    const register = fixture.componentInstance;
    expect(register).toBeTruthy();
  });

  it(`should have as title 'Playlist Generator'`, async() => {
    const fixture = TestBed.createComponent(PlaylistGeneratorComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('Playlist Generator');
  });

  it('should render "title" in a a tag', async() => {
    const fixture = TestBed.createComponent(PlaylistGeneratorComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).toContain('Create a playlist');
  });
});
