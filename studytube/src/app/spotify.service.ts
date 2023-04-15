import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface PlayList {
  playlistTitle: string; 
}

export interface Track {
  title: string; 
  artist: string; 
  url: string; 
}

class track {
  title: string | undefined; 
  artist: string | undefined; 
  url: string | undefined; 
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

  updateTrackOnPlaylist(track: Track, playlistName: string, newSongName: string, newArtistName: string, newURL: string) {
    /*headers = new Headers({'Content-Type': 'application/json'});
    theTrack = track {
      title: newSongName,
      artist: newArtistName,
      url: newURL,
    }
    return this.http.put(ap) */ 
  }

  createPlaylist(playlistName: string) {
    return this.http.post(this.api + "books", playlistName)  //<Book>(this.api + "books", book)
    .subscribe((res) => {
      console.log(res);
    });
  }

  getPlaylist(playlistName: string) { 
    return this.http.get(this.api = "playlistName")
  }
}