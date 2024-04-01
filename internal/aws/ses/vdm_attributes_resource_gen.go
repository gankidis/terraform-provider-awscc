// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ses

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ses_vdm_attributes", vdmAttributesResource)
}

// vdmAttributesResource returns the Terraform awscc_ses_vdm_attributes resource.
// This Terraform resource corresponds to the CloudFormation AWS::SES::VdmAttributes resource.
func vdmAttributesResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DashboardAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Preferences regarding the Dashboard feature.",
		//	  "properties": {
		//	    "EngagementMetrics": {
		//	      "description": "Whether emails sent from this account have engagement tracking enabled.",
		//	      "pattern": "ENABLED|DISABLED",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"dashboard_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EngagementMetrics
				"engagement_metrics": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Whether emails sent from this account have engagement tracking enabled.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("ENABLED|DISABLED"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Preferences regarding the Dashboard feature.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GuardianAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Preferences regarding the Guardian feature.",
		//	  "properties": {
		//	    "OptimizedSharedDelivery": {
		//	      "description": "Whether emails sent from this account have optimized delivery algorithm enabled.",
		//	      "pattern": "ENABLED|DISABLED",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"guardian_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: OptimizedSharedDelivery
				"optimized_shared_delivery": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Whether emails sent from this account have optimized delivery algorithm enabled.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("ENABLED|DISABLED"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Preferences regarding the Guardian feature.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VdmAttributesResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique identifier for this resource",
		//	  "type": "string"
		//	}
		"vdm_attributes_resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique identifier for this resource",
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
		Description: "Resource Type definition for AWS::SES::VdmAttributes",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::VdmAttributes").WithTerraformTypeName("awscc_ses_vdm_attributes")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dashboard_attributes":       "DashboardAttributes",
		"engagement_metrics":         "EngagementMetrics",
		"guardian_attributes":        "GuardianAttributes",
		"optimized_shared_delivery":  "OptimizedSharedDelivery",
		"vdm_attributes_resource_id": "VdmAttributesResourceId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
