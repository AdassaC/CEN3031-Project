import { Component, OnInit, Pipe, PipeTransform } from '@angular/core';
import { Input, Output } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
// Importing Iframely's embed.js as library.
//import { iframely } from '@iframely/embed.js';

@Component({
  selector: 'app-spotify-player',
  templateUrl: './spotify-player.component.html',
  styleUrls: ['./spotify-player.component.css']
})

export class SpotifyPlayerComponent {
  public srcURL: string = "https://open.spotify.com/embed/playlist/37i9dQZF1DX8NTLI2TtZa6?utm_source=generator&theme=0";
}

/*getSrcURL() {
  this.srcURL = "https://open.spotify.com/embed/playlist/37i9dQZF1DX8NTLI2TtZa6?utm_source=generator&theme=0";
}*/

/*@Component({
  selector: 'app-spotify-player',
  templateUrl: './spotify-player.component.html',
  styleUrls: ['./spotify-player.component.css'],
  selector: 'app-iframely-embed',
  template: `<div>
    <p bind-innerHTML="htmlCode | safeHtml"></p>
  </div>`,
})
/*export class SpotifyPlayerComponent {
  /** Get html via JSON API calls to /api/oembed or /api/iframely. 
  htmlCode = '<script>
    window.onSpotifyIframeApiReady = (IFrameAPI) => {
        let element = document.getElementById('embed-iframe');
        let options = {
            uri: 'spotify:episode:7makk4oTQel546B0PZlDM5'
          };
        let callback = (EmbedController) => {};
        IFrameAPI.createController(element, options, callback);
      };
  </script>';

  /*constructor() {
    / Trigger on data load from source in case html has embed.js. 
    iframely.load();
  }

  ngOnInit() {}*/
  /*
  window.onSpotifyIframeApiReady = (IFrameAPI) => {
    let element = document.getElementById('embed-iframe');
    let options = {
      uri: 'spotify:episode:7makk4oTQel546B0PZlDM5'
    };
    let callback = (EmbedController) => {};
    IFrameAPI.createController(element, options, callback);
  };

}*/
