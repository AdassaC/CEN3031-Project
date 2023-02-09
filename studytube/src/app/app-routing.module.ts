import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

//Import components and routes to them
import { AboutComponent } from './pages/about/about.component';
import { ReportBugsComponent } from './pages/report-bugs/report-bugs.component';
import { HomeComponent } from './pages/home/home.component';
//import { DashboardComponent } from './pages/dashboard/dashboard.component';

//Set routes using components
const routes: Routes = [
  {path:'',component:HomeComponent},
  {path:'about',component:AboutComponent},
  {path:'report-bugs',component:ReportBugsComponent},
  //Gonna have to change to include user id in path
  //{path: 'dashboard', component:DashboardComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
