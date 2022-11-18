// Copyright 2022 SaferPlace

// @generated by protoc-gen-es v0.0.9 with parameter "target=js+dts"
// @generated from file incident/v1/incident.proto (package incident.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {proto3} from "@bufbuild/protobuf";

/**
 * Resolution defines how the incident was resolved, or not yet.
 *
 * @generated from enum incident.v1.Resolution
 */
export const Resolution = proto3.makeEnum(
  "incident.v1.Resolution",
  [
    {no: 0, name: "RESOLUTION_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "RESOLUTION_REJECTED", localName: "REJECTED"},
    {no: 2, name: "RESOLUTION_ACCEPTED", localName: "ACCEPTED"},
    {no: 3, name: "RESOLUTION_ALERTED", localName: "ALERTED"},
  ],
);

/**
 * Incident defines an individual incident as it happened.
 *
 * @generated from message incident.v1.Incident
 */
export const Incident = proto3.makeMessageType(
  "incident.v1.Incident",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "timestamp", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "lat", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 4, name: "lon", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 5, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "resolution", kind: "enum", T: proto3.getEnumType(Resolution) },
    { no: 7, name: "reviewer_comments", kind: "message", T: Comment, repeated: true },
    { no: 8, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * Comment left by the reviewer.
 *
 * @generated from message incident.v1.Comment
 */
export const Comment = proto3.makeMessageType(
  "incident.v1.Comment",
  () => [
    { no: 1, name: "timestamp", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "author_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

