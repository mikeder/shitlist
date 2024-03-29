// @generated by protoc-gen-es v1.2.1
// @generated from file shitlist/v1/shitlist.proto (package shitlist.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * ClickRequest is a request to record a click event.
 *
 * @generated from message shitlist.v1.ClickRequest
 */
export declare class ClickRequest extends Message<ClickRequest> {
  /**
   * user_id of the user to record a click event for.
   *
   * @generated from field: string user_id = 1;
   */
  userId: string;

  constructor(data?: PartialMessage<ClickRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "shitlist.v1.ClickRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClickRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClickRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClickRequest;

  static equals(a: ClickRequest | PlainMessage<ClickRequest> | undefined, b: ClickRequest | PlainMessage<ClickRequest> | undefined): boolean;
}

/**
 * ClickResponse is a response to a click event.
 *
 * @generated from message shitlist.v1.ClickResponse
 */
export declare class ClickResponse extends Message<ClickResponse> {
  /**
   * clicks recorded for the user.
   *
   * @generated from field: uint64 clicks = 1;
   */
  clicks: bigint;

  constructor(data?: PartialMessage<ClickResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "shitlist.v1.ClickResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClickResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClickResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClickResponse;

  static equals(a: ClickResponse | PlainMessage<ClickResponse> | undefined, b: ClickResponse | PlainMessage<ClickResponse> | undefined): boolean;
}

/**
 * LeadersRequest is a request for the top clickers.
 *
 * @generated from message shitlist.v1.LeadersRequest
 */
export declare class LeadersRequest extends Message<LeadersRequest> {
  constructor(data?: PartialMessage<LeadersRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "shitlist.v1.LeadersRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LeadersRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LeadersRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LeadersRequest;

  static equals(a: LeadersRequest | PlainMessage<LeadersRequest> | undefined, b: LeadersRequest | PlainMessage<LeadersRequest> | undefined): boolean;
}

/**
 * Clicker represents a single clicker user.
 *
 * @generated from message shitlist.v1.Clicker
 */
export declare class Clicker extends Message<Clicker> {
  /**
   * user_id of the user thats clicking.
   *
   * @generated from field: string user_id = 1;
   */
  userId: string;

  /**
   * clicks is the number of times the user has clicked.
   *
   * @generated from field: uint64 clicks = 2;
   */
  clicks: bigint;

  constructor(data?: PartialMessage<Clicker>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "shitlist.v1.Clicker";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Clicker;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Clicker;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Clicker;

  static equals(a: Clicker | PlainMessage<Clicker> | undefined, b: Clicker | PlainMessage<Clicker> | undefined): boolean;
}

/**
 * LeadersResponse is the top clickers.
 *
 * @generated from message shitlist.v1.LeadersResponse
 */
export declare class LeadersResponse extends Message<LeadersResponse> {
  /**
   * top_clickers are the top 10 clicking users.
   *
   * @generated from field: repeated shitlist.v1.Clicker top_clickers = 1;
   */
  topClickers: Clicker[];

  constructor(data?: PartialMessage<LeadersResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "shitlist.v1.LeadersResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LeadersResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LeadersResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LeadersResponse;

  static equals(a: LeadersResponse | PlainMessage<LeadersResponse> | undefined, b: LeadersResponse | PlainMessage<LeadersResponse> | undefined): boolean;
}

