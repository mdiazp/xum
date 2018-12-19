import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpModule } from '@angular/http';
import { FlexLayoutModule } from '@angular/flex-layout';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';

import { DemoMaterialModule, } from '../../demo-material-module';
import { XGroupsRoutes } from './xgroups.routing';
import { ListComponent } from './list/list.component';

@NgModule({
  imports: [
    CommonModule,
    HttpModule,
    FlexLayoutModule,
    RouterModule.forChild(XGroupsRoutes),
    ReactiveFormsModule,
    DemoMaterialModule,
  ],
  declarations: [ListComponent]
})
export class XGroupsModule { }
