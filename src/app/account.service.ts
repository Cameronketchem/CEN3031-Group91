import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Observable } from 'rxjs';
import  Web3  from 'web3';

declare let window: any;
let web3: Web3;

@Injectable({
  providedIn: 'root'
})
export class AccountService {
  web3: any
  userHasAccount = false

  constructor() { }

  public hasAccount = new BehaviorSubject<boolean>(false);
  nonce: any;
  userAddress: any;
  walletId: any;

  get hasAccount$(): Observable<boolean> {
    return this.hasAccount.asObservable();
  }

  public async checkProfileExists() {
    const web3 = new Web3(window.ethereum);
    await window.ethereum.request({ method: 'eth_requestAccounts' });
    const account = await web3.eth.getAccounts();
    this.userAddress = account[0];
    this.nonce = await web3.eth.getTransactionCount(this.userAddress);
    this.walletId = web3.utils.soliditySha3(this.userAddress, this.nonce);
    try {
      fetch('http://localhost:8080/api/users/${walletId}', {
        method: 'GET'
      })
        .then(res => res.json())
        .then(data => {
          console.log(data)
          if (data[0] == this.walletId) {
            this.userHasAccount = true;
          }
          else{
            this.userHasAccount = false;
          }
        })
      const accounts = await window.ethereum.request({ method: 'eth_accounts' });
      this.hasAccount.next(this.userHasAccount);
    } catch (error) {
      console.log(error);
    }
      this.hasAccount.next(true);
  }
}

