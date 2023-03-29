import { Router } from '@angular/router';
import { Component } from '@angular/core';
import { AccountService } from '../account.service';
import {HttpClient} from '@angular/common/http'

@Component({
  selector: 'app-profile-creation',
  templateUrl: './profile-creation.component.html',
  styleUrls: ['./profile-creation.component.css']
})
export class ProfileCreationComponent {
  Account = {
    userId: '',
    addr: '',
    nonce: '',
    profilePicUrl: '',
    bio: ''
  }

  constructor(private accountService: AccountService, private http: HttpClient, private router:Router){
  }

  onSubmit(){
    this.Account.userId = this.accountService.walletId;
    this.Account.addr = this.accountService.userAddress;
    this.Account.nonce = this.accountService.nonce;
    const accountInfo = {
        profilePic: this.Account.profilePicUrl,
        bio: this.Account.bio
    }
    if(this.Account.profilePicUrl == ''){
      this.Account.profilePicUrl = 'data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBwgHBgkIBwgKCgkLDRYPDQwMDRsUFRAWIB0iIiAdHx8kKDQsJCYxJx8fLT0tMTU3Ojo6Iys/RD84QzQ5OjcBCgoKDQwNGg8PGjclHyU3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3N//AABEIAIEAgQMBIgACEQEDEQH/xAAbAAEAAgMBAQAAAAAAAAAAAAAAAwUBBAYCB//EADIQAQABAwEDBw0BAQAAAAAAAAABAgMEEQUSITEyNEFRcbEVIiNCUlNhYnJzkqHhwRP/xAAWAQEBAQAAAAAAAAAAAAAAAAAAAQL/xAAWEQEBAQAAAAAAAAAAAAAAAAAAEQH/2gAMAwEAAhEDEQA/APogDbIAAAAAAAAAAAAAAM6AMAAAAAACfGxL+TPo6Y3fankb1GxtY8+9GvywVVULS5saqI9HeiZ7JhX3rF2xVu3aJp7J6pSiMBUAAAAAAAAAAGzg4s5V7cnmRxq7mtK82Jb3MWa5jjXP6TVb1FFNFMU0xpERpEQ9MsMqIsmzTftVW644TyfBKyDlr1qqzcqt186JRrTblqIu27sRzo0mVW1moAKgAAAAAAABLotl9Btd0+LnV3sW7FWPNvXjRPJ8JTVWQDKgEgq9ux6C39f+KZZbdu71yi3E82NZVsNYgEioAAAAAAAAJ8PInGvxcjjHJVHbCANV1Vqum5RFdE60zyPbmcXLvYusW5jd9meRYW9tU7sf9LNW98ssxVsgzMmjGs1V1cvqx2yrru2ZmNLNrSe2tXZGRcyLm/dq1nq+BEebtdV27VcqnWap1eQaQAAAAAAAAGY4zpHKt9n7NimIuZFMTV1Uz1G6qvx8PIyONFGlPtTwhu0bGq08+9Ed1K35GWaqp8ix7+fxY8ix7+fx/q3EoqPIse/n8f6eRY9/P4/1bi0Ulex70cy7RV3xMNO/iZFiJm5aqiO2OMOnYmNSpHJi32hs2JibuNTpPrUR19yoaAAQAAA5eECrLY2Nv3Jv1x5tPCnXtXaLGtRYsUW49WP2lYUAAAAAAAAUO1saLF+LlEeZc/Ur5pbXo3sGueumYqXBz4QNMgACTHjeyLVPbXEftGmxOl2fuU+IOnAYaAAAAAAAAGvtDoN/6JbCHNjXEvR8k+AOYAbQAEE2J0qz9ynxAHTgMNAAAAAAAACLJ6Nd+ifAAcvADaAAj//Z';
    }

    if(this.Account.bio == ''){
      this.Account.bio = 'This is my bio!';
    }

    this.http.post("http://localhost:8080/api/users/", this.Account).subscribe(response => {
      console.log(response);
    });
    this.accountService.hasAccount.next(true);
    this.router.navigate(['']);
  }
}
