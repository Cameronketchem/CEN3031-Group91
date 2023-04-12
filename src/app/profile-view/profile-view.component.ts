import { Component } from '@angular/core';

@Component({
  selector: 'app-profile-view',
  templateUrl: './profile-view.component.html',
  styleUrls: ['./profile-view.component.css']
})

export class ProfileViewComponent {

  userData: {
    user_id: string,
    wallet_pubkey: string,
    public_pic_url: string,
    bio: string
  };
  
  constructor() {
    // Will set values once we implement call to API
    this.userData = {
      user_id: 'Username',
      wallet_pubkey: '',
      public_pic_url: '',
      bio: ''
    };
  }

}
