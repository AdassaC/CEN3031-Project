import { Component } from '@angular/core';
import { FormControl, FormGroup, FormBuilder, FormArray, Form } from '@angular/forms';
import { Validators } from '@angular/forms';
import { MAT_FORM_FIELD, MatFormField, MatFormFieldControl } from '@angular/material/form-field';
import { EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-playlist-generator',
  templateUrl: './playlist-generator.component.html',
  styleUrls: ['./playlist-generator.component.css']
})

export class PlaylistGeneratorComponent {
  title = "Playlist Generator";
  @Output() submit = new EventEmitter();

  playlistForm = this.fb.group({
    playlistName:['', Validators.required],
    numberOfSongs: [''],
    genre1: [''],
    genre2: [''],
    
    userGenres: this.fb.array([
      this.fb.control('')
    ])
  });
  
  constructor(public fb: FormBuilder) {}

  get userGenres() {
    return this.playlistForm.get('userGenres') as FormArray;
  }

  addUserGenres() {
    this.userGenres.push(this.fb.control(''));
  }
  
  onSubmit() {
    // TODO: Use EventEmitter with form value
    console.warn(this.playlistForm.value);
    this.submit.emit(this.playlistForm.value);
    alert("New playlist created");
  }

  /*userArray = [];

  storeUser (userArrayOut) {

    this.userArray.push(userArrayOut);

    console.log(this.userArray);
  }*/
}


