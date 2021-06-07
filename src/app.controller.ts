import { Controller, Get, Post, Request, UseGuards } from '@nestjs/common';
import { AuthGuard } from '@nestjs/passport';
import { AuthService } from './auth/auth.service';

@Controller()
export class AppController {
  constructor(
    private auth: AuthService,
  ) {
  }


  @Get()
  index() {
    return {
      msg: 'hi',
    };
  }

  @UseGuards(AuthGuard('local'))
  @Post('auth/login')
  login(@Request() req) {
    return this.auth.login(req.email);
  }
}
