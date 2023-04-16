import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule } from '@angular/forms';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatChipsModule} from '@angular/material/chips';


import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { RegisterComponent } from './components/register/register.component';
import { LoginComponent } from './components/login/login.component';

// Firebase services + environment module
import { AngularFireModule } from '@angular/fire/compat';
import { AngularFireAuthModule } from '@angular/fire/compat/auth';
import { AngularFireStorageModule } from '@angular/fire/compat/storage';
import { AngularFirestoreModule } from '@angular/fire/compat/firestore';
import { AngularFireDatabaseModule } from '@angular/fire/compat/database';
import { environment } from '../environments/environment';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { SignInComponent } from './components/sign-in/sign-in.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { ForgotPasswordComponent } from './components/forgot-password/forgot-password.component';
import { VerifyEmailComponent } from './components/verify-email/verify-email.component';
import { AuthService } from './shared/services/auth';

import { AboutComponent } from './pages/about/about.component';
import { HomeComponent } from './pages/home/home.component';
import { ReportBugsComponent } from './pages/report-bugs/report-bugs.component';
import { FormsModule} from '@angular/forms';
import { FooterComponent } from './shared/footer/footer.component';
import { NavbarComponent } from './shared/navbar/navbar.component';
//import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { SpotifyPlayerComponent } from './components/spotify-player/spotify-player.component';
import { HttpClientModule } from '@angular/common/http';
import { SpotifyComponent } from './components/spotify/spotify.component';
import { StripeComponent } from './components/stripe/stripe.component';
import { TaskViewComponent } from './tasksManager/task-view/task-view.component';
import { NewListComponent } from './tasksManager/new-list/new-list.component';
import { TodoComponent } from './components/todo/todo.component';
import { PlaylistGeneratorComponent } from './components/playlist-generator/playlist-generator.component';
import { ListsComponent } from './components/lists/lists.component';
import { SettingsComponent } from './components/settings/settings.component';


@NgModule({
  declarations: [
    AppComponent,
    DashboardComponent,
    SignInComponent,
    SignUpComponent,
    ForgotPasswordComponent,
    VerifyEmailComponent,
    RegisterComponent,
    LoginComponent,
    AboutComponent,
    HomeComponent,
    ReportBugsComponent,
    FooterComponent,
    NavbarComponent,
    SpotifyPlayerComponent,
    SpotifyComponent,
    StripeComponent,
    TaskViewComponent,
    NewListComponent,
    TodoComponent,
    PlaylistGeneratorComponent,
    ListsComponent,
    SettingsComponent
  ],

  imports: [
    BrowserModule,
    AppRoutingModule,
    AngularFireModule.initializeApp(environment.firebase),
    AngularFireAuthModule,
    AngularFirestoreModule,
    AngularFireStorageModule,
    AngularFireDatabaseModule,
   // BrowserAnimationsModule,
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatChipsModule
  ],

  providers: [AuthService],

  bootstrap: [AppComponent]
})
export class AppModule { }
