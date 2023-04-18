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
  users: Array<{
    user_id: Number,
    addr: String,
    bio: String,
    nonce: String,
    profile_pic_url: String
  }>;
  showFeed: boolean;
  NFTsOffSetData: number;
  usersOffSetData: number;
  loading: boolean;
  noMoreData: boolean;
  feedCategory: string;


  constructor(){
    this.NFTs = new Array();
    this.users = new Array();
    this.showFeed = true;
    this.NFTsOffSetData = -1;
    this.usersOffSetData = -1;
    this.loading = false;
    this.noMoreData = false;
    this.feedCategory = "nfts";
  }
  ngOnInit(): void {
    fetch(`http://localhost:8080/api/assets/${this.NFTsOffSetData}`, {
      method: 'GET'
    })
      .then(res => res.json())
      .then(data => {
        console.log(data)
        this.NFTs = [...data]
      })
      .catch(err => console.log(err))

    fetch(`http://localhost:8080/api/users/${this.usersOffSetData}`, {
      method: 'GET'
    })
      .then(res => res.json())
      .then(data => {
        console.log(data)
        this.users = [...data]
      })
      .catch(err => console.log(err))
   
  }

  loadMoreData(){
    console.log("loading more data!")
    this.loading = true;
    if(this.feedCategory == 'nfts'){
      this.NFTsOffSetData += 20;
      fetch(`http://localhost:8080/api/assets/${this.NFTsOffSetData}`, {
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
        // If we receive error then most probably there are no more NFTs left
        this.noMoreData = true;
        this.loading = false;
      })
    } else if(this.feedCategory == 'users') {
      fetch(`http://localhost:8080/api/users/${this.usersOffSetData}`, {
        method: 'GET'
      })
      .then(res => res.json())
      .then(data => {
        console.log(data)
        this.users = [...this.users, ...data];
        this.loading = false;
        if(data.length != 0){
          this.usersOffSetData += 20;
        }
      })
      .catch(err => {
        console.log(err);
        // If we receive error then most probably there are no more NFTs left
        this.noMoreData = true;
        this.loading = false;
      })
    }
  }

  onSortByChange(e: any){
    let sortValue = e.target.value;
    // Will sort NFTs array by asset_id or price
    if(sortValue == "id-htl"){
      // Sort by asset_id from high to low
      this.NFTs.sort((a: any, b: any) => b.asset_id - a.asset_id);
    } else if(sortValue == "id-lth"){
      // Sort by asset_id from low to high
      this.NFTs.sort((a: any, b: any) => a.asset_id - b.asset_id);
    } else if(sortValue == "price-htl"){
      // Sort by price from high to low
      this.NFTs.sort((a: any, b: any) => b.price - a.price);
    } else if(sortValue == "price-lth"){
      // Sort by price from low to high
      this.NFTs.sort((a: any, b: any) => a.price - b.price);
    }

  }

  toggleSelected(category: string) {
    if(this.feedCategory != category){

    }
    this.feedCategory = category;
    console.log("Changed to " + category)
  }


  

}
