import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as purchase_rpc_buy_subscription_plan_pb from '../purchase/rpc/buy_subscription_plan_pb';
import * as purchase_rpc_i_allocate_voucher_by_campaign_id_pb from '../purchase/rpc/i_allocate_voucher_by_campaign_id_pb';


export class PurchaseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  iAllocateVoucherByCampaignID(
    request: purchase_rpc_i_allocate_voucher_by_campaign_id_pb.IAllocateVoucherByCampaignIDRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  buySubscriptionPlan(
    request: purchase_rpc_buy_subscription_plan_pb.BuySubscriptionPlanRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: purchase_rpc_buy_subscription_plan_pb.BuySubscriptionPlanReply) => void
  ): grpcWeb.ClientReadableStream<purchase_rpc_buy_subscription_plan_pb.BuySubscriptionPlanReply>;

}

export class PurchasePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  iAllocateVoucherByCampaignID(
    request: purchase_rpc_i_allocate_voucher_by_campaign_id_pb.IAllocateVoucherByCampaignIDRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  buySubscriptionPlan(
    request: purchase_rpc_buy_subscription_plan_pb.BuySubscriptionPlanRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<purchase_rpc_buy_subscription_plan_pb.BuySubscriptionPlanReply>;

}

