import { BadRequestException, Body, Controller, Get, HttpException, Param, Post, ValidationPipe } from '@nestjs/common';
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
  findAll(): any {
    return [
      { username: 'kain' },
      { username: 'kain1' },
      { username: 'kain2' },
    ];
  }

  @Get(':id')
  findOne(@Param('id') id: any): any {
    return { id, username: 'kain' };
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
    const exists = await this.user.findOne(body.email);
    if (exists) {
      throw new BadRequestException([
        `The email of [${body.email}] already exists.`,
      ]);
    }
    const password = await hash(body.password, {
      type: argon2id,
    });
    await this.user.create({
      email: body.email,
      password,
    });
    return {
      message: ['ok'],
    };
  }
}
