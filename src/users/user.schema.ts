import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Date, Document } from 'mongoose';

export type UserDoc = User & Document;

@Schema()
export class User {

  @Prop({
    required: true,
    unique: true,
  })
  email: string;

  @Prop({
    required: true,
  })
  password: string;

  @Prop()
  call?: string;

  @Prop({
    default: true,
  })
  status?: boolean;

  @Prop({
    type: Date,
  })
  create_time?: Date;

  @Prop({
    type: Date,
  })
  update_time?: Date;
}

export const UserSchema = SchemaFactory.createForClass(User);
