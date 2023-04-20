import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';

import { SignInComponent } from './components/sign-in/sign-in.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { DashboardComponent } from './components/dashboardComponents/dashboard/dashboard.component';
import { ForgotPasswordComponent } from './components/forgot-password/forgot-password.component';
import { VerifyEmailComponent } from './components/verify-email/verify-email.component';
// route guard
import { AuthGuard } from './shared/guard/auth.guard';
import { SpotifyPlayerComponent } from './components/spotify-player/spotify-player.component';
import { HomeComponent } from './pages/home/home.component';
import { AboutComponent } from './pages/about/about.component';
import { ReportBugsComponent } from './pages/report-bugs/report-bugs.component';
import { NavbarComponent } from './shared/navbar/navbar.component';
import { FooterComponent } from './shared/footer/footer.component';
import { TaskViewComponent } from './tasksManager/task-view/task-view.component';
import { NewListComponent } from './tasksManager/new-list/new-list.component';
import { PlaylistGeneratorComponent } from './components/dashboardComponents/playlist-generator/playlist-generator.component';
import { SettingsComponent } from './components/dashboardComponents/settings/settings.component';
import { ListsComponent } from './components/dashboardComponents/lists/lists.component';
import { TodoComponent } from './components/todo/todo.component';
import { StripeComponent } from './components/dashboardComponents/stripeSubcription/stripe/stripe.component';
import { SubscribeComponent } from './components/dashboardComponents/stripeSubcription/subscribe/subscribe.component';
import { UpdatePlaylistComponent } from './components/dashboardComponents/update-playlist/update-playlist.component';
import { CreateListComponent } from './components/dashboardComponents/create-list/create-list.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'sign-in', component: SignInComponent },
  { path: 'register-user', component: SignUpComponent },
  { path: 'dashboard', component: DashboardComponent, canActivate: [AuthGuard] },
  { path: 'forgot-password', component: ForgotPasswordComponent },
  { path: 'verify-email-address', component: VerifyEmailComponent },

  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  
  { path: 'player', component: SpotifyPlayerComponent},
  { path: 'home', component: HomeComponent},
  { path: 'about', component: AboutComponent},
  { path: 'report-bugs', component: ReportBugsComponent},
  { path: 'playlist', component: PlaylistGeneratorComponent},

  { path: 'view-tasks', component: TaskViewComponent},
  { path: 'lists/:listId', component: TaskViewComponent},
  { path: 'new-list', component: NewListComponent},

  { path: 'navbar', component: NavbarComponent},

  { path: 'footer', component: FooterComponent},

  { path: 'lists', component: ListsComponent},
  { path: 'todo', component: TodoComponent},
  { path: 'settings', component: SettingsComponent},
  { path: 'stripe', component: StripeComponent},
  { path: 'subscribe', component: SubscribeComponent},
  { path: 'update-playlist', component: UpdatePlaylistComponent},
  { path: 'create-list', component: CreateListComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
