import { Component } from '@angular/core';
import { FormControl, FormGroup, FormBuilder, FormArray, Form } from '@angular/forms';
import { Validators } from '@angular/forms';
import { MAT_FORM_FIELD, MatFormField, MatFormFieldControl } from '@angular/material/form-field';
import { EventEmitter, Output } from '@angular/core';
import { SpotifyService } from 'src/app/spotify.service';
import { AuthService } from 'src/app/shared/services/auth';
import { StripeService } from 'src/app/stripe.service';

@Component({
  selector: 'app-update-playlist',
  templateUrl: './update-playlist.component.html',
  styleUrls: ['./update-playlist.component.css',  '../dashboard.scss']
})
export class UpdatePlaylistComponent {
  constructor(
    public authService: AuthService, 
    private spotifyService : SpotifyService, 
    public fb: FormBuilder
  ) {}

  title = "Update Playlist";
  @Output() submit = new EventEmitter();

  updatePlaylist = this.fb.group({
    playlistName:['', Validators.required],
    numberOfSongs: [''],
    genre1: [''],
    genre2: [''],
    
    userGenres: this.fb.array([
      this.fb.control('')
    ])
  });
  
  onSubmit(pName: string) {
    // TODO: Use EventEmitter with form value
    /*console.warn(this.playlistForm.value);
    this.submit.emit(this.playlistForm.value);*/

    this.spotifyService.createPlaylist(
      pName
    )
    alert("Playlist has been updated!");
  }

}
