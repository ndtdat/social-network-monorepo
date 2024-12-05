/**
 * @fileoverview gRPC-Web generated client stub for purchase
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v0.0.0
// source: purchase/purchase.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('../google/api/annotations_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var purchase_rpc_i_allocate_voucher_by_campaign_id_pb = require('../purchase/rpc/i_allocate_voucher_by_campaign_id_pb.js')
const proto = {};
proto.purchase = require('./purchase_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.purchase.PurchaseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.purchase.PurchasePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rpc.IAllocateVoucherByCampaignIDRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_Purchase_IAllocateVoucherByCampaignID = new grpc.web.MethodDescriptor(
  '/purchase.Purchase/IAllocateVoucherByCampaignID',
  grpc.web.MethodType.UNARY,
  purchase_rpc_i_allocate_voucher_by_campaign_id_pb.IAllocateVoucherByCampaignIDRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.rpc.IAllocateVoucherByCampaignIDRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.rpc.IAllocateVoucherByCampaignIDRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.purchase.PurchaseClient.prototype.iAllocateVoucherByCampaignID =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/purchase.Purchase/IAllocateVoucherByCampaignID',
      request,
      metadata || {},
      methodDescriptor_Purchase_IAllocateVoucherByCampaignID,
      callback);
};


/**
 * @param {!proto.rpc.IAllocateVoucherByCampaignIDRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.purchase.PurchasePromiseClient.prototype.iAllocateVoucherByCampaignID =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/purchase.Purchase/IAllocateVoucherByCampaignID',
      request,
      metadata || {},
      methodDescriptor_Purchase_IAllocateVoucherByCampaignID);
};


module.exports = proto.purchase;

