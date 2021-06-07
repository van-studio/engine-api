import { Controller, Get } from '@nestjs/common';
import { UsersService } from './users/users.service';

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
}
