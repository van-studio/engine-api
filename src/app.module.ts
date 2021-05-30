import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { ConfigModule, ConfigService } from '@nestjs/config';

import { AppController } from './app.controller';
import { UsersModule } from './users/users.module';
import { TeamModule } from './team/team.module';
import { ProjectsModule } from './projects/projects.module';

@Module({
  imports: [
    UsersModule,
    TeamModule,
    ProjectsModule,
    ConfigModule.forRoot(),
    MongooseModule.forRootAsync({
      imports: [ConfigModule],
      useFactory: async (config: ConfigService) => ({
        uri: config.get<string>('database'),
      }),
      inject: [ConfigService],
    }),
  ],
  controllers: [
    AppController,
  ],
  providers: [],
})
export class AppModule {
}
