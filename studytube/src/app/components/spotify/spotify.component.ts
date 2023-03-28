import { Component } from '@angular/core';
import { Track, SpotifyService, PlayList } from 'src/app/spotify.service';


@Component({
  selector: 'app-spotify',
  templateUrl: './spotify.component.html',
  styleUrls: ['./spotify.component.css']
})
export class SpotifyComponent {
    public tracks: Track[] = []; 
    public playlist: PlayList[] = []; 

    constructor(
      private spotifyService: SpotifyService, 
    ) {
      //this.tracks = bookService.get();
     }

     addTrackPlaylist(songName: string, artistName: string, urlName: string, playlistName: string) {
        this.spotifyService.addTrackToPlaylist({
              title: songName,
              artist: artistName, 
              url: urlName,
        }, playlistName)
     }

     removeTrackPlaylist(songName: string, artistName: string, urlName: string, playlistName: string) {
        this.spotifyService.removeTrackFromPlaylist({
          title: songName,
          artist: artistName, 
          url: urlName,
        }, playlistName)
     }

     updateTrackPlaylist(songName: string, artistName: string, urlName: string, playlistName: string, updatedSongName: string, updatedArtistName: string, updatedURL: string) {
        this.spotifyService.updateTrackOnPlaylist({
          title: songName,
          artist: artistName, 
          url: urlName,
        }, playlistName,
        updatedSongName, updatedArtistName, updatedURL)
     }

     createPlaylist(playlistName: string) {
      this.spotifyService.createPlaylist(playlistName);
     }

     getPlaylist(playlistName: string) {
        this.spotifyService.getPlaylist(playlistName).subscribe(res => {
          this.playlist = res; 
        });
     }

}
