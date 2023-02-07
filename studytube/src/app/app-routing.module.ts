import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

//Import components and routes to them
import { AboutComponent } from './pages/about/about.component';
import { ReportBugsComponent } from './pages/report-bugs/report-bugs.component';
import { HomeComponent } from './pages/home/home.component';

//Set routes using components
const routes: Routes = [
  {path:'',component:HomeComponent},
  {path:'about',component:AboutComponent},
  {path:'report-bugs',component:ReportBugsComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
