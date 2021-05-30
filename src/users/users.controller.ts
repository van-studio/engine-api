import { BadRequestException, Body, Controller, Get, Param, Post } from '@nestjs/common';
import { UsersService } from './users.service';
import { CreateUserBody } from './types';
import { argon2id, hash } from 'argon2';

@Controller('users')
export class UsersController {
  constructor(
    private user: UsersService,
  ) {
  }

  @Get()
  lists(): any {
    return this.user.find();
  }

  @Get(':id')
  get(@Param('id') id: any): any {
    return this.user.findOne({
      _id: id,
    });
  }

  @Get('pages/:page')
  pages(@Param('page') page: number): any {
    return {
      result: {
        page,
      },
    };
  }

  @Post()
  async create(@Body() body: CreateUserBody) {
    const exists = await this.user.findOne({
      $or: [
        { username: body.username },
        { email: body.email },
      ],
    });
    if (exists) {
      const field = body.username === exists.username ? 'username' : 'email';
      throw new BadRequestException([
        `The ${field} of [${body[field]}] already exists`,
      ]);
    }
    const password = await hash(body.password, {
      type: argon2id,
    });
    await this.user.create({
      username: body.username,
      email: body.email,
      password,
      call: body?.call,
    });
    return {
      message: ['ok'],
    };
  }
}
