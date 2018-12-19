import { Component, ViewChild, OnInit, AfterViewInit, ElementRef } from '@angular/core';
import { MatPaginator, MatSort, MatSelect } from '@angular/material';
import {
  APIXUserService,
  XUserFilter,
  XUsersDataSource,
  ErrorHandlerService,
  Paginator,
  OrderBy,
  APIXGroupService,
} from '../../../services/core';
import { tap, merge, debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { fromEvent } from 'rxjs';
import { XGroup } from '../../../models/core';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit, AfterViewInit {

  xgroups: XGroup[] = [];
  dataSource: XUsersDataSource;
  displayedColumns= ['id', 'name', 'lastName', 'xgroupName', 'actived', 'operations'];

  @ViewChild(MatSort) sort: MatSort;
  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild('nameFilter') nameFilter: ElementRef;
  @ViewChild('lastNameFilter') lastNameFilter: ElementRef;
  @ViewChild('identificationFilter') identificationFilter: ElementRef;
  @ViewChild(MatSelect) xgroupFilter: MatSelect;

  initialPageSize = 20;
  pageSizeOptions = [20, 50, 100];

  constructor(private api: APIXUserService,
              private apiXG: APIXGroupService,
              private eh: ErrorHandlerService) {}

  ngOnInit() {
      this.apiXG.GetXGroups().subscribe(
        (xgroups) => {
          this.xgroups = xgroups;
        },
        (e) => this.eh.HandleError(e),
      );
      this.dataSource = new XUsersDataSource(this.api, this.eh);
      this.dataSource.load(
        true,
        new XUserFilter(
          null, null, null, null, null, null,
          new Paginator(
            0,
            this.initialPageSize,
          ),
          null,
        ),
      );
  }

  ngAfterViewInit() {
    fromEvent(this.nameFilter.nativeElement, 'keyup')
            .pipe(
                debounceTime(150),
                distinctUntilChanged(),
                tap(() => {
                    this.paginator.pageIndex = 0;
                    this.paginator.page.emit();
                    this.load(true);
                })
            )
            .subscribe();

    fromEvent(this.lastNameFilter.nativeElement, 'keyup')
            .pipe(
                debounceTime(150),
                distinctUntilChanged(),
                tap(() => {
                    this.paginator.pageIndex = 0;
                    this.paginator.page.emit();
                    this.load(true);
                })
            )
            .subscribe();

    fromEvent(this.identificationFilter.nativeElement, 'keyup')
            .pipe(
                debounceTime(150),
                distinctUntilChanged(),
                tap(() => {
                    this.paginator.pageIndex = 0;
                    this.paginator.page.emit();
                    this.load(true);
                })
            )
            .subscribe();

    this.xgroupFilter.valueChange.subscribe(
      () => {
        this.paginator.pageIndex = 0;
        this.paginator.page.emit();
        this.load(true);
      },
    );

    this.sort.sortChange.subscribe(
      () => {
        this.paginator.pageIndex = 0;
        this.paginator.page.emit();
      }
    );

    this.paginator.page
      .pipe(
        tap(() => this.load(false)),
      )
      .subscribe();
  }

  load(loadCount: boolean) {
    this.dataSource.load(
      loadCount,
      new XUserFilter(
        this.nameFilter.nativeElement.value,
        this.lastNameFilter.nativeElement.value,
        null,
        this.identificationFilter.nativeElement.value,
        null,
        this.xgroupFilter.value,
        new Paginator(
          this.paginator.pageIndex * this.paginator.pageSize,
          this.paginator.pageSize
        ),
        new OrderBy(
          this.sort.active,
          (this.sort.direction !== 'asc'),
        ),
      ),
    );
  }
}
