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

import {
  XGroup,
} from '../../models/core';

import { SessionService } from '../session.service';
import { APIService } from './api.service';
import { Paginator, OrderBy } from './util';

@Injectable()
export class APIXGroupService extends APIService {
  constructor(protected http: Http,
              protected session: SessionService) {
    super(http, session);
  }

  public GetXGroups(filter?: XGroupFilter): Observable<XGroup[]> {
    if ( filter && filter !== null ) {
      return this.get('/xgroups', { params: filter.GetUSP() });
    } else {
      return this.get('/xgroups');
    }
  }
}

export class XGroupFilter {
  constructor(public NameSubstr: string,
              public AdminID: number,
              public Actived: boolean,
              public paginator: Paginator,
              public orderby: OrderBy) {}

  public GetUSP(): URLSearchParams {
    let usp: URLSearchParams;
    usp = new URLSearchParams();
    if ( this.NameSubstr !== null && this.NameSubstr !== '' ) {
      usp.append('nameSubstr', this.NameSubstr.toString());
    }
    if ( this.Actived !== null ) {
      usp.append('actived', this.Actived.toString());
    }
    if ( this.AdminID !== null && this.AdminID !== 0 ) {
      usp.append('adminID', this.NameSubstr.toString());
    }
    usp.appendAll(this.paginator.GetUSP());
    usp.appendAll(this.orderby.GetUSP());
    return usp;
  }
}
