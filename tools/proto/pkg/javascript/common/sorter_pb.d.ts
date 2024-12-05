import * as jspb from 'google-protobuf'

import * as validate_validate_pb from '../validate/validate_pb';


export class Sorter extends jspb.Message {
  getField(): string;
  setField(value: string): Sorter;

  getOrder(): string;
  setOrder(value: string): Sorter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Sorter.AsObject;
  static toObject(includeInstance: boolean, msg: Sorter): Sorter.AsObject;
  static serializeBinaryToWriter(message: Sorter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Sorter;
  static deserializeBinaryFromReader(message: Sorter, reader: jspb.BinaryReader): Sorter;
}

export namespace Sorter {
  export type AsObject = {
    field: string,
    order: string,
  }
}

