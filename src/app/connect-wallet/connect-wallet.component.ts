import { AccountService } from './../account.service';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import Web3 from 'web3';

declare let window: any;
let web3: Web3;

@Component({
  selector: 'app-connect-wallet',
  templateUrl: './connect-wallet.component.html',
  styleUrls: ['./connect-wallet.component.css']
})


export class ConnectWalletComponent {
  web3: any
  buttonText = 'Connect'
  public metaMaskInstalled: boolean = false;

  constructor(private accountService: AccountService, private router: Router) { }

  public async login() {
    if (typeof window.ethereum !== 'undefined') {
    try {
      await window.ethereum.enable();
      this.web3 = new Web3(window.ethereum);
      this.metaMaskInstalled = true;
      console.log('Web3 instance:', this.web3);
    } catch (error) {
      console.error(error);
    }
    this.accountService.checkProfileExists();
    this.accountService.hasAccount$.subscribe(hasAccount => {
      if (!hasAccount) {
        this.router.navigate(['profile-creation']);
        this.buttonText = 'Connect';
      }
      else{
        this.buttonText = 'View Profile';
        this.router.navigate(['']);
      }
    });
  }
  else{
    console.log("User does not have Meta Mask extension");
    alert('You don\'t have the Metamask extension installed!  Please visit "https://metamask.io/" to create an account');
  }
  }
    
  
  async readAccount() {
    const accounts = await this.web3.eth.getAccounts();
    console.log('Accounts:', accounts);
  }
}