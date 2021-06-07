import { Injectable } from '@nestjs/common';
import { UsersService } from '../users/users.service';
import { verify } from 'argon2';

@Injectable()
export class AuthService {
  constructor(
    private users: UsersService,
  ) {
  }

  async validateUser(email: string, password: string): Promise<any> {
    const user = await this.users.findOne({
      email,
    });
    if (user && await verify(user.password, password)) {
      return true;
    }
    return null;
  }
}
