// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package nimblestudio

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
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
	registry.AddResourceFactory("awscc_nimblestudio_streaming_image", streamingImageResource)
}

// streamingImageResource returns the Terraform awscc_nimblestudio_streaming_image resource.
// This Terraform resource corresponds to the CloudFormation AWS::NimbleStudio::StreamingImage resource.
func streamingImageResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eA human-readable description of the streaming image.\u003c/p\u003e",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>A human-readable description of the streaming image.</p>",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Ec2ImageId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe ID of an EC2 machine image with which to create this streaming image.\u003c/p\u003e",
		//	  "pattern": "^ami-[0-9A-z]+$",
		//	  "type": "string"
		//	}
		"ec_2_image_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The ID of an EC2 machine image with which to create this streaming image.</p>",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^ami-[0-9A-z]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EncryptionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "\u003cp\u003eTODO\u003c/p\u003e",
		//	  "properties": {
		//	    "KeyArn": {
		//	      "description": "\u003cp\u003eThe ARN for a KMS key that is used to encrypt studio data.\u003c/p\u003e",
		//	      "minLength": 4,
		//	      "pattern": "^arn:.*",
		//	      "type": "string"
		//	    },
		//	    "KeyType": {
		//	      "description": "\u003cp/\u003e",
		//	      "enum": [
		//	        "CUSTOMER_MANAGED_KEY"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "KeyType"
		//	  ],
		//	  "type": "object"
		//	}
		"encryption_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KeyArn
				"key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>The ARN for a KMS key that is used to encrypt studio data.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: KeyType
				"key_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p/>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "<p>TODO</p>",
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EulaIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe list of EULAs that must be accepted before a Streaming Session can be started using this streaming image.\u003c/p\u003e",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"eula_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "<p>The list of EULAs that must be accepted before a Streaming Session can be started using this streaming image.</p>",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eA friendly name for a streaming image resource.\u003c/p\u003e",
		//	  "maxLength": 64,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>A friendly name for a streaming image resource.</p>",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 64),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Owner
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe owner of the streaming image, either the studioId that contains the streaming image, or 'amazon' for images that are provided by Amazon Nimble Studio.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"owner": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The owner of the streaming image, either the studioId that contains the streaming image, or 'amazon' for images that are provided by Amazon Nimble Studio.</p>",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Platform
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe platform of the streaming image, either WINDOWS or LINUX.\u003c/p\u003e",
		//	  "pattern": "^[a-zA-Z]*$",
		//	  "type": "string"
		//	}
		"platform": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The platform of the streaming image, either WINDOWS or LINUX.</p>",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StreamingImageId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"streaming_image_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StudioId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe studioId. \u003c/p\u003e",
		//	  "type": "string"
		//	}
		"studio_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The studioId. </p>",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
				mapplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
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
		Description: "Represents a streaming session machine image that can be used to launch a streaming session",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::StreamingImage").WithTerraformTypeName("awscc_nimblestudio_streaming_image")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":              "Description",
		"ec_2_image_id":            "Ec2ImageId",
		"encryption_configuration": "EncryptionConfiguration",
		"eula_ids":                 "EulaIds",
		"key_arn":                  "KeyArn",
		"key_type":                 "KeyType",
		"name":                     "Name",
		"owner":                    "Owner",
		"platform":                 "Platform",
		"streaming_image_id":       "StreamingImageId",
		"studio_id":                "StudioId",
		"tags":                     "Tags",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
