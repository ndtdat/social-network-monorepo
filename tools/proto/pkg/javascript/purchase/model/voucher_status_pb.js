// source: purchase/model/voucher_status.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

goog.exportSymbol('proto.model.VoucherStatus', null, global);
/**
 * @enum {number}
 */
proto.model.VoucherStatus = {
  VS_NONE: 0,
  VS_DRAFT: 1,
  VS_AVAILABLE: 2,
  VS_UNAVAILABLE: 3
};

goog.object.extend(exports, proto.model);
