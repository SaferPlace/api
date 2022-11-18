// Copyright 2022 SaferPlace

// @generated by protoc-gen-es v0.0.9 with parameter "target=js+dts"
// @generated from file report/v1/report.proto (package report.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";
import type {Incident} from "../../incident/v1/incident_pb.js";

/**
 * @generated from message report.v1.SendReportRequest
 */
export declare class SendReportRequest extends Message<SendReportRequest> {
  /**
   * @generated from field: incident.v1.Incident incident = 1;
   */
  incident?: Incident;

  constructor(data?: PartialMessage<SendReportRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "report.v1.SendReportRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendReportRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendReportRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendReportRequest;

  static equals(a: SendReportRequest | PlainMessage<SendReportRequest> | undefined, b: SendReportRequest | PlainMessage<SendReportRequest> | undefined): boolean;
}

/**
 * @generated from message report.v1.SendReportResponse
 */
export declare class SendReportResponse extends Message<SendReportResponse> {
  constructor(data?: PartialMessage<SendReportResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "report.v1.SendReportResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendReportResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendReportResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendReportResponse;

  static equals(a: SendReportResponse | PlainMessage<SendReportResponse> | undefined, b: SendReportResponse | PlainMessage<SendReportResponse> | undefined): boolean;
}

