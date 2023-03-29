import { Component } from '@angular/core';
import { AccountService } from './account.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'CrowdNFT';
  displayWelcomeMessage: boolean;
  profileLink = 'connect-wallet';
  buttonText = 'Connect';

  constructor(private accountService: AccountService) {
    this.displayWelcomeMessage = true;
  }

  ngOnInit() {
    this.accountService.hasAccount$.subscribe(hasAccount => {
      if (hasAccount) {
        this.profileLink = 'profile';
        this.buttonText = 'View Profile';
      }
      else{
        this.profileLink = 'connect-wallet';
        this.buttonText = 'Connect'
      }
    });
  }

  

  public closeWelcomeMessage(){
    this.displayWelcomeMessage = false;
  }

}

