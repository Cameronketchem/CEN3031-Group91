import { Router } from '@angular/router';
import { Component } from '@angular/core';
import { AccountService } from '../account.service';
import { HttpClient } from '@angular/common/http'

@Component({
  selector: 'app-profile-creation',
  templateUrl: './profile-creation.component.html',
  styleUrls: ['./profile-creation.component.css']
})

export class ProfileCreationComponent {
  Account = {
    profile_pic_url: '',
    bio: '',
  }

  constructor(private accountService: AccountService, private http: HttpClient, private router:Router){
  }

  async onSubmit(){
    const addr = await this.accountService.address();

    if(this.Account.profile_pic_url == '')
      this.Account.profile_pic_url = 'https://i.seadn.io/gae/AL_lAfiHmqpFk72vHLLTEO8zuRWWYsqRH2uiCld3UuCo8ZAqTah5PwwhtmfFYvlGlTmLh9M7blNUK8kUqzXGN-Lk4hRtUPqAVL9W?auto=format&w=1000';

    if(this.Account.bio == '')
      this.Account.bio = 'This is my bio!';

    // Get signature
    const sig = await this.accountService.signMessage(`ADDR:${addr}`)
    if (sig == '') {
      this.router.navigate([''])
      alert("Invalid signature")
    }

    // Signup
    const xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = () => {
      if (xhttp.readyState === 4) {
        if (xhttp.status !== 200) {
          alert(JSON.parse(xhttp.responseText).error)
          return
        }

        alert("Succesfully signed up")
        this.accountService.hasAccount.next(true);
        this.router.navigate(['']);
      }
    }

    xhttp.open('POST', 'http://localhost:8080/api/auth/signup');
    xhttp.send(JSON.stringify({
      addr: addr,
      profile_pic_url: this.Account.profile_pic_url,
      bio: this.Account.bio,
      sig: sig
    }));
  }
}
