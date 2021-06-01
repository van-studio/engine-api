import { Injectable } from '@nestjs/common';
import { CloudBaseService } from '@common/cloudbase.service';
import { Database } from '@cloudbase/node-sdk';

@Injectable()
export class UsersService {
  private collection: Database.CollectionReference;
  private projection: any = { password: false };

  constructor(
    private cloudBaseService: CloudBaseService,
  ) {
    this.collection = cloudBaseService.db.collection('users');
  }

  async find(query: any = {}, projection = true) {
    const { data } = await this.collection
      .where(query)
      .field(projection ? this.projection : null)
      .get();
    return data;
  }

  async findOne(query: any = {}, projection = true) {
    const data = await this.find(query, projection);
    return data.length !== 0 ? data[0] : undefined;
  }

  create(data: any) {
    return this.collection.add(data);
  }


}
