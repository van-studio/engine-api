import { Module } from '@nestjs/common';
import { CloudBaseService } from './cloudbase.service';

@Module({
  providers: [
    CloudBaseService,
  ],
  exports: [
    CloudBaseService,
  ],
})
export class CommonModule {
}
