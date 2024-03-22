// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_transit_gateway_route_table", transitGatewayRouteTableDataSource)
}

// transitGatewayRouteTableDataSource returns the Terraform awscc_ec2_transit_gateway_route_table data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::TransitGatewayRouteTable resource.
func transitGatewayRouteTableDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key of the associated tag key-value pair",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of the associated tag key-value pair",
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
						Description: "The key of the associated tag key-value pair",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of the associated tag key-value pair",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway.",
		//	  "type": "string"
		//	}
		"transit_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayRouteTableId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Transit Gateway Route Table primary identifier",
		//	  "type": "string"
		//	}
		"transit_gateway_route_table_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Transit Gateway Route Table primary identifier",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::TransitGatewayRouteTable",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayRouteTable").WithTerraformTypeName("awscc_ec2_transit_gateway_route_table")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                            "Key",
		"tags":                           "Tags",
		"transit_gateway_id":             "TransitGatewayId",
		"transit_gateway_route_table_id": "TransitGatewayRouteTableId",
		"value":                          "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
