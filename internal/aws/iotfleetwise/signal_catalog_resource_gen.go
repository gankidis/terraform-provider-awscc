// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_iotfleetwise_signal_catalog", signalCatalogResource)
}

// signalCatalogResource returns the Terraform awscc_iotfleetwise_signal_catalog resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTFleetWise::SignalCatalog resource.
func signalCatalogResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastModificationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"last_modification_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z\\d\\-_:]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 100),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_:]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NodeCounts
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "TotalActuators": {
		//	      "type": "number"
		//	    },
		//	    "TotalAttributes": {
		//	      "type": "number"
		//	    },
		//	    "TotalBranches": {
		//	      "type": "number"
		//	    },
		//	    "TotalNodes": {
		//	      "type": "number"
		//	    },
		//	    "TotalSensors": {
		//	      "type": "number"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"node_counts": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: TotalActuators
				"total_actuators": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TotalAttributes
				"total_attributes": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TotalBranches
				"total_branches": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TotalNodes
				"total_nodes": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TotalSensors
				"total_sensors": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Nodes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "properties": {
		//	      "Actuator": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "AllowedValues": {
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "minItems": 1,
		//	            "type": "array"
		//	          },
		//	          "AssignedValue": {
		//	            "type": "string"
		//	          },
		//	          "DataType": {
		//	            "enum": [
		//	              "INT8",
		//	              "UINT8",
		//	              "INT16",
		//	              "UINT16",
		//	              "INT32",
		//	              "UINT32",
		//	              "INT64",
		//	              "UINT64",
		//	              "BOOLEAN",
		//	              "FLOAT",
		//	              "DOUBLE",
		//	              "STRING",
		//	              "UNIX_TIMESTAMP",
		//	              "INT8_ARRAY",
		//	              "UINT8_ARRAY",
		//	              "INT16_ARRAY",
		//	              "UINT16_ARRAY",
		//	              "INT32_ARRAY",
		//	              "UINT32_ARRAY",
		//	              "INT64_ARRAY",
		//	              "UINT64_ARRAY",
		//	              "BOOLEAN_ARRAY",
		//	              "FLOAT_ARRAY",
		//	              "DOUBLE_ARRAY",
		//	              "STRING_ARRAY",
		//	              "UNIX_TIMESTAMP_ARRAY",
		//	              "UNKNOWN"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Description": {
		//	            "maxLength": 2048,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          },
		//	          "FullyQualifiedName": {
		//	            "type": "string"
		//	          },
		//	          "Max": {
		//	            "type": "number"
		//	          },
		//	          "Min": {
		//	            "type": "number"
		//	          },
		//	          "Unit": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "DataType",
		//	          "FullyQualifiedName"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Attribute": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "AllowedValues": {
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "minItems": 1,
		//	            "type": "array"
		//	          },
		//	          "AssignedValue": {
		//	            "type": "string"
		//	          },
		//	          "DataType": {
		//	            "enum": [
		//	              "INT8",
		//	              "UINT8",
		//	              "INT16",
		//	              "UINT16",
		//	              "INT32",
		//	              "UINT32",
		//	              "INT64",
		//	              "UINT64",
		//	              "BOOLEAN",
		//	              "FLOAT",
		//	              "DOUBLE",
		//	              "STRING",
		//	              "UNIX_TIMESTAMP",
		//	              "INT8_ARRAY",
		//	              "UINT8_ARRAY",
		//	              "INT16_ARRAY",
		//	              "UINT16_ARRAY",
		//	              "INT32_ARRAY",
		//	              "UINT32_ARRAY",
		//	              "INT64_ARRAY",
		//	              "UINT64_ARRAY",
		//	              "BOOLEAN_ARRAY",
		//	              "FLOAT_ARRAY",
		//	              "DOUBLE_ARRAY",
		//	              "STRING_ARRAY",
		//	              "UNIX_TIMESTAMP_ARRAY",
		//	              "UNKNOWN"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "DefaultValue": {
		//	            "type": "string"
		//	          },
		//	          "Description": {
		//	            "maxLength": 2048,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          },
		//	          "FullyQualifiedName": {
		//	            "type": "string"
		//	          },
		//	          "Max": {
		//	            "type": "number"
		//	          },
		//	          "Min": {
		//	            "type": "number"
		//	          },
		//	          "Unit": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "DataType",
		//	          "FullyQualifiedName"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Branch": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Description": {
		//	            "maxLength": 2048,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          },
		//	          "FullyQualifiedName": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "FullyQualifiedName"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Sensor": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "AllowedValues": {
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "minItems": 1,
		//	            "type": "array"
		//	          },
		//	          "DataType": {
		//	            "enum": [
		//	              "INT8",
		//	              "UINT8",
		//	              "INT16",
		//	              "UINT16",
		//	              "INT32",
		//	              "UINT32",
		//	              "INT64",
		//	              "UINT64",
		//	              "BOOLEAN",
		//	              "FLOAT",
		//	              "DOUBLE",
		//	              "STRING",
		//	              "UNIX_TIMESTAMP",
		//	              "INT8_ARRAY",
		//	              "UINT8_ARRAY",
		//	              "INT16_ARRAY",
		//	              "UINT16_ARRAY",
		//	              "INT32_ARRAY",
		//	              "UINT32_ARRAY",
		//	              "INT64_ARRAY",
		//	              "UINT64_ARRAY",
		//	              "BOOLEAN_ARRAY",
		//	              "FLOAT_ARRAY",
		//	              "DOUBLE_ARRAY",
		//	              "STRING_ARRAY",
		//	              "UNIX_TIMESTAMP_ARRAY",
		//	              "UNKNOWN"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Description": {
		//	            "maxLength": 2048,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          },
		//	          "FullyQualifiedName": {
		//	            "type": "string"
		//	          },
		//	          "Max": {
		//	            "type": "number"
		//	          },
		//	          "Min": {
		//	            "type": "number"
		//	          },
		//	          "Unit": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "DataType",
		//	          "FullyQualifiedName"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 500,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"nodes": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Actuator
					"actuator": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AllowedValues
							"allowed_values": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Optional:    true,
								Computed:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.SizeAtLeast(1),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									generic.Multiset(),
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: AssignedValue
							"assigned_value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: DataType
							"data_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"INT8",
										"UINT8",
										"INT16",
										"UINT16",
										"INT32",
										"UINT32",
										"INT64",
										"UINT64",
										"BOOLEAN",
										"FLOAT",
										"DOUBLE",
										"STRING",
										"UNIX_TIMESTAMP",
										"INT8_ARRAY",
										"UINT8_ARRAY",
										"INT16_ARRAY",
										"UINT16_ARRAY",
										"INT32_ARRAY",
										"UINT32_ARRAY",
										"INT64_ARRAY",
										"UINT64_ARRAY",
										"BOOLEAN_ARRAY",
										"FLOAT_ARRAY",
										"DOUBLE_ARRAY",
										"STRING_ARRAY",
										"UNIX_TIMESTAMP_ARRAY",
										"UNKNOWN",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 2048),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: FullyQualifiedName
							"fully_qualified_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
							}, /*END ATTRIBUTE*/
							// Property: Max
							"max": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Min
							"min": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Unit
							"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Attribute
					"attribute": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AllowedValues
							"allowed_values": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Optional:    true,
								Computed:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.SizeAtLeast(1),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									generic.Multiset(),
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: AssignedValue
							"assigned_value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: DataType
							"data_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"INT8",
										"UINT8",
										"INT16",
										"UINT16",
										"INT32",
										"UINT32",
										"INT64",
										"UINT64",
										"BOOLEAN",
										"FLOAT",
										"DOUBLE",
										"STRING",
										"UNIX_TIMESTAMP",
										"INT8_ARRAY",
										"UINT8_ARRAY",
										"INT16_ARRAY",
										"UINT16_ARRAY",
										"INT32_ARRAY",
										"UINT32_ARRAY",
										"INT64_ARRAY",
										"UINT64_ARRAY",
										"BOOLEAN_ARRAY",
										"FLOAT_ARRAY",
										"DOUBLE_ARRAY",
										"STRING_ARRAY",
										"UNIX_TIMESTAMP_ARRAY",
										"UNKNOWN",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: DefaultValue
							"default_value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 2048),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: FullyQualifiedName
							"fully_qualified_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
							}, /*END ATTRIBUTE*/
							// Property: Max
							"max": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Min
							"min": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Unit
							"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Branch
					"branch": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 2048),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: FullyQualifiedName
							"fully_qualified_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Sensor
					"sensor": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AllowedValues
							"allowed_values": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Optional:    true,
								Computed:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.SizeAtLeast(1),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									generic.Multiset(),
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: DataType
							"data_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"INT8",
										"UINT8",
										"INT16",
										"UINT16",
										"INT32",
										"UINT32",
										"INT64",
										"UINT64",
										"BOOLEAN",
										"FLOAT",
										"DOUBLE",
										"STRING",
										"UNIX_TIMESTAMP",
										"INT8_ARRAY",
										"UINT8_ARRAY",
										"INT16_ARRAY",
										"UINT16_ARRAY",
										"INT32_ARRAY",
										"UINT32_ARRAY",
										"INT64_ARRAY",
										"UINT64_ARRAY",
										"BOOLEAN_ARRAY",
										"FLOAT_ARRAY",
										"DOUBLE_ARRAY",
										"STRING_ARRAY",
										"UNIX_TIMESTAMP_ARRAY",
										"UNKNOWN",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 2048),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: FullyQualifiedName
							"fully_qualified_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
							}, /*END ATTRIBUTE*/
							// Property: Max
							"max": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Min
							"min": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Unit
							"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(1, 500),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(0, 50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Definition of AWS::IoTFleetWise::SignalCatalog Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTFleetWise::SignalCatalog").WithTerraformTypeName("awscc_iotfleetwise_signal_catalog")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actuator":               "Actuator",
		"allowed_values":         "AllowedValues",
		"arn":                    "Arn",
		"assigned_value":         "AssignedValue",
		"attribute":              "Attribute",
		"branch":                 "Branch",
		"creation_time":          "CreationTime",
		"data_type":              "DataType",
		"default_value":          "DefaultValue",
		"description":            "Description",
		"fully_qualified_name":   "FullyQualifiedName",
		"key":                    "Key",
		"last_modification_time": "LastModificationTime",
		"max":                    "Max",
		"min":                    "Min",
		"name":                   "Name",
		"node_counts":            "NodeCounts",
		"nodes":                  "Nodes",
		"sensor":                 "Sensor",
		"tags":                   "Tags",
		"total_actuators":        "TotalActuators",
		"total_attributes":       "TotalAttributes",
		"total_branches":         "TotalBranches",
		"total_nodes":            "TotalNodes",
		"total_sensors":          "TotalSensors",
		"unit":                   "Unit",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
