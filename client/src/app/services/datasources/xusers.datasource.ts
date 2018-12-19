import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {catchError, finalize} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';

import {XUser} from '../../models/core';
import {APIXUserService, XUserFilter} from '../api/core';
import {ErrorHandlerService} from '../error-handler.service';



export class XUsersDataSource implements DataSource<XUser> {

    private xusersSubject = new BehaviorSubject<XUser[]>([]);
    private countSubject = new BehaviorSubject<number>(0);

    private loadingSubject = new BehaviorSubject<boolean>(false);

    public loading$ = this.loadingSubject.asObservable();

    public count$ = this.countSubject.asObservable();

    constructor(private api: APIXUserService,
                private eh: ErrorHandlerService) {}

    load(loadCount: boolean, filter?: XUserFilter) {
      this.loadingSubject.next(true);
      if ( loadCount ) {
        this.loadCount(filter);
      } else {
        this.loadXUsers(filter);
      }
    }

    private loadCount(filter?: XUserFilter) {
      this.api.GetCount(filter).subscribe(
        count => {
          this.countSubject.next(count);
          this.loadXUsers(filter);
        },
        (e) => {
          this.countSubject.next(0);
          this.xusersSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    private loadXUsers(filter?: XUserFilter) {
      this.api.GetXUsers(filter).subscribe(
        xusers => this.xusersSubject.next(xusers),
        (e) => {
          this.xusersSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    connect(collectionViewer: CollectionViewer): Observable<XUser[]> {
        console.log('Connecting data source');
        return this.xusersSubject.asObservable();
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.xusersSubject.complete();
        this.loadingSubject.complete();
    }

}

