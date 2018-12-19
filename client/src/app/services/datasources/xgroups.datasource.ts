import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {catchError, finalize} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';

import {XGroup} from '../../models/core';
import {APIXGroupService, XGroupFilter} from '../../services/api/core';
import {ErrorHandlerService} from '../error-handler.service';



export class XGroupsDataSource implements DataSource<XGroup> {

    private xgroupsSubject = new BehaviorSubject<XGroup[]>([]);

    private loadingSubject = new BehaviorSubject<boolean>(false);

    public loading$ = this.loadingSubject.asObservable();

    constructor(private api: APIXGroupService,
                private eh: ErrorHandlerService) {}

    loadXGroups(filter?: XGroupFilter) {
        this.loadingSubject.next(true);
        /*
        this.api.GetXGroups(filter).pipe(
                catchError(() => of([])),
                finalize(() => this.loadingSubject.next(false))
            )
            .subscribe(xgroups => this.xgroupsSubject.next(xgroups));
        */
       this.api.GetXGroups(filter).pipe(
         finalize(() => this.loadingSubject.next(false)),
       ).subscribe(
         xgroups => this.xgroupsSubject.next(xgroups),
         (e) => {
           this.xgroupsSubject.next([]);
           this.eh.HandleError(e);
         }
        );
    }

    connect(collectionViewer: CollectionViewer): Observable<XGroup[]> {
        console.log('Connecting data source');
        return this.xgroupsSubject.asObservable();
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.xgroupsSubject.complete();
        this.loadingSubject.complete();
    }

}

