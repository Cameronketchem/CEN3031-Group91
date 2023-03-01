import { Component, Input } from '@angular/core';

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


}
