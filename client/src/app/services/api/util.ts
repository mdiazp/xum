import { URLSearchParams } from '@angular/http';

export class Paginator {
  constructor(public offset: number,
              public limit: number) {}

  GetUSP(): URLSearchParams {
    let usp: URLSearchParams;
    usp = new URLSearchParams();
    if (this.limit !== null && this.limit !== 0 ) {
      usp.append('limit', this.limit.toString());
    }
    if (this.offset !== null) {
      usp.append('offset', this.offset.toString());
    }
    return usp;
  }
}

export class OrderBy {
  constructor(public by: string,
              public desc: boolean) {}

  GetUSP(): URLSearchParams {
    let usp: URLSearchParams;
    usp = new URLSearchParams();
    if (this.by !== null && this.by !== '' ) {
      usp.append('orderby', this.by);
    }
    if (this.desc !== null) {
      usp.append('desc', this.desc.toString());
    }
    return usp;
  }
}
