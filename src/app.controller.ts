import { BadRequestException, Body, Controller, Get, Post } from '@nestjs/common';
import { UsersService } from './users/users.service';
import { AuthBody } from './types';
import { argon2id, verify } from 'argon2';

@Controller()
export class AppController {
  constructor(
    private users: UsersService,
  ) {
  }

  @Get()
  index() {
    return {
      msg: 'hi',
    };
  }

  @Post('auth')
  async auth(@Body() body: AuthBody) {
    const data = await this.users.findOne(body.email);
    if (!data) {
      throw new BadRequestException(['The email does not exist or password is incorrect']);
    }
    const check = await verify(data.password, body.password, {
      type: argon2id,
    });
    if (!check) {
      throw new BadRequestException(['The email does not exist or password is incorrect']);
    }
    // TODO: Factory Auth
    return {
      message: ['ok'],
    };
  }
}
