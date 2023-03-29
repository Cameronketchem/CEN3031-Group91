import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'CrowdNFT';
  displayWelcomeMessage: boolean;

  constructor(){
    this.displayWelcomeMessage = true;
  };

  public closeWelcomeMessage(){
    this.displayWelcomeMessage = false;
  }
}

