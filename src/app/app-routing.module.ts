import { HomeComponent } from './home/home.component';
import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { RouterModule, Routes } from "@angular/router";
import { ConnectWalletComponent } from './connect-wallet/connect-wallet.component';
import { NftCardPageComponent } from './nft-card-page/nft-card-page.component';


const routes: Routes = [  {path:'', component: HomeComponent},
{ path: 'connect-wallet', component: ConnectWalletComponent },
{ path: 'nftcard/:id', component: NftCardPageComponent }]

@NgModule({
  imports: [RouterModule.forRoot(routes), RouterModule],
  exports: [RouterModule]
})
export class AppRoutingModule {

 }
