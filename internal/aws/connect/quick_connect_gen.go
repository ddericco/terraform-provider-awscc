// Code generated by generators/resource/main.go; DO NOT EDIT.

package connect

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_connect_quick_connect", quickConnectResourceType)
}

// quickConnectResourceType returns the Terraform awscc_connect_quick_connect resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Connect::QuickConnect resource type.
func quickConnectResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the quick connect.",
			//   "maxLength": 250,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The description of the quick connect.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 250)},
			Optional:    true,
		},
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the Amazon Connect instance.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The identifier of the Amazon Connect instance.",
			Type:        types.StringType,
			Required:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the quick connect.",
			//   "maxLength": 127,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the quick connect.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 127)},
			Required:    true,
		},
		"quick_connect_arn": {
			// Property: QuickConnectArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the quick connect.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the quick connect.",
			Type:        types.StringType,
			Computed:    true,
		},
		"quick_connect_config": {
			// Property: QuickConnectConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration settings for the quick connect.",
			//   "properties": {
			//     "PhoneConfig": {
			//       "additionalProperties": false,
			//       "description": "The phone configuration. This is required only if QuickConnectType is PHONE_NUMBER.",
			//       "properties": {
			//         "PhoneNumber": {
			//           "description": "The phone number in E.164 format.",
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "PhoneNumber"
			//       ],
			//       "type": "object"
			//     },
			//     "QueueConfig": {
			//       "additionalProperties": false,
			//       "description": "The queue configuration. This is required only if QuickConnectType is QUEUE.",
			//       "properties": {
			//         "ContactFlowArn": {
			//           "description": "The identifier of the contact flow.",
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "QueueArn": {
			//           "description": "The identifier for the queue.",
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "ContactFlowArn",
			//         "QueueArn"
			//       ],
			//       "type": "object"
			//     },
			//     "QuickConnectType": {
			//       "description": "The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).",
			//       "enum": [
			//         "PHONE_NUMBER",
			//         "QUEUE",
			//         "USER"
			//       ],
			//       "type": "string"
			//     },
			//     "UserConfig": {
			//       "additionalProperties": false,
			//       "description": "The user configuration. This is required only if QuickConnectType is USER.",
			//       "properties": {
			//         "ContactFlowArn": {
			//           "description": "The identifier of the contact flow.",
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "UserArn": {
			//           "description": "The identifier of the user.",
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "ContactFlowArn",
			//         "UserArn"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "QuickConnectType"
			//   ],
			//   "type": "object"
			// }
			Description: "Configuration settings for the quick connect.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"phone_config": {
						// Property: PhoneConfig
						Description: "The phone configuration. This is required only if QuickConnectType is PHONE_NUMBER.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"phone_number": {
									// Property: PhoneNumber
									Description: "The phone number in E.164 format.",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Optional: true,
					},
					"queue_config": {
						// Property: QueueConfig
						Description: "The queue configuration. This is required only if QuickConnectType is QUEUE.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"contact_flow_arn": {
									// Property: ContactFlowArn
									Description: "The identifier of the contact flow.",
									Type:        types.StringType,
									Required:    true,
								},
								"queue_arn": {
									// Property: QueueArn
									Description: "The identifier for the queue.",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Optional: true,
					},
					"quick_connect_type": {
						// Property: QuickConnectType
						Description: "The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).",
						Type:        types.StringType,
						Required:    true,
					},
					"user_config": {
						// Property: UserConfig
						Description: "The user configuration. This is required only if QuickConnectType is USER.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"contact_flow_arn": {
									// Property: ContactFlowArn
									Description: "The identifier of the contact flow.",
									Type:        types.StringType,
									Required:    true,
								},
								"user_arn": {
									// Property: UserArn
									Description: "The identifier of the user.",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more tags.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "One or more tags.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 200,
				},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Connect::QuickConnect",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::QuickConnect").WithTerraformTypeName("awscc_connect_quick_connect")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"contact_flow_arn":     "ContactFlowArn",
		"description":          "Description",
		"instance_arn":         "InstanceArn",
		"key":                  "Key",
		"name":                 "Name",
		"phone_config":         "PhoneConfig",
		"phone_number":         "PhoneNumber",
		"queue_arn":            "QueueArn",
		"queue_config":         "QueueConfig",
		"quick_connect_arn":    "QuickConnectArn",
		"quick_connect_config": "QuickConnectConfig",
		"quick_connect_type":   "QuickConnectType",
		"tags":                 "Tags",
		"user_arn":             "UserArn",
		"user_config":          "UserConfig",
		"value":                "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_connect_quick_connect", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
