import { Injectable } from '@nestjs/common';
import { UsersService } from '../users/users.service';
import { verify } from 'argon2';
import { JwtService } from '@nestjs/jwt';

@Injectable()
export class AuthService {
  constructor(
    private users: UsersService,
    private jwtService: JwtService,
  ) {
  }

  async validateUser(email: string, password: string): Promise<any> {
    const user = await this.users.findOne({
      email,
    });
    if (user && await verify(user.password, password)) {
      return {
        email: user.email,
      };
    }
    return null;
  }

  async login(email: any) {
    const payload = { sub: email };
    return {
      access_token: this.jwtService.sign(payload),
    };
  }
}
