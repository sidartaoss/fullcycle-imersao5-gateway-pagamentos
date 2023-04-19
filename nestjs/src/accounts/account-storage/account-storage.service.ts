import { Injectable, Scope } from '@nestjs/common';
import { AccountsService } from '../accounts.service';
import { Account } from '../entities/account.entity';

@Injectable({
  scope: Scope.REQUEST,
})
export class AccountStorageService {
  private _account: Account | null;

  constructor(private accountService: AccountsService) {}

  get account() {
    return this._account;
  }

  async setBy(token: string) {
    this._account = await this.accountService.findOne(token);
  }
}
