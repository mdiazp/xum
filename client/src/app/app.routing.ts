import { Routes } from '@angular/router';

import { FullComponent } from './layouts/full/full.component';
import { AuthGuard } from './guards/core';
// import { LoginComponent } from './login/login.component';

export const AppRoutes: Routes = [
  {
    path: 'login',
    loadChildren: './views/login/login.module#LoginModule'
  },
  {
    path: '',
    component: FullComponent,
    canActivate: [AuthGuard],
    children: [
      {
        path: '',
        redirectTo: '/starter',
        pathMatch: 'full'
      },
      {
        path: '',
        loadChildren:
          './material-component/material.module#MaterialComponentsModule'
      },
      {
        path: 'starter',
        loadChildren: './starter/starter.module#StarterModule'
      },
      {
        path: 'users',
        loadChildren: './views/users/users.module#UsersModule'
      },
      {
        path: 'xgroups',
        loadChildren: './views/xgroups/xgroups.module#XGroupsModule'
      },
      {
        path: 'xusers',
        loadChildren: './views/xusers/xusers.module#XUsersModule'
      },
    ]
  }
];
