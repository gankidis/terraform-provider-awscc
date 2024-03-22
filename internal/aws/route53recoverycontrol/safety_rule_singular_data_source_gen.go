// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53recoverycontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53recoverycontrol_safety_rule", safetyRuleDataSource)
}

// safetyRuleDataSource returns the Terraform awscc_route53recoverycontrol_safety_rule data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53RecoveryControl::SafetyRule resource.
func safetyRuleDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssertionRule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An assertion rule enforces that, when a routing control state is changed, that the criteria set by the rule configuration is met. Otherwise, the change to the routing control is not accepted.",
		//	  "properties": {
		//	    "AssertedControls": {
		//	      "description": "The routing controls that are part of transactions that are evaluated to determine if a request to change a routing control state is allowed. For example, you might include three routing controls, one for each of three AWS Regions.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "WaitPeriodMs": {
		//	      "description": "An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail. This helps prevent \"flapping\" of state. The wait period is 5000 ms by default, but you can choose a custom value.",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "AssertedControls",
		//	    "WaitPeriodMs"
		//	  ],
		//	  "type": "object"
		//	}
		"assertion_rule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AssertedControls
				"asserted_controls": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The routing controls that are part of transactions that are evaluated to determine if a request to change a routing control state is allowed. For example, you might include three routing controls, one for each of three AWS Regions.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: WaitPeriodMs
				"wait_period_ms": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail. This helps prevent \"flapping\" of state. The wait period is 5000 ms by default, but you can choose a custom value.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An assertion rule enforces that, when a routing control state is changed, that the criteria set by the rule configuration is met. Otherwise, the change to the routing control is not accepted.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ControlPanelArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the control panel.",
		//	  "type": "string"
		//	}
		"control_panel_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the control panel.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GatingRule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A gating rule verifies that a set of gating controls evaluates as true, based on a rule configuration that you specify. If the gating rule evaluates to true, Amazon Route 53 Application Recovery Controller allows a set of routing control state changes to run and complete against the set of target controls.",
		//	  "properties": {
		//	    "GatingControls": {
		//	      "description": "The gating controls for the gating rule. That is, routing controls that are evaluated by the rule configuration that you specify.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "TargetControls": {
		//	      "description": "Routing controls that can only be set or unset if the specified RuleConfig evaluates to true for the specified GatingControls. For example, say you have three gating controls, one for each of three AWS Regions. Now you specify AtLeast 2 as your RuleConfig. With these settings, you can only change (set or unset) the routing controls that you have specified as TargetControls if that rule evaluates to true. \nIn other words, your ability to change the routing controls that you have specified as TargetControls is gated by the rule that you set for the routing controls in GatingControls.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "WaitPeriodMs": {
		//	      "description": "An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail. This helps prevent \"flapping\" of state. The wait period is 5000 ms by default, but you can choose a custom value.",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "WaitPeriodMs",
		//	    "TargetControls",
		//	    "GatingControls"
		//	  ],
		//	  "type": "object"
		//	}
		"gating_rule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: GatingControls
				"gating_controls": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The gating controls for the gating rule. That is, routing controls that are evaluated by the rule configuration that you specify.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: TargetControls
				"target_controls": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Routing controls that can only be set or unset if the specified RuleConfig evaluates to true for the specified GatingControls. For example, say you have three gating controls, one for each of three AWS Regions. Now you specify AtLeast 2 as your RuleConfig. With these settings, you can only change (set or unset) the routing controls that you have specified as TargetControls if that rule evaluates to true. \nIn other words, your ability to change the routing controls that you have specified as TargetControls is gated by the rule that you set for the routing controls in GatingControls.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: WaitPeriodMs
				"wait_period_ms": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail. This helps prevent \"flapping\" of state. The wait period is 5000 ms by default, but you can choose a custom value.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A gating rule verifies that a set of gating controls evaluates as true, based on a rule configuration that you specify. If the gating rule evaluates to true, Amazon Route 53 Application Recovery Controller allows a set of routing control state changes to run and complete against the set of target controls.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the safety rule.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the safety rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RuleConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The rule configuration for an assertion rule or gating rule. This is the criteria that you set for specific assertion controls (routing controls) or gating controls. This configuration specifies how many controls must be enabled after a transaction completes.",
		//	  "properties": {
		//	    "Inverted": {
		//	      "description": "Logical negation of the rule. If the rule would usually evaluate true, it's evaluated as false, and vice versa.",
		//	      "type": "boolean"
		//	    },
		//	    "Threshold": {
		//	      "description": "The value of N, when you specify an ATLEAST rule type. That is, Threshold is the number of controls that must be set when you specify an ATLEAST type.",
		//	      "type": "integer"
		//	    },
		//	    "Type": {
		//	      "description": "A rule can be one of the following: ATLEAST, AND, or OR.",
		//	      "enum": [
		//	        "AND",
		//	        "OR",
		//	        "ATLEAST"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type",
		//	    "Threshold",
		//	    "Inverted"
		//	  ],
		//	  "type": "object"
		//	}
		"rule_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Inverted
				"inverted": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Logical negation of the rule. If the rule would usually evaluate true, it's evaluated as false, and vice versa.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Threshold
				"threshold": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The value of N, when you specify an ATLEAST rule type. That is, Threshold is the number of controls that must be set when you specify an ATLEAST type.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A rule can be one of the following: ATLEAST, AND, or OR.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The rule configuration for an assertion rule or gating rule. This is the criteria that you set for specific assertion controls (routing controls) or gating controls. This configuration specifies how many controls must be enabled after a transaction completes.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SafetyRuleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the safety rule.",
		//	  "type": "string"
		//	}
		"safety_rule_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the safety rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
		//	  "enum": [
		//	    "PENDING",
		//	    "DEPLOYED",
		//	    "PENDING_DELETION"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of tags associated with a resource",
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
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A collection of tags associated with a resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53RecoveryControl::SafetyRule",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryControl::SafetyRule").WithTerraformTypeName("awscc_route53recoverycontrol_safety_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"asserted_controls": "AssertedControls",
		"assertion_rule":    "AssertionRule",
		"control_panel_arn": "ControlPanelArn",
		"gating_controls":   "GatingControls",
		"gating_rule":       "GatingRule",
		"inverted":          "Inverted",
		"key":               "Key",
		"name":              "Name",
		"rule_config":       "RuleConfig",
		"safety_rule_arn":   "SafetyRuleArn",
		"status":            "Status",
		"tags":              "Tags",
		"target_controls":   "TargetControls",
		"threshold":         "Threshold",
		"type":              "Type",
		"value":             "Value",
		"wait_period_ms":    "WaitPeriodMs",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
