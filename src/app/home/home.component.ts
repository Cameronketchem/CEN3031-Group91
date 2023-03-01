import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  NFTs: Array<{
    asset_id: Number,
    asset_contract: Number,
    desc: String,
    exec_contract: Number,
    executor: Number,
    img_preview: String,
    price: Number
  }>;

  constructor(){
    this.NFTs = new Array();
  }
  ngOnInit(): void {
    fetch('http://localhost:8080/api/assets/-1', {
      method: 'GET'
    })
      .then(res => res.json())
      .then(data => {
        console.log(data)
        for(let i = 0; i < data.length; i++){
          this.NFTs[i] = {...data[i]};
        }
      })
      .catch(err => console.log(err))
   
  }


  

}
