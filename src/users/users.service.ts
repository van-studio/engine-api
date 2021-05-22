import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { User, UserDoc } from './user.schema';
import { Model } from 'mongoose';

@Injectable()
export class UsersService {
  constructor(
    @InjectModel(User.name) private model: Model<UserDoc>,
  ) {
  }

  async findOne(email: string) {
    return this.model.findOne({ email });
  }

  async create(data: User) {
    return this.model.create(data);
  }
}
