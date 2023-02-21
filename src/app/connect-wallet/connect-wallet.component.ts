import { Component } from '@angular/core';
import Web3 from 'web3';
import { AbiItem } from 'web3-utils';

declare let window: any;
let web3: Web3;

@Component({
  selector: 'app-connect-wallet',
  templateUrl: './connect-wallet.component.html',
  styleUrls: ['./connect-wallet.component.css']
})


export class ConnectWalletComponent {
  web3: any

  public async login() {
    try {
      await window.ethereum.enable();
      this.web3 = new Web3(window.ethereum);
      console.log('Web3 instance:', this.web3);
      // Perform other actions with the web3 instance as needed
    } catch (error) {
      console.error(error);
    }
  }
    
  
  async readAccount() {
    const accounts = await this.web3.eth.getAccounts();
    console.log('Accounts:', accounts);
  }
}