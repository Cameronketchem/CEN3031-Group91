import { Component, Input } from '@angular/core';
import { SignupComponent } from '../signup/signup.component';
import { LoginComponent } from '../login/login.component';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {

  rando = Math.random() * 100;

  @Input() value: 'X' | 'Y' | undefined;

  constructor(){};

  public handleClick(){
    console.log("Clickeamos!");
  }

}
