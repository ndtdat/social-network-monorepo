import * as jspb from 'google-protobuf'

import * as validate_validate_pb from '../../validate/validate_pb';


export class RegisterRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): RegisterRequest;

  getPassword(): string;
  setPassword(value: string): RegisterRequest;

  getCampaignCode(): string;
  setCampaignCode(value: string): RegisterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterRequest): RegisterRequest.AsObject;
  static serializeBinaryToWriter(message: RegisterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterRequest;
  static deserializeBinaryFromReader(message: RegisterRequest, reader: jspb.BinaryReader): RegisterRequest;
}

export namespace RegisterRequest {
  export type AsObject = {
    email: string,
    password: string,
    campaignCode: string,
  }
}

export class RegisterReply extends jspb.Message {
  getAccessToken(): string;
  setAccessToken(value: string): RegisterReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterReply.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterReply): RegisterReply.AsObject;
  static serializeBinaryToWriter(message: RegisterReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterReply;
  static deserializeBinaryFromReader(message: RegisterReply, reader: jspb.BinaryReader): RegisterReply;
}

export namespace RegisterReply {
  export type AsObject = {
    accessToken: string,
  }
}

