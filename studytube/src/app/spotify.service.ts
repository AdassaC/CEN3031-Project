import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';

export interface PlayList {
  playlistTitle: string; 
}

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
    return this.http.post(this.api + "addplaylist/" + playlistName + "/title/" + track.title + "/artist/" + track.artist + "/trackURL/" + track.url, track) 
    .subscribe((res) => {
      console.log(res);
    });
  }

  removeTrackFromPlaylist(track : Track, playlistName: string) {
    return this.http.post(this.api + "removetrack/" + playlistName + "/title/" + track.title + "/artist/" + track.artist, track) 
    .subscribe((res) => {
      console.log(res);
    });
  }

  updateTrackOnPlaylist(track: Track, playlistName: string, newSongName: string, newArtistName: string, newURL: string) {
    return this.http.post(this.api + "updatetrack/" + playlistName + "/title/" + track.title + "/artist/" + track.artist + "/newSong/" + newSongName + "/newArtist/" + newArtistName + "/newURL/" + newURL, track) 
    .subscribe((res) => {
      console.log(res);
    });
  }

  createPlaylist(playlistName: string, userID: string) {
    return this.http.post(this.api + "createPlaylist/" + playlistName + "/" + userID, playlistName)  //<Book>(this.api + "books", book)
    .subscribe((res) => {
      console.log(res);
    });
  }

  getPlaylist(playlistName: string) : Observable<PlayList[]> {
    return this.http.post<PlayList[]>(this.api + "getPlaylist/" + playlistName, playlistName);
  }
}
