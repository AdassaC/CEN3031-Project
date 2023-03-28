import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';

import { SignInComponent } from './components/sign-in/sign-in.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
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
import { SettingsComponent } from './components/settings/settings.component';
import { ListsComponent } from './components/lists/lists.component';
import { TodoComponent } from './components/todo/todo.component';

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
  { path: 'view-tasks', component: TaskViewComponent},

  { path: 'navbar', component: NavbarComponent},
  { path: 'footer', component: FooterComponent},

  { path: 'lists', component: ListsComponent},
  { path: 'todo', component: TodoComponent},
  { path: 'settings', component: SettingsComponent}
  //Gonna have to change to include user id in path
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
