// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotfleethub

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
	registry.AddResourceTypeFactory("awscc_iotfleethub_application", applicationResourceType)
}

// applicationResourceType returns the Terraform awscc_iotfleethub_application resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTFleetHub::Application resource type.
func applicationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"application_arn": {
			// Property: ApplicationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the application.",
			//   "maxLength": 1600,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ARN of the application.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 1600)},
			Computed:    true,
		},
		"application_creation_date": {
			// Property: ApplicationCreationDate
			// CloudFormation resource type schema:
			// {
			//   "description": "When the Application was created",
			//   "type": "integer"
			// }
			Description: "When the Application was created",
			Type:        types.NumberType,
			Computed:    true,
		},
		"application_description": {
			// Property: ApplicationDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "Application Description, should be between 1 and 2048 characters.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Application Description, should be between 1 and 2048 characters.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 2048)},
			Optional:    true,
		},
		"application_id": {
			// Property: ApplicationId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the application.",
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the application.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(36, 36)},
			Computed:    true,
		},
		"application_last_update_date": {
			// Property: ApplicationLastUpdateDate
			// CloudFormation resource type schema:
			// {
			//   "description": "When the Application was last updated",
			//   "type": "integer"
			// }
			Description: "When the Application was last updated",
			Type:        types.NumberType,
			Computed:    true,
		},
		"application_name": {
			// Property: ApplicationName
			// CloudFormation resource type schema:
			// {
			//   "description": "Application Name, should be between 1 and 256 characters.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Application Name, should be between 1 and 256 characters.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 256)},
			Required:    true,
		},
		"application_state": {
			// Property: ApplicationState
			// CloudFormation resource type schema:
			// {
			//   "description": "The current state of the application.",
			//   "type": "string"
			// }
			Description: "The current state of the application.",
			Type:        types.StringType,
			Computed:    true,
		},
		"application_url": {
			// Property: ApplicationUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the application.",
			//   "type": "string"
			// }
			Description: "The URL of the application.",
			Type:        types.StringType,
			Computed:    true,
		},
		"error_message": {
			// Property: ErrorMessage
			// CloudFormation resource type schema:
			// {
			//   "description": "A message indicating why Create or Delete Application failed.",
			//   "type": "string"
			// }
			Description: "A message indicating why Create or Delete Application failed.",
			Type:        types.StringType,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the role that the web application assumes when it interacts with AWS IoT Core. For more info on configuring this attribute, see https://docs.aws.amazon.com/iot/latest/apireference/API_iotfleethub_CreateApplication.html#API_iotfleethub_CreateApplication_RequestSyntax",
			//   "maxLength": 1600,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ARN of the role that the web application assumes when it interacts with AWS IoT Core. For more info on configuring this attribute, see https://docs.aws.amazon.com/iot/latest/apireference/API_iotfleethub_CreateApplication.html#API_iotfleethub_CreateApplication_RequestSyntax",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 1600)},
			Required:    true,
		},
		"sso_client_id": {
			// Property: SsoClientId
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS SSO application generated client ID (used with AWS SSO APIs).",
			//   "type": "string"
			// }
			Description: "The AWS SSO application generated client ID (used with AWS SSO APIs).",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the application.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the application.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 256)},
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 50,
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
		Description: "Resource schema for AWS::IoTFleetHub::Application",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTFleetHub::Application").WithTerraformTypeName("awscc_iotfleethub_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_arn":              "ApplicationArn",
		"application_creation_date":    "ApplicationCreationDate",
		"application_description":      "ApplicationDescription",
		"application_id":               "ApplicationId",
		"application_last_update_date": "ApplicationLastUpdateDate",
		"application_name":             "ApplicationName",
		"application_state":            "ApplicationState",
		"application_url":              "ApplicationUrl",
		"error_message":                "ErrorMessage",
		"key":                          "Key",
		"role_arn":                     "RoleArn",
		"sso_client_id":                "SsoClientId",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iotfleethub_application", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
