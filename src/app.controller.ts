import { Controller, Get, Post, Request, UseGuards } from '@nestjs/common';
import { AuthGuard } from '@nestjs/passport';

@Controller()
export class AppController {
  @Get()
  index() {
    return {
      msg: 'hi',
    };
  }

  @UseGuards(AuthGuard('local'))
  @Post('auth/login')
  login() {
  }
}
