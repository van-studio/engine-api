import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { CloudBase as App, Database, init as AppInit } from '@cloudbase/node-sdk';

const CloudBase = require('@cloudbase/manager-node');

@Injectable()
export class CloudBaseService {
  private manager: any;
  private app: App;
  db: Database.Db;

  constructor(
    config: ConfigService,
  ) {
    const cfg = {
      env: config.get('cloudbase_env'),
      secretId: config.get('cloudbase_id'),
      secretKey: config.get('cloudbase_key'),
    };
    this.manager = CloudBase.init(cfg);
    this.app = AppInit(cfg);
    this.db = this.app.database();
  }

  DbManager() {
    return this.manager.database;
  }
}
