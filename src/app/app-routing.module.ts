import { HomeComponent } from './home/home.component';
import { SignupComponent } from './signup/signup.component';
import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { RouterModule, Routes } from "@angular/router";
import { LoginComponent } from './login/login.component';

const routes: Routes = [  {path:'', component: HomeComponent},
{ path: 'signup', component: SignupComponent },
{path: 'login', component: LoginComponent},];

@NgModule({
  imports: [RouterModule.forRoot(routes), RouterModule],
  exports: [RouterModule]
})
export class AppRoutingModule { }
