// Copyright 2022 SaferPlace

// @generated by protoc-gen-es v0.0.9 with parameter "target=js+dts"
// @generated from file incident/v1/incident.proto (package incident.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";

/**
 * Resolution defines how the incident was resolved, or not yet.
 *
 * @generated from enum incident.v1.Resolution
 */
export declare enum Resolution {
  /**
   * @generated from enum value: RESOLUTION_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: RESOLUTION_REJECTED = 1;
   */
  REJECTED = 1,

  /**
   * @generated from enum value: RESOLUTION_ACCEPTED = 2;
   */
  ACCEPTED = 2,

  /**
   * @generated from enum value: RESOLUTION_ALERTED = 3;
   */
  ALERTED = 3,
}

/**
 * Incident defines an individual incident as it happened.
 *
 * @generated from message incident.v1.Incident
 */
export declare class Incident extends Message<Incident> {
  /**
   * ID of the report. This would be generated on ingestion, and therefore
   * does not have to be specified by the client as it will be overriden.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * Timestamp (unix) at which the incident occured, in seconds.
   *
   * @generated from field: int64 timestamp = 2;
   */
  timestamp: bigint;

  /**
   * Lat (lattitude) of the incident.
   *
   * @generated from field: float lat = 3;
   */
  lat: number;

  /**
   * Lon (Longitude) of the incident.
   *
   * @generated from field: float lon = 4;
   */
  lon: number;

  /**
   * Description provided by the user when sending the incident.
   *
   * @generated from field: string description = 5;
   */
  description: string;

  /**
   * Resolution of the incident.
   *
   * @generated from field: incident.v1.Resolution resolution = 6;
   */
  resolution: Resolution;

  /**
   * Comments provided by the reviewers.
   *
   * @generated from field: repeated incident.v1.Comment reviewer_comments = 7;
   */
  reviewerComments: Comment[];

  /**
   * Tags added by the reviewer to categorize the incident. This might be
   * updated in the future to maybe not be free-form but rather just specific
   * categories.
   *
   * @generated from field: repeated string tags = 8;
   */
  tags: string[];

  constructor(data?: PartialMessage<Incident>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "incident.v1.Incident";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Incident;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Incident;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Incident;

  static equals(a: Incident | PlainMessage<Incident> | undefined, b: Incident | PlainMessage<Incident> | undefined): boolean;
}

/**
 * Comment left by the reviewer.
 *
 * @generated from message incident.v1.Comment
 */
export declare class Comment extends Message<Comment> {
  /**
   * Timestamp (unix) of the comment, in seconds.
   *
   * @generated from field: int64 timestamp = 1;
   */
  timestamp: bigint;

  /**
   * AuthorID is the author of the comment
   *
   * @generated from field: string author_id = 2;
   */
  authorId: string;

  /**
   * Message left in the comment
   *
   * @generated from field: string message = 3;
   */
  message: string;

  constructor(data?: PartialMessage<Comment>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "incident.v1.Comment";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Comment;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Comment;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Comment;

  static equals(a: Comment | PlainMessage<Comment> | undefined, b: Comment | PlainMessage<Comment> | undefined): boolean;
}

