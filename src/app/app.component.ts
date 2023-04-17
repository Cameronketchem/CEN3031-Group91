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
  searchCategory = "nft" || "user";

  constructor(private router: Router){
    this.displayWelcomeMessage = true;
    this.searchCategory = "nft";
  };

  public closeWelcomeMessage(){
    this.displayWelcomeMessage = false;
  }

  handleSearchClick(event: any){
    console.log("Searching ID " + this.searchQuery)
    if(this.searchQuery.length > 0){
      if(this.searchCategory == "nft"){
        console.log("Searching nft with ID " + this.searchQuery)
        this.router.navigate(['/nftcard/', this.searchQuery]);
      } else if(this.searchCategory == "user"){
        console.log("Searching user with ID " + this.searchQuery)
        this.router.navigate(['/user/', this.searchQuery]);
      }
    }
  }

  onInputChange(event: any){
    this.searchQuery = event.target.value;
  }

  onSelectChange(event: any){
    this.searchCategory = event.target.value;
    console.log("Category is now ")
    console.log(this.searchCategory)
  }
}

