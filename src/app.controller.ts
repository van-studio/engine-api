import { BadRequestException, Body, Controller, Get, Post, Put } from '@nestjs/common';
import { UsersService } from './users/users.service';
import { AuthBody } from './types';
import { argon2id, verify } from 'argon2';
import { CloudBaseService } from '@common/cloudbase.service';

@Controller()
export class AppController {
  constructor(
    private users: UsersService,
    private cloudBaseService: CloudBaseService,
  ) {
  }

  @Get()
  index() {
    this.cloudBaseService.db.collection('test').add({
      style: {
        name: '',
      },
      name: 'kain',
    });
    return {
      msg: 'hi',
    };
  }

  @Post('login')
  async auth(@Body() body: AuthBody) {
    // const data = await this.users.findOne(body.email);
    // if (!data) {
    //   throw new BadRequestException(['The email does not exist or password is incorrect']);
    // }
    // const check = await verify(data.password, body.password, {
    //   type: argon2id,
    // });
    // if (!check) {
    //   throw new BadRequestException(['The email does not exist or password is incorrect']);
    // }
    // // TODO: Factory Auth
    // return {
    //   message: ['ok'],
    // };
  }

  @Post('refresh_token')
  async refreshToken() {

  }

  @Put('update')
  async update() {
  }

  @Put('email')
  async email() {

  }
}
