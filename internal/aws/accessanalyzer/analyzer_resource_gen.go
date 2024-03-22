// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package accessanalyzer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
	registry.AddResourceFactory("awscc_accessanalyzer_analyzer", analyzerResource)
}

// analyzerResource returns the Terraform awscc_accessanalyzer_analyzer resource.
// This Terraform resource corresponds to the CloudFormation AWS::AccessAnalyzer::Analyzer resource.
func analyzerResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AnalyzerConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration for the analyzer",
		//	  "properties": {
		//	    "UnusedAccessConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "The Configuration for Unused Access Analyzer",
		//	      "properties": {
		//	        "UnusedAccessAge": {
		//	          "description": "The specified access age in days for which to generate findings for unused access. For example, if you specify 90 days, the analyzer will generate findings for IAM entities within the accounts of the selected organization for any access that hasn't been used in 90 or more days since the analyzer's last scan. You can choose a value between 1 and 180 days.",
		//	          "maximum": 180,
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"analyzer_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: UnusedAccessConfiguration
				"unused_access_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: UnusedAccessAge
						"unused_access_age": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The specified access age in days for which to generate findings for unused access. For example, if you specify 90 days, the analyzer will generate findings for IAM entities within the accounts of the selected organization for any access that hasn't been used in 90 or more days since the analyzer's last scan. You can choose a value between 1 and 180 days.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.Between(1, 180),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The Configuration for Unused Access Analyzer",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration for the analyzer",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AnalyzerName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Analyzer name",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"analyzer_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Analyzer name",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ArchiveRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An Access Analyzer archive rule. Archive rules automatically archive new findings that meet the criteria you define when you create the rule.",
		//	    "properties": {
		//	      "Filter": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Contains": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "Eq": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "Exists": {
		//	              "type": "boolean"
		//	            },
		//	            "Neq": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "Property": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Property"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "minItems": 1,
		//	        "type": "array"
		//	      },
		//	      "RuleName": {
		//	        "description": "The archive rule name",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Filter",
		//	      "RuleName"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"archive_rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Filter
					"filter": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Contains
								"contains": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										generic.Multiset(),
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Eq
								"eq": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										generic.Multiset(),
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Exists
								"exists": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
										boolplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Neq
								"neq": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										generic.Multiset(),
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Property
								"property": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Required: true,
						Validators: []validator.List{ /*START VALIDATORS*/
							listvalidator.SizeAtLeast(1),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							generic.Multiset(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: RuleName
					"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The archive rule name",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the analyzer",
		//	  "maxLength": 1600,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the analyzer",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 255,
		//	        "minLength": 1,
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 127),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 255),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the analyzer, must be one of ACCOUNT, ORGANIZATION, ACCOUNT_UNUSED_ACCESS or ORGANIZATION_UNUSED_ACCESS",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the analyzer, must be one of ACCOUNT, ORGANIZATION, ACCOUNT_UNUSED_ACCESS or ORGANIZATION_UNUSED_ACCESS",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "The AWS::AccessAnalyzer::Analyzer type specifies an analyzer of the user's account",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AccessAnalyzer::Analyzer").WithTerraformTypeName("awscc_accessanalyzer_analyzer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"analyzer_configuration":      "AnalyzerConfiguration",
		"analyzer_name":               "AnalyzerName",
		"archive_rules":               "ArchiveRules",
		"arn":                         "Arn",
		"contains":                    "Contains",
		"eq":                          "Eq",
		"exists":                      "Exists",
		"filter":                      "Filter",
		"key":                         "Key",
		"neq":                         "Neq",
		"property":                    "Property",
		"rule_name":                   "RuleName",
		"tags":                        "Tags",
		"type":                        "Type",
		"unused_access_age":           "UnusedAccessAge",
		"unused_access_configuration": "UnusedAccessConfiguration",
		"value":                       "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
