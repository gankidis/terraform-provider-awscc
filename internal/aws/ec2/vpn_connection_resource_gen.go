// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_vpn_connection", vPNConnectionResource)
}

// vPNConnectionResource returns the Terraform awscc_ec2_vpn_connection resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::VPNConnection resource.
func vPNConnectionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CustomerGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the customer gateway at your end of the VPN connection.",
		//	  "type": "string"
		//	}
		"customer_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the customer gateway at your end of the VPN connection.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StaticRoutesOnly
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the VPN connection uses static routes only.",
		//	  "type": "boolean"
		//	}
		"static_routes_only": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the VPN connection uses static routes only.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Any tags assigned to the VPN connection.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Any tags assigned to the VPN connection.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway associated with the VPN connection.",
		//	  "type": "string"
		//	}
		"transit_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway associated with the VPN connection.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of VPN connection.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of VPN connection.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpnConnectionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The provider-assigned unique ID for this managed resource",
		//	  "type": "string"
		//	}
		"vpn_connection_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The provider-assigned unique ID for this managed resource",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpnGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the virtual private gateway at the AWS side of the VPN connection.",
		//	  "type": "string"
		//	}
		"vpn_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the virtual private gateway at the AWS side of the VPN connection.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpnTunnelOptionsSpecifications
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tunnel options for the VPN connection.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "PreSharedKey": {
		//	        "type": "string"
		//	      },
		//	      "TunnelInsideCidr": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"vpn_tunnel_options_specifications": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: PreSharedKey
					"pre_shared_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: TunnelInsideCidr
					"tunnel_inside_cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tunnel options for the VPN connection.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::EC2::VPNConnection",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VPNConnection").WithTerraformTypeName("awscc_ec2_vpn_connection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"customer_gateway_id":               "CustomerGatewayId",
		"key":                               "Key",
		"pre_shared_key":                    "PreSharedKey",
		"static_routes_only":                "StaticRoutesOnly",
		"tags":                              "Tags",
		"transit_gateway_id":                "TransitGatewayId",
		"tunnel_inside_cidr":                "TunnelInsideCidr",
		"type":                              "Type",
		"value":                             "Value",
		"vpn_connection_id":                 "VpnConnectionId",
		"vpn_gateway_id":                    "VpnGatewayId",
		"vpn_tunnel_options_specifications": "VpnTunnelOptionsSpecifications",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
