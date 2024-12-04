// source: purchase/model/voucher_calculation_type.proto
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

goog.exportSymbol('proto.model.VoucherCalculationType', null, global);
/**
 * @enum {number}
 */
proto.model.VoucherCalculationType = {
  VCT_NONE: 0,
  VCT_PERCENTAGE: 1,
  VCT_AMOUNT: 2
};

goog.object.extend(exports, proto.model);
