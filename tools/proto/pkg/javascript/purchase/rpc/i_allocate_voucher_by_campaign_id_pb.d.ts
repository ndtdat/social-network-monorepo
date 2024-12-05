import * as jspb from 'google-protobuf'

import * as validate_validate_pb from '../../validate/validate_pb';


export class IAllocateVoucherByCampaignIDRequest extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): IAllocateVoucherByCampaignIDRequest;

  getCampaignId(): number;
  setCampaignId(value: number): IAllocateVoucherByCampaignIDRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IAllocateVoucherByCampaignIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: IAllocateVoucherByCampaignIDRequest): IAllocateVoucherByCampaignIDRequest.AsObject;
  static serializeBinaryToWriter(message: IAllocateVoucherByCampaignIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IAllocateVoucherByCampaignIDRequest;
  static deserializeBinaryFromReader(message: IAllocateVoucherByCampaignIDRequest, reader: jspb.BinaryReader): IAllocateVoucherByCampaignIDRequest;
}

export namespace IAllocateVoucherByCampaignIDRequest {
  export type AsObject = {
    userId: number,
    campaignId: number,
  }
}

