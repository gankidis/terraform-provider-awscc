// Code generated by generators/resource/main.go; DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_lambda_url", urlResourceType)
}

// urlResourceType returns the Terraform awscc_lambda_url resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Lambda::Url resource type.
func urlResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"auth_type": {
			// Property: AuthType
			// CloudFormation resource type schema:
			// {
			//   "description": "Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.",
			//   "enum": [
			//     "AWS_IAM",
			//     "NONE"
			//   ],
			//   "type": "string"
			// }
			Description: "Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"AWS_IAM",
					"NONE",
				}),
			},
		},
		"cors": {
			// Property: Cors
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AllowCredentials": {
			//       "description": "Specifies whether credentials are included in the CORS request.",
			//       "type": "boolean"
			//     },
			//     "AllowHeaders": {
			//       "description": "Represents a collection of allowed headers.",
			//       "insertionOrder": true,
			//       "items": {
			//         "maxLength": 1024,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "maxItems": 100,
			//       "minItems": 1,
			//       "type": "array"
			//     },
			//     "AllowMethods": {
			//       "description": "Represents a collection of allowed HTTP methods.",
			//       "insertionOrder": true,
			//       "items": {
			//         "enum": [
			//           "GET",
			//           "PUT",
			//           "HEAD",
			//           "POST",
			//           "PATCH",
			//           "DELETE",
			//           "*"
			//         ],
			//         "type": "string"
			//       },
			//       "maxItems": 6,
			//       "minItems": 1,
			//       "type": "array"
			//     },
			//     "AllowOrigins": {
			//       "description": "Represents a collection of allowed origins.",
			//       "insertionOrder": true,
			//       "items": {
			//         "maxLength": 253,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "maxItems": 100,
			//       "minItems": 1,
			//       "type": "array"
			//     },
			//     "ExposeHeaders": {
			//       "description": "Represents a collection of exposed headers.",
			//       "insertionOrder": true,
			//       "items": {
			//         "maxLength": 1024,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "maxItems": 100,
			//       "minItems": 1,
			//       "type": "array"
			//     },
			//     "MaxAge": {
			//       "maximum": 86400,
			//       "minimum": 0,
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"allow_credentials": {
						// Property: AllowCredentials
						Description: "Specifies whether credentials are included in the CORS request.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"allow_headers": {
						// Property: AllowHeaders
						Description: "Represents a collection of allowed headers.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 100),
							validate.ArrayForEach(validate.StringLenBetween(1, 1024)),
						},
					},
					"allow_methods": {
						// Property: AllowMethods
						Description: "Represents a collection of allowed HTTP methods.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 6),
							validate.ArrayForEach(validate.StringInSlice([]string{
								"GET",
								"PUT",
								"HEAD",
								"POST",
								"PATCH",
								"DELETE",
								"*",
							})),
						},
					},
					"allow_origins": {
						// Property: AllowOrigins
						Description: "Represents a collection of allowed origins.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 100),
							validate.ArrayForEach(validate.StringLenBetween(1, 253)),
						},
					},
					"expose_headers": {
						// Property: ExposeHeaders
						Description: "Represents a collection of exposed headers.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 100),
							validate.ArrayForEach(validate.StringLenBetween(1, 1024)),
						},
					},
					"max_age": {
						// Property: MaxAge
						Type:     types.Int64Type,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 86400),
						},
					},
				},
			),
			Optional: true,
		},
		"function_arn": {
			// Property: FunctionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The full Amazon Resource Name (ARN) of the function associated with the Function URL.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The full Amazon Resource Name (ARN) of the function associated with the Function URL.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"function_url": {
			// Property: FunctionUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "The generated url for this resource.",
			//   "type": "string"
			// }
			Description: "The generated url for this resource.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"invoke_mode": {
			// Property: InvokeMode
			// CloudFormation resource type schema:
			// {
			//   "description": "The invocation mode for the function?s URL. Set to BUFFERED if you want to buffer responses before returning them to the client. Set to RESPONSE_STREAM if you want to stream responses, allowing faster time to first byte and larger response payload sizes. If not set, defaults to BUFFERED.",
			//   "enum": [
			//     "BUFFERED",
			//     "RESPONSE_STREAM"
			//   ],
			//   "type": "string"
			// }
			Description: "The invocation mode for the function?s URL. Set to BUFFERED if you want to buffer responses before returning them to the client. Set to RESPONSE_STREAM if you want to stream responses, allowing faster time to first byte and larger response payload sizes. If not set, defaults to BUFFERED.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"BUFFERED",
					"RESPONSE_STREAM",
				}),
			},
		},
		"qualifier": {
			// Property: Qualifier
			// CloudFormation resource type schema:
			// {
			//   "description": "The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
			// Qualifier is a write-only property.
		},
		"target_function_arn": {
			// Property: TargetFunctionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the function associated with the Function URL.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the function associated with the Function URL.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
			// TargetFunctionArn is a write-only property.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Lambda::Url",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::Url").WithTerraformTypeName("awscc_lambda_url")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_credentials":   "AllowCredentials",
		"allow_headers":       "AllowHeaders",
		"allow_methods":       "AllowMethods",
		"allow_origins":       "AllowOrigins",
		"auth_type":           "AuthType",
		"cors":                "Cors",
		"expose_headers":      "ExposeHeaders",
		"function_arn":        "FunctionArn",
		"function_url":        "FunctionUrl",
		"invoke_mode":         "InvokeMode",
		"max_age":             "MaxAge",
		"qualifier":           "Qualifier",
		"target_function_arn": "TargetFunctionArn",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/TargetFunctionArn",
		"/properties/Qualifier",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}