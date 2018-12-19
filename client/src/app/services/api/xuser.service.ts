import { Injectable } from '@angular/core';
import {
  Http,
  URLSearchParams,
  Response,
  RequestOptions,
  RequestOptionsArgs,
  Headers,
} from '@angular/http';
import { BehaviorSubject, Observable, Operator } from 'rxjs';
import { map } from 'rxjs/operators';
import { isNullOrUndefined } from 'util';

import {
  XUser,
} from '../../models/core';

import { SessionService } from '../session.service';
import { APIService } from './api.service';
import { Paginator, OrderBy } from './util';

@Injectable()
export class APIXUserService extends APIService {
  constructor(protected http: Http,
              protected session: SessionService) {
    super(http, session);
  }

  public GetXUsers(filter?: XUserFilter): Observable<XUser[]> {
    if ( filter && filter !== null ) {
      return this.get('/xusers', { params: filter.GetUSP() });
    } else {
      return this.get('/xusers');
    }
  }

  public GetCount(filter?: XUserFilter): Observable<number> {
    if ( filter && filter !== null ) {
      return this.get('/xuserscount', { params: filter.GetUSP() });
    } else {
      return this.get('/xuserscount');
    }
  }
}

export class XUserFilter {
  constructor(public NameSubstr: string,
              public LastNameSubstr: string,
              public PhoneNumberSubstr: string,
              public IdentificationPrefix: string,
              public Actived: boolean,
              public XGroupID: number,
              public paginator: Paginator,
              public orderby: OrderBy) {}

  public GetUSP(): URLSearchParams {
    let usp: URLSearchParams;
    usp = new URLSearchParams();
    if (  this.NameSubstr !== null && this.NameSubstr !== '' ) {
      usp.append('nameSubstr', this.NameSubstr.toString());
    }
    if ( this.LastNameSubstr !== null && this.LastNameSubstr !== '' ) {
      usp.append('lastnameSubstr', this.LastNameSubstr.toString());
    }
    if ( this.PhoneNumberSubstr !== null && this.PhoneNumberSubstr !== '' ) {
      usp.append('phoneSubstr', this.PhoneNumberSubstr.toString());
    }
    if ( this.IdentificationPrefix !== null && this.IdentificationPrefix !== '' ) {
      usp.append('identificationPrefix', this.IdentificationPrefix.toString());
    }
    if ( this.Actived !== null ) {
      usp.append('actived', this.Actived.toString());
    }
    if ( !isNullOrUndefined(this.XGroupID) && this.XGroupID !== 0 ) {
      usp.append('xgroupid', this.XGroupID.toString());
    }
    if ( this.paginator !== null ) {
      usp.appendAll(this.paginator.GetUSP());
    }
    if ( this.orderby !== null ) {
      usp.appendAll(this.orderby.GetUSP());
    }
    return usp;
  }
}
