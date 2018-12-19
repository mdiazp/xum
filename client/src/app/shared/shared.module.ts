import { NgModule } from '@angular/core';

import { MenuItems } from './menu-items/menu-items';
import {
  AccordionAnchorDirective,
  AccordionLinkDirective,
  AccordionDirective
} from './accordion';
import {
  AuthGuard
} from '../guards/core';
import {
  LocalStorageService,
  SessionService,
  ErrorHandlerService,
  APIAccountService,
  APIXGroupService,
} from '../services/core';
import { APIXUserService } from '../services/api/xuser.service';
// import { APIAccountService } from '../services/api/account.services';


@NgModule({
  declarations: [
    AccordionAnchorDirective,
    AccordionLinkDirective,
    AccordionDirective
  ],
  exports: [
    AccordionAnchorDirective,
    AccordionLinkDirective,
    AccordionDirective
   ],
  providers: [
    MenuItems,

    AuthGuard,

    LocalStorageService,
    SessionService,
    ErrorHandlerService,

    APIAccountService,
    APIXGroupService,
    APIXUserService,
  ]
})
export class SharedModule { }
