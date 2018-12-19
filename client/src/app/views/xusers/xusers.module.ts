import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpModule } from '@angular/http';
import { FlexLayoutModule } from '@angular/flex-layout';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';

import { DemoMaterialModule, } from '../../demo-material-module';
import { XUsersRoutes } from './xusers.routing';
import { ListComponent } from './list/list.component';
import { List2Component } from './list2/list2.component';

@NgModule({
  imports: [
    CommonModule,
    HttpModule,
    FlexLayoutModule,
    RouterModule.forChild(XUsersRoutes),
    ReactiveFormsModule,
    DemoMaterialModule,
  ],
  declarations: [ListComponent, List2Component]
})
export class XUsersModule { }
