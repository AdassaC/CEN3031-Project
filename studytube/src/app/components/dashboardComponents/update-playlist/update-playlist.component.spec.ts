import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UpdatePlaylistComponent } from './update-playlist.component';

describe('UpdatePlaylistComponent', () => {
  let component: UpdatePlaylistComponent;
  let fixture: ComponentFixture<UpdatePlaylistComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UpdatePlaylistComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UpdatePlaylistComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create UpdatePlaylistComponent', async() => {
    const fixture = TestBed.createComponent(UpdatePlaylistComponent);
    const register = fixture.componentInstance;
    expect(register).toBeTruthy();
  });

  it(`should have as title 'Update a Playlist'`, async() => {
    const fixture = TestBed.createComponent(UpdatePlaylistComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('Update a Playlist');
  });

  it('should render "title" in a a tag', async() => {
    const fixture = TestBed.createComponent(UpdatePlaylistComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).toContain('Update a Playlist');
  });
});
