import { Component } from '@angular/core';
import { Router } from '@angular/router';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'CrowdNFT';
  displayWelcomeMessage: boolean;
  searchQuery = "";

  constructor(private router: Router){
    this.displayWelcomeMessage = true;
  };

  public closeWelcomeMessage(){
    this.displayWelcomeMessage = false;
  }

  handleSearchClick(event: any){
    console.log("Searching ID " + this.searchQuery)
    if(this.searchQuery.length > 0){
      this.router.navigate(['/nftcard/', this.searchQuery]);
    }
  }

  onInputChange(event: any){
    this.searchQuery = event.target.value;
  }
}

