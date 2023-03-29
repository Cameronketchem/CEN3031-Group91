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
  showFeed: boolean;
  dataOffSet: number;
  loading: boolean;


  constructor(){
    this.NFTs = new Array();
    this.showFeed = true;
    this.dataOffSet = -1;
    this.loading = false;
  }
  ngOnInit(): void {
    fetch(`http://localhost:8080/api/assets/${this.dataOffSet}`, {
      method: 'GET'
    })
      .then(res => res.json())
      .then(data => {
        console.log(data)
        this.NFTs = [...data]
      })
      .catch(err => console.log(err))
   
  }

  loadMoreData(){
    console.log("loading more data!")
    this.loading = true;
    this.dataOffSet += 20;
    fetch(`http://localhost:8080/api/assets/${this.dataOffSet}`, {
      method: 'GET'
    })
      .then(res => res.json())
      .then(data => {
        console.log(data)
        this.NFTs = [...this.NFTs, ...data];
        this.loading = false;
      })
      .catch(err => {
        console.log(err);
        this.loading = false;
      })
  }


  

}
