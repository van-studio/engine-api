import { IsEmail, IsNotEmpty, IsString, Length, Matches } from 'class-validator';

export class CreateUserBody {
  @IsNotEmpty()
  @IsString()
  @Length(4, 16)
  username: string;

  @IsNotEmpty()
  @IsEmail()
  email: string;

  @IsNotEmpty()
  @IsString()
  @Length(12, 20)
  @Matches(
    '^(?=.*[a-z])(?=.*[A-Z])(?=.*[@$!%*?&-+])(?=.*[0-9])[\\w|@$!%*?&-+]+$',
    undefined,
    {
      message: 'The password consists of a combination of uppercase letters, lowercase letters, symbols, and numbers',
    },
  )
  password: string;

  call?: string;
}
