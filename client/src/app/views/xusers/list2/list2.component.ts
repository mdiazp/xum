import { Component, ViewChild, OnInit, AfterViewInit } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import {
  APIXUserService,
  XUserFilter,
  XUsersDataSource,
  ErrorHandlerService,
  Paginator,
  OrderBy,
} from '../../../services/core';
import { tap, merge } from 'rxjs/operators';

@Component({
  selector: 'app-list2',
  templateUrl: './list2.component.html',
  styleUrls: ['./list2.component.scss']
})
export class List2Component implements OnInit, AfterViewInit {

  dataSource: XUsersDataSource;
  displayedColumns= ['xuser'];

  @ViewChild(MatPaginator) paginator: MatPaginator;
  initialPageSize = 20;
  pageSizeOptions = [20, 50, 100];

  constructor(private api: APIXUserService,
              private eh: ErrorHandlerService) {}

  ngOnInit() {
      this.dataSource = new XUsersDataSource(this.api, this.eh);
      this.dataSource.load(
        true,
        new XUserFilter(
          null,
          null,
          null,
          null,
          null,
          null,
          new Paginator(
            0,
            this.initialPageSize,
          ),
          null,
        ),
      );
  }

  ngAfterViewInit() {
    this.paginator.page
      .pipe(
        tap(() => this.load(false)),
      )
      .subscribe();
  }

  load(loadCount: boolean) {
    console.log('xusers load');

    this.dataSource.load(
      loadCount,
      new XUserFilter(
        null,
        null,
        null,
        null,
        null,
        null,
        new Paginator(
          this.paginator.pageIndex * this.paginator.pageSize,
          this.paginator.pageSize
        ),
        new OrderBy(
          'id',
          false,
        ),
      ),
    );
  }
}
