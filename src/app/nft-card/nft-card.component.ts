import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-nft-card',
  templateUrl: './nft-card.component.html',
  styleUrls: ['./nft-card.component.css']
})
export class NftCardComponent {


  randomNumber = Math.random() * 100;

  @Input() name: String | undefined;
  @Input() price: Number | undefined;




}
