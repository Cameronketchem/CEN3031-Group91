import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {

  randomNumbers: Array<Number>;

  constructor(){
    this.randomNumbers = new Array(20);
    for(let i = 0; i < this.randomNumbers.length; i++){
      this.randomNumbers[i] = Math.random() * 100;
    }
  }
  

}
