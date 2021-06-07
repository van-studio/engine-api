import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Date, Document } from 'mongoose';

export type ProjectDoc = Project & Document;

@Schema({})
export class Project {
  @Prop({
    required: true,
    unique: true,
  })
  key: string;

  @Prop({
    required: true,
  })
  name: string;

  @Prop()
  description?: string;

  @Prop()
  partner: string[];

  @Prop()
  star: boolean;

  @Prop()
  label: string[];

  @Prop({
    type: Date,
    default: Date.now,
  })
  create_time?: number;

  @Prop({
    type: Date,
    default: Date.now,
  })
  update_time?: number;
}

export const ProjectSchema = SchemaFactory.createForClass(Project);
