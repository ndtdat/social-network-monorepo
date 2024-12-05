import * as jspb from 'google-protobuf'

import * as validate_validate_pb from '../../validate/validate_pb';
import * as purchase_model_detailed_subscription_plan_pb from '../../purchase/model/detailed_subscription_plan_pb';


export class BuySubscriptionPlanRequest extends jspb.Message {
  getSubscriptionPlanId(): number;
  setSubscriptionPlanId(value: number): BuySubscriptionPlanRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BuySubscriptionPlanRequest.AsObject;
  static toObject(includeInstance: boolean, msg: BuySubscriptionPlanRequest): BuySubscriptionPlanRequest.AsObject;
  static serializeBinaryToWriter(message: BuySubscriptionPlanRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BuySubscriptionPlanRequest;
  static deserializeBinaryFromReader(message: BuySubscriptionPlanRequest, reader: jspb.BinaryReader): BuySubscriptionPlanRequest;
}

export namespace BuySubscriptionPlanRequest {
  export type AsObject = {
    subscriptionPlanId: number,
  }
}

export class BuySubscriptionPlanReply extends jspb.Message {
  getData(): purchase_model_detailed_subscription_plan_pb.DetailedSubscriptionPlan | undefined;
  setData(value?: purchase_model_detailed_subscription_plan_pb.DetailedSubscriptionPlan): BuySubscriptionPlanReply;
  hasData(): boolean;
  clearData(): BuySubscriptionPlanReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BuySubscriptionPlanReply.AsObject;
  static toObject(includeInstance: boolean, msg: BuySubscriptionPlanReply): BuySubscriptionPlanReply.AsObject;
  static serializeBinaryToWriter(message: BuySubscriptionPlanReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BuySubscriptionPlanReply;
  static deserializeBinaryFromReader(message: BuySubscriptionPlanReply, reader: jspb.BinaryReader): BuySubscriptionPlanReply;
}

export namespace BuySubscriptionPlanReply {
  export type AsObject = {
    data?: purchase_model_detailed_subscription_plan_pb.DetailedSubscriptionPlan.AsObject,
  }
}

