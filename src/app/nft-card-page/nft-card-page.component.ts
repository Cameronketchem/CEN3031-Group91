import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-nft-card-page',
  templateUrl: './nft-card-page.component.html',
  styleUrls: ['./nft-card-page.component.css']
})
export class NftCardPageComponent {
  
  id: string | null;
  invalidID: Boolean;
  loading: Boolean;

  NFTData: {
    asset_id: Number,
    asset_contract: {
      addr: String,
      contract_id: Number,
      is_exec: Boolean,
      is_finished: Boolean
    },
    desc: String,
    exec_contract: {
      addr: String,
      contract_id: Number,
      is_exec: Boolean,
      is_finished: Boolean
    },
    executor: {
      addr: String,
      bio: String,
      nonce: String,
      profile_pic_url: String,
      user_id: Number
    },
    img_preview: String,
    price: Number
  } | null;

  constructor(private route: ActivatedRoute){
    
    this.id = null;
    this.NFTData = null;
    this.invalidID = false;
    this.loading = false;

    this.route.paramMap.subscribe(params => {
      this.id = params.get('id');

      this.loading = true;
      fetch(`http://localhost:8080/api/asset/${this.id}`)
      .then(res => res.json())
      .then(data => {
        if(!data.error){
          this.invalidID = false;
          console.log("Data is here")
          console.log(data)
          this.NFTData = {...data};
        } else {
          this.invalidID = true;
        }
        this.loading = false;
      })
      .catch(err => {
        console.log(err)
        this.loading = false;
      })


    })
  }

}
