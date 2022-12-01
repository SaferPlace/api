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
export const ReviewService = {
  typeName: "review.v1.ReviewService",
  methods: {
    /**
     * ViewIncident displays the incident information for review.
     *
     * @generated from rpc review.v1.ReviewService.ViewIncident
     */
    viewIncident: {
      name: "ViewIncident",
      I: ViewIncidentRequest,
      O: ViewIncidentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ReviewIncident allows the client to send reviewed incident.
     *
     * @generated from rpc review.v1.ReviewService.ReviewIncident
     */
    reviewIncident: {
      name: "ReviewIncident",
      I: ReviewIncidentRequest,
      O: ReviewIncidentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * IncidentsWithoutReview shows basic information about incidents which
     * have not been reviewed yet.
     *
     * @generated from rpc review.v1.ReviewService.IncidentsWithoutReview
     */
    incidentsWithoutReview: {
      name: "IncidentsWithoutReview",
      I: IncidentsWithoutReviewRequest,
      O: IncidentsWithoutReviewResponse,
      kind: MethodKind.Unary,
    },
  }
};

