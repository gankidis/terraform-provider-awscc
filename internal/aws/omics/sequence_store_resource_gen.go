// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package omics

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_omics_sequence_store", sequenceStoreResource)
}

// sequenceStoreResource returns the Terraform awscc_omics_sequence_store resource.
// This Terraform resource corresponds to the CloudFormation AWS::Omics::SequenceStore resource.
func sequenceStoreResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The store's ARN.",
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "pattern": "^arn:.+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The store's ARN.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When the store was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "When the store was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the store.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[\\p{L}||\\p{M}||\\p{Z}||\\p{S}||\\p{N}||\\p{P}]+$",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the store.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\p{L}||\\p{M}||\\p{Z}||\\p{S}||\\p{N}||\\p{P}]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FallbackLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An S3 URI representing the bucket and folder to store failed read set uploads.",
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"fallback_location": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An S3 URI representing the bucket and folder to store failed read set uploads.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the store.",
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "pattern": "^[\\p{L}||\\p{M}||\\p{Z}||\\p{S}||\\p{N}||\\p{P}]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the store.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 127),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\p{L}||\\p{M}||\\p{Z}||\\p{S}||\\p{N}||\\p{P}]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SequenceStoreId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 10,
		//	  "pattern": "^[0-9]+$",
		//	  "type": "string"
		//	}
		"sequence_store_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SseConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Server-side encryption (SSE) settings for a store.",
		//	  "properties": {
		//	    "KeyArn": {
		//	      "description": "An encryption key ARN.",
		//	      "maxLength": 2048,
		//	      "minLength": 20,
		//	      "pattern": "arn:([^:\n]*):([^:\n]*):([^:\n]*):([0-9]{12}):([^:\n]*)",
		//	      "type": "string"
		//	    },
		//	    "Type": {
		//	      "enum": [
		//	        "KMS"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type"
		//	  ],
		//	  "type": "object"
		//	}
		"sse_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KeyArn
				"key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "An encryption key ARN.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(20, 2048),
						stringvalidator.RegexMatches(regexp.MustCompile("arn:([^:\n]*):([^:\n]*):([^:\n]*):([0-9]{12}):([^:\n]*)"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"KMS",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Server-side encryption (SSE) settings for a store.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "minLength": 0,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
				mapplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// Tags is a write-only property.
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Definition of AWS::Omics::SequenceStore Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Omics::SequenceStore").WithTerraformTypeName("awscc_omics_sequence_store")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"creation_time":     "CreationTime",
		"description":       "Description",
		"fallback_location": "FallbackLocation",
		"key_arn":           "KeyArn",
		"name":              "Name",
		"sequence_store_id": "SequenceStoreId",
		"sse_config":        "SseConfig",
		"tags":              "Tags",
		"type":              "Type",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
