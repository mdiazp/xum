import { User as user } from './user';

export class Session {
    constructor(public Token: string,
                public User: user) {}
}
