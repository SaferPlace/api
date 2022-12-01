// Copyright 2022 SaferPlace

// @generated by protoc-gen-connect-web v0.3.0 with parameter "target=js+dts"
// @generated from file review/v1/review.proto (package review.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {IncidentsWithoutReviewRequest, IncidentsWithoutReviewResponse, ReviewIncidentRequest, ReviewIncidentResponse, ViewIncidentRequest, ViewIncidentResponse} from "./review_pb.js";
import {MethodKind} from "@bufbuild/protobuf";

/**
 * Review API
 *
 * The review API is used by the reviewers to decide the faith of a given report
 *
 * @generated from service review.v1.ReviewService
 */
export declare const ReviewService: {
  readonly typeName: "review.v1.ReviewService",
  readonly methods: {
    /**
     * ViewIncident displays the incident information for review.
     *
     * @generated from rpc review.v1.ReviewService.ViewIncident
     */
    readonly viewIncident: {
      readonly name: "ViewIncident",
      readonly I: typeof ViewIncidentRequest,
      readonly O: typeof ViewIncidentResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * ReviewIncident allows the client to send reviewed incident.
     *
     * @generated from rpc review.v1.ReviewService.ReviewIncident
     */
    readonly reviewIncident: {
      readonly name: "ReviewIncident",
      readonly I: typeof ReviewIncidentRequest,
      readonly O: typeof ReviewIncidentResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * IncidentsWithoutReview shows basic information about incidents which
     * have not been reviewed yet.
     *
     * @generated from rpc review.v1.ReviewService.IncidentsWithoutReview
     */
    readonly incidentsWithoutReview: {
      readonly name: "IncidentsWithoutReview",
      readonly I: typeof IncidentsWithoutReviewRequest,
      readonly O: typeof IncidentsWithoutReviewResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

