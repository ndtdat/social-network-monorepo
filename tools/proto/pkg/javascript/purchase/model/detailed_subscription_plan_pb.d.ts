import * as jspb from 'google-protobuf'

import * as purchase_model_subscription_plan_tier_pb from '../../purchase/model/subscription_plan_tier_pb';


export class DetailedSubscriptionPlan extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): DetailedSubscriptionPlan;

  getSubscriptionPlanId(): number;
  setSubscriptionPlanId(value: number): DetailedSubscriptionPlan;

  getTier(): purchase_model_subscription_plan_tier_pb.SubscriptionPlanTier;
  setTier(value: purchase_model_subscription_plan_tier_pb.SubscriptionPlanTier): DetailedSubscriptionPlan;

  getCurrencySymbol(): string;
  setCurrencySymbol(value: string): DetailedSubscriptionPlan;

  getAmount(): string;
  setAmount(value: string): DetailedSubscriptionPlan;

  getDiscountAmount(): string;
  setDiscountAmount(value: string): DetailedSubscriptionPlan;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DetailedSubscriptionPlan.AsObject;
  static toObject(includeInstance: boolean, msg: DetailedSubscriptionPlan): DetailedSubscriptionPlan.AsObject;
  static serializeBinaryToWriter(message: DetailedSubscriptionPlan, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DetailedSubscriptionPlan;
  static deserializeBinaryFromReader(message: DetailedSubscriptionPlan, reader: jspb.BinaryReader): DetailedSubscriptionPlan;
}

export namespace DetailedSubscriptionPlan {
  export type AsObject = {
    userId: number,
    subscriptionPlanId: number,
    tier: purchase_model_subscription_plan_tier_pb.SubscriptionPlanTier,
    currencySymbol: string,
    amount: string,
    discountAmount: string,
  }
}

