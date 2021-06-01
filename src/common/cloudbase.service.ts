import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { CloudBase, Database, init } from '@cloudbase/node-sdk';

@Injectable()
export class CloudBaseService {
  private cloudBase: CloudBase;
  db: Database.Db;

  constructor(
    config: ConfigService,
  ) {
    this.cloudBase = init({
      env: config.get('cloudbase_env'),
      secretId: config.get('cloudbase_id'),
      secretKey: config.get('cloudbase_key'),
    });
    this.db = this.cloudBase.database();
  }
}
