import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-user-card',
  templateUrl: './user-card.component.html',
  styleUrls: ['./user-card.component.css']
})
export class UserCardComponent {

  @Input() id: Number | undefined;
  @Input() bio: String | undefined;
  @Input() addr: String | undefined;
  @Input() profile_pic_url: String | undefined;

  constructor(private router: Router){

  }

  visitUser(id: Number | undefined){
    console.log("Visit user with id " + id);
    if(id){
      this.router.navigateByUrl(`/user/${id}`)
    }
  }

}


