// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_ipam_pool", iPAMPoolResource)
}

// iPAMPoolResource returns the Terraform awscc_ec2_ipam_pool resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::IPAMPool resource.
func iPAMPoolResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AddressFamily
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The address family of the address space in this pool. Either IPv4 or IPv6.",
		//	  "type": "string"
		//	}
		"address_family": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The address family of the address space in this pool. Either IPv4 or IPv6.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AllocationDefaultNetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.",
		//	  "type": "integer"
		//	}
		"allocation_default_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AllocationMaxNetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum allowed netmask length for allocations made from this pool.",
		//	  "type": "integer"
		//	}
		"allocation_max_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum allowed netmask length for allocations made from this pool.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AllocationMinNetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The minimum allowed netmask length for allocations made from this pool.",
		//	  "type": "integer"
		//	}
		"allocation_min_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The minimum allowed netmask length for allocations made from this pool.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AllocationResourceTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When specified, an allocation will not be allowed unless a resource has a matching set of tags.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"allocation_resource_tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "When specified, an allocation will not be allowed unless a resource has a matching set of tags.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IPAM Pool.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IPAM Pool.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AutoImport
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.",
		//	  "type": "boolean"
		//	}
		"auto_import": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AwsService
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Limits which service in Amazon Web Services that the pool can be used in.",
		//	  "enum": [
		//	    "ec2"
		//	  ],
		//	  "type": "string"
		//	}
		"aws_service": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Limits which service in Amazon Web Services that the pool can be used in.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ec2",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpamArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IPAM this pool is a part of.",
		//	  "type": "string"
		//	}
		"ipam_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IPAM this pool is a part of.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the IPAM Pool.",
		//	  "type": "string"
		//	}
		"ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the IPAM Pool.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpamScopeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the scope this pool is a part of.",
		//	  "type": "string"
		//	}
		"ipam_scope_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the scope this pool is a part of.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpamScopeId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Id of the scope this pool is a part of.",
		//	  "type": "string"
		//	}
		"ipam_scope_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Id of the scope this pool is a part of.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpamScopeType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Determines whether this scope contains publicly routable space or space for a private network",
		//	  "enum": [
		//	    "public",
		//	    "private"
		//	  ],
		//	  "type": "string"
		//	}
		"ipam_scope_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Determines whether this scope contains publicly routable space or space for a private network",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Locale
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The region of this pool. If not set, this will default to \"None\" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match.",
		//	  "type": "string"
		//	}
		"locale": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The region of this pool. If not set, this will default to \"None\" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PoolDepth
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The depth of this pool in the source pool hierarchy.",
		//	  "type": "integer"
		//	}
		"pool_depth": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The depth of this pool in the source pool hierarchy.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProvisionedCidrs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of cidrs representing the address space available for allocation in this pool.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An address space to be inserted into this pool. All allocations must be made from this address space.",
		//	    "properties": {
		//	      "Cidr": {
		//	        "description": "Represents a single IPv4 or IPv6 CIDR",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Cidr"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"provisioned_cidrs": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Cidr
					"cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Represents a single IPv4 or IPv6 CIDR",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of cidrs representing the address space available for allocation in this pool.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PublicIpSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Default is `byoip`.",
		//	  "enum": [
		//	    "byoip",
		//	    "amazon"
		//	  ],
		//	  "type": "string"
		//	}
		"public_ip_source": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Default is `byoip`.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"byoip",
					"amazon",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PubliclyAdvertisable
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.",
		//	  "type": "boolean"
		//	}
		"publicly_advertisable": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceIpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Id of this pool's source. If set, all space provisioned in this pool must be free space provisioned in the parent pool.",
		//	  "type": "string"
		//	}
		"source_ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Id of this pool's source. If set, all space provisioned in this pool must be free space provisioned in the parent pool.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceResource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The resource associated with this pool's space. Depending on the ResourceType, setting a SourceResource changes which space can be provisioned in this pool and which types of resources can receive allocations",
		//	  "properties": {
		//	    "ResourceId": {
		//	      "type": "string"
		//	    },
		//	    "ResourceOwner": {
		//	      "type": "string"
		//	    },
		//	    "ResourceRegion": {
		//	      "type": "string"
		//	    },
		//	    "ResourceType": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "ResourceId",
		//	    "ResourceType",
		//	    "ResourceRegion",
		//	    "ResourceOwner"
		//	  ],
		//	  "type": "object"
		//	}
		"source_resource": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ResourceId
				"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: ResourceOwner
				"resource_owner": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: ResourceRegion
				"resource_region": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: ResourceType
				"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The resource associated with this pool's space. Depending on the ResourceType, setting a SourceResource changes which space can be provisioned in this pool and which types of resources can receive allocations",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of this pool. This can be one of the following values: \"create-in-progress\", \"create-complete\", \"modify-in-progress\", \"modify-complete\", \"delete-in-progress\", or \"delete-complete\"",
		//	  "enum": [
		//	    "create-in-progress",
		//	    "create-complete",
		//	    "modify-in-progress",
		//	    "modify-complete",
		//	    "delete-in-progress",
		//	    "delete-complete"
		//	  ],
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of this pool. This can be one of the following values: \"create-in-progress\", \"create-complete\", \"modify-in-progress\", \"modify-complete\", \"delete-in-progress\", or \"delete-complete\"",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StateMessage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An explanation of how the pool arrived at it current state.",
		//	  "type": "string"
		//	}
		"state_message": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An explanation of how the pool arrived at it current state.",
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
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Schema of AWS::EC2::IPAMPool Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::IPAMPool").WithTerraformTypeName("awscc_ec2_ipam_pool")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address_family":                    "AddressFamily",
		"allocation_default_netmask_length": "AllocationDefaultNetmaskLength",
		"allocation_max_netmask_length":     "AllocationMaxNetmaskLength",
		"allocation_min_netmask_length":     "AllocationMinNetmaskLength",
		"allocation_resource_tags":          "AllocationResourceTags",
		"arn":                               "Arn",
		"auto_import":                       "AutoImport",
		"aws_service":                       "AwsService",
		"cidr":                              "Cidr",
		"description":                       "Description",
		"ipam_arn":                          "IpamArn",
		"ipam_pool_id":                      "IpamPoolId",
		"ipam_scope_arn":                    "IpamScopeArn",
		"ipam_scope_id":                     "IpamScopeId",
		"ipam_scope_type":                   "IpamScopeType",
		"key":                               "Key",
		"locale":                            "Locale",
		"pool_depth":                        "PoolDepth",
		"provisioned_cidrs":                 "ProvisionedCidrs",
		"public_ip_source":                  "PublicIpSource",
		"publicly_advertisable":             "PubliclyAdvertisable",
		"resource_id":                       "ResourceId",
		"resource_owner":                    "ResourceOwner",
		"resource_region":                   "ResourceRegion",
		"resource_type":                     "ResourceType",
		"source_ipam_pool_id":               "SourceIpamPoolId",
		"source_resource":                   "SourceResource",
		"state":                             "State",
		"state_message":                     "StateMessage",
		"tags":                              "Tags",
		"value":                             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
