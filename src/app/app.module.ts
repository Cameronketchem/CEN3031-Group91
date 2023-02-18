import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { NftCardComponent } from './nft-card/nft-card.component';
import { ConnectWalletComponent } from './connect-wallet/connect-wallet.component';

const routes: Routes = [ {path: '', component: HomeComponent}, 
  { path: 'connect-wallet', component: ConnectWalletComponent }];

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    NftCardComponent,
    ConnectWalletComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    RouterModule,
    [RouterModule.forRoot(routes)]
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
