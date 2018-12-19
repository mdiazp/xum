import { Component, OnInit } from '@angular/core';
import {
  APIXGroupService,
  XGroupFilter,
  XGroupsDataSource,
  ErrorHandlerService,
} from '../../../services/core';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {

  dataSource: XGroupsDataSource;
  displayedColumns= ['id', 'name', 'description', 'actived'];

  constructor(private api: APIXGroupService,
              private eh: ErrorHandlerService) {}

  ngOnInit() {
      this.dataSource = new XGroupsDataSource(this.api, this.eh);
      this.dataSource.loadXGroups();
  }
}
