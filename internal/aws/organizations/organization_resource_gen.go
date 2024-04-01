// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_organizations_organization", organizationResource)
}

// organizationResource returns the Terraform awscc_organizations_organization resource.
// This Terraform resource corresponds to the CloudFormation AWS::Organizations::Organization resource.
func organizationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of an organization.",
		//	  "pattern": "^arn:aws.*:organizations::\\d{12}:organization\\/o-[a-z0-9]{10,32}",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of an organization.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FeatureSet
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "ALL",
		//	  "description": "Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.",
		//	  "enum": [
		//	    "ALL",
		//	    "CONSOLIDATED_BILLING"
		//	  ],
		//	  "type": "string"
		//	}
		"feature_set": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ALL",
					"CONSOLIDATED_BILLING",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("ALL"),
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier (ID) of an organization.",
		//	  "pattern": "^o-[a-z0-9]{10,32}$",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier (ID) of an organization.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ManagementAccountArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the account that is designated as the management account for the organization.",
		//	  "pattern": "^arn:aws.*:organizations::\\d{12}:account\\/o-[a-z0-9]{10,32}\\/\\d{12}",
		//	  "type": "string"
		//	}
		"management_account_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the account that is designated as the management account for the organization.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ManagementAccountEmail
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The email address that is associated with the AWS account that is designated as the management account for the organization.",
		//	  "maxLength": 64,
		//	  "minLength": 6,
		//	  "pattern": "[^\\s@]+@[^\\s@]+\\.[^\\s@]+",
		//	  "type": "string"
		//	}
		"management_account_email": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The email address that is associated with the AWS account that is designated as the management account for the organization.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ManagementAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier (ID) of the management account of an organization.",
		//	  "pattern": "^\\d{12}$",
		//	  "type": "string"
		//	}
		"management_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier (ID) of the management account of an organization.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RootId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier (ID) for the root.",
		//	  "maxLength": 64,
		//	  "pattern": "^r-[0-9a-z]{4,32}$",
		//	  "type": "string"
		//	}
		"root_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier (ID) for the root.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Resource schema for AWS::Organizations::Organization",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Organizations::Organization").WithTerraformTypeName("awscc_organizations_organization")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                      "Arn",
		"feature_set":              "FeatureSet",
		"management_account_arn":   "ManagementAccountArn",
		"management_account_email": "ManagementAccountEmail",
		"management_account_id":    "ManagementAccountId",
		"organization_id":          "Id",
		"root_id":                  "RootId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
