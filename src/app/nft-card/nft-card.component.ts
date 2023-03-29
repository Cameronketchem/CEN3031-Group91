import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';


@Component({
  selector: 'app-nft-card',
  templateUrl: './nft-card.component.html',
  styleUrls: ['./nft-card.component.css']
})
export class NftCardComponent {

  
  @Input() id: Number | undefined;
  @Input() description: String | undefined;
  @Input() price: Number | undefined;
  @Input() image: String | undefined;

  constructor(private router: Router){

  }

  visitNFT(id: Number | undefined){
    console.log("Visit nft with id " + id);
    if(id){
      this.router.navigateByUrl(`/nftcard/${id}`)
    }
  }

}
