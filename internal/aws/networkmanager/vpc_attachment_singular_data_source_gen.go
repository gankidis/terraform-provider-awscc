// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_networkmanager_vpc_attachment", vpcAttachmentDataSourceType)
}

// vpcAttachmentDataSourceType returns the Terraform awscc_networkmanager_vpc_attachment data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::NetworkManager::VpcAttachment resource type.
func vpcAttachmentDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"attachment_id": {
			// Property: AttachmentId
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of the attachment.",
			//   "type": "string"
			// }
			Description: "Id of the attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"attachment_policy_rule_number": {
			// Property: AttachmentPolicyRuleNumber
			// CloudFormation resource type schema:
			// {
			//   "description": "The policy rule number associated with the attachment.",
			//   "type": "integer"
			// }
			Description: "The policy rule number associated with the attachment.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"attachment_type": {
			// Property: AttachmentType
			// CloudFormation resource type schema:
			// {
			//   "description": "Attachment type.",
			//   "type": "string"
			// }
			Description: "Attachment type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"core_network_arn": {
			// Property: CoreNetworkArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of a core network for the VPC attachment.",
			//   "type": "string"
			// }
			Description: "The ARN of a core network for the VPC attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"core_network_id": {
			// Property: CoreNetworkId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of a core network for the VPC attachment.",
			//   "type": "string"
			// }
			Description: "The ID of a core network for the VPC attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "Creation time of the attachment.",
			//   "type": "string"
			// }
			Description: "Creation time of the attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"edge_location": {
			// Property: EdgeLocation
			// CloudFormation resource type schema:
			// {
			//   "description": "The Region where the edge is located.",
			//   "type": "string"
			// }
			Description: "The Region where the edge is located.",
			Type:        types.StringType,
			Computed:    true,
		},
		"options": {
			// Property: Options
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Vpc options of the attachment.",
			//   "properties": {
			//     "Ipv6Support": {
			//       "default": false,
			//       "description": "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Vpc options of the attachment.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"ipv_6_support": {
						// Property: Ipv6Support
						Description: "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
						Type:        types.BoolType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"owner_account_id": {
			// Property: OwnerAccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "Owner account of the attachment.",
			//   "type": "string"
			// }
			Description: "Owner account of the attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"proposed_segment_change": {
			// Property: ProposedSegmentChange
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The attachment to move from one segment to another.",
			//   "properties": {
			//     "AttachmentPolicyRuleNumber": {
			//       "description": "The rule number in the policy document that applies to this change.",
			//       "type": "integer"
			//     },
			//     "SegmentName": {
			//       "description": "The name of the segment to change.",
			//       "type": "string"
			//     },
			//     "Tags": {
			//       "description": "The key-value tags that changed for the segment.",
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "A key-value pair to associate with a resource.",
			//         "insertionOrder": false,
			//         "properties": {
			//           "Key": {
			//             "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//             "type": "string"
			//           },
			//           "Value": {
			//             "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Key",
			//           "Value"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The attachment to move from one segment to another.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"attachment_policy_rule_number": {
						// Property: AttachmentPolicyRuleNumber
						Description: "The rule number in the policy document that applies to this change.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"segment_name": {
						// Property: SegmentName
						Description: "The name of the segment to change.",
						Type:        types.StringType,
						Computed:    true,
					},
					"tags": {
						// Property: Tags
						Description: "The key-value tags that changed for the segment.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
									Type:        types.StringType,
									Computed:    true,
								},
								"value": {
									// Property: Value
									Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"resource_arn": {
			// Property: ResourceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the Resource.",
			//   "type": "string"
			// }
			Description: "The ARN of the Resource.",
			Type:        types.StringType,
			Computed:    true,
		},
		"segment_name": {
			// Property: SegmentName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the segment attachment..",
			//   "type": "string"
			// }
			Description: "The name of the segment attachment..",
			Type:        types.StringType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "State of the attachment.",
			//   "type": "string"
			// }
			Description: "State of the attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subnet_arns": {
			// Property: SubnetArns
			// CloudFormation resource type schema:
			// {
			//   "description": "Subnet Arn list",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "Subnet Arn list",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags for the attachment.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "insertionOrder": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Tags for the attachment.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"updated_at": {
			// Property: UpdatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "Last update time of the attachment.",
			//   "type": "string"
			// }
			Description: "Last update time of the attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"vpc_arn": {
			// Property: VpcArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the VPC.",
			//   "type": "string"
			// }
			Description: "The ARN of the VPC.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::NetworkManager::VpcAttachment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::VpcAttachment").WithTerraformTypeName("awscc_networkmanager_vpc_attachment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attachment_id":                 "AttachmentId",
		"attachment_policy_rule_number": "AttachmentPolicyRuleNumber",
		"attachment_type":               "AttachmentType",
		"core_network_arn":              "CoreNetworkArn",
		"core_network_id":               "CoreNetworkId",
		"created_at":                    "CreatedAt",
		"edge_location":                 "EdgeLocation",
		"ipv_6_support":                 "Ipv6Support",
		"key":                           "Key",
		"options":                       "Options",
		"owner_account_id":              "OwnerAccountId",
		"proposed_segment_change":       "ProposedSegmentChange",
		"resource_arn":                  "ResourceArn",
		"segment_name":                  "SegmentName",
		"state":                         "State",
		"subnet_arns":                   "SubnetArns",
		"tags":                          "Tags",
		"updated_at":                    "UpdatedAt",
		"value":                         "Value",
		"vpc_arn":                       "VpcArn",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}