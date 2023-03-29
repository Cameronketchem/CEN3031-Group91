import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { NftCardComponent } from './nft-card/nft-card.component';
import { ConnectWalletComponent } from './connect-wallet/connect-wallet.component';
import { ProfileCreationComponent } from './profile-creation/profile-creation.component';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';


const routes: Routes = [ {path: '', component: HomeComponent}, 
  { path: 'connect-wallet', component: ConnectWalletComponent },
{ path: 'profile-creation', component: ProfileCreationComponent}];

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    NftCardComponent,
    ConnectWalletComponent,
    ProfileCreationComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    RouterModule,
    [RouterModule.forRoot(routes)]
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
