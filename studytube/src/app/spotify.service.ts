import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Track {
  title: string; 
  artist: string; 
  url: string; 
}

@Injectable({
  providedIn: 'root'
})
export class SpotifyService {

  api: string = 'http://localhost:4201/';
  constructor(private http: HttpClient) {}

  addTrackToPlaylist(track : Track, playlistName: string) {
    return this.http.post(this.api + "books", track)  //<Book>(this.api + "books", book)
    .subscribe((res) => {
      console.log(res);
    });
  }

  removeTrackFromPlaylist(track : Track, playlistName: string) {

  }

  updateTrackOnPlaylist(track: Track, playlistName: string, newSongName: string, newArtistName: string, updatedURL: string) {

  }

  createPlaylist(playlistName: string) {
    return this.http.post(this.api + "books", track)  //<Book>(this.api + "books", book)
    .subscribe((res) => {
      console.log(res);
    });
  }

  getPlaylist(playlistName: string) { 

  }
}
