import { Component } from '@angular/core';
import { Track, SpotifyService } from 'src/app/spotify.service';


@Component({
  selector: 'app-spotify',
  templateUrl: './spotify.component.html',
  styleUrls: ['./spotify.component.css']
})
export class SpotifyComponent {
    public tracks: Track[] = []; 

    constructor(
      private spotifyService: SpotifyService, 
    ) {
      //this.tracks = bookService.get();
     }

     addTrackPlaylist() {

     }

     removeTrackPlaylist() {

     }

     updateTrackPlaylist() {

     }

     createPlaylist() {

     }

     getPlaylist() {
      
     }

}
