// Copyright 2023 SaferPlace

// @generated by protoc-gen-es v1.3.0 with parameter "target=js+dts"
// @generated from file viewer/v1/viewer.proto (package viewer.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Coordinates, Incident } from "../../incident/v1/incident_pb.js";

/**
 * ViewInRadiusRequest shows all incidents within specified radius (in meters).
 *
 * @generated from message viewer.v1.ViewInRadiusRequest
 */
export declare class ViewInRadiusRequest extends Message<ViewInRadiusRequest> {
  /**
   * Center of the search radius
   *
   * @generated from field: incident.v1.Coordinates center = 1;
   */
  center?: Coordinates;

  /**
   * Radius from the place that the incident happened in, in meters.
   *
   * @generated from field: double radius = 2;
   */
  radius: number;

  constructor(data?: PartialMessage<ViewInRadiusRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "viewer.v1.ViewInRadiusRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ViewInRadiusRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ViewInRadiusRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ViewInRadiusRequest;

  static equals(a: ViewInRadiusRequest | PlainMessage<ViewInRadiusRequest> | undefined, b: ViewInRadiusRequest | PlainMessage<ViewInRadiusRequest> | undefined): boolean;
}

/**
 * @generated from message viewer.v1.ViewInRadiusResponse
 */
export declare class ViewInRadiusResponse extends Message<ViewInRadiusResponse> {
  /**
   * @generated from field: repeated incident.v1.Incident incidents = 1;
   */
  incidents: Incident[];

  constructor(data?: PartialMessage<ViewInRadiusResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "viewer.v1.ViewInRadiusResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ViewInRadiusResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ViewInRadiusResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ViewInRadiusResponse;

  static equals(a: ViewInRadiusResponse | PlainMessage<ViewInRadiusResponse> | undefined, b: ViewInRadiusResponse | PlainMessage<ViewInRadiusResponse> | undefined): boolean;
}

