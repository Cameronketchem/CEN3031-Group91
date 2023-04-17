import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-profile-view',
  templateUrl: './profile-view.component.html',
  styleUrls: ['./profile-view.component.css']
})

export class ProfileViewComponent {

  id: String | null;
  loading: boolean;
  invalidID: boolean;
  userData: {
    user_id: string,
    addr: string,
    profile_pic_url: string,
    bio: string
  };

  
  constructor(private route: ActivatedRoute) {
    // Will set values once we implement call to API
    this.id = null;
    this.loading = false;
    this.invalidID = false;
    this.userData = {
      user_id: '',
      addr: '',
      profile_pic_url: '',
      bio: ''
    };

    this.route.paramMap.subscribe(params => {
      this.id = params.get('id');
  
      this.loading = true;
      fetch(`http://localhost:8080/api/user/${this.id}`)
      .then(res => res.json())
      .then(data => {
        if(!data.error){
          this.invalidID = false;
          console.log("Data is here")
          console.log(data)
          this.userData = {...data};
        } else {
          this.invalidID = true;
        }
        this.loading = false;
      })
      .catch(err => {
        console.log(err)
        this.loading = false;
        this.invalidID = true;
        console.log("here")
      })
    
    })
  }
  
  


}