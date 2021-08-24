// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

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
	registry.AddResourceTypeFactory("awscc_datasync_location_f_sx_windows", locationFSxWindowsResourceType)
}

// locationFSxWindowsResourceType returns the Terraform awscc_datasync_location_f_sx_windows resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DataSync::LocationFSxWindows resource type.
func locationFSxWindowsResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"domain": {
			// Property: Domain
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Windows domain that the FSx for Windows server belongs to.",
			//   "maxLength": 253,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the Windows domain that the FSx for Windows server belongs to.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Domain is a force-new attribute.
		},
		"fsx_filesystem_arn": {
			// Property: FsxFilesystemArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the FSx for Windows file system.",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the FSx for Windows file system.",
			Type:        types.StringType,
			Required:    true,
			// FsxFilesystemArn is a force-new attribute.
			// FsxFilesystemArn is a write-only attribute.
		},
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created.",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the FSx for Windows location that was described.",
			//   "maxLength": 4356,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The URL of the FSx for Windows location that was described.",
			Type:        types.StringType,
			Computed:    true,
		},
		"password": {
			// Property: Password
			// CloudFormation resource type schema:
			// {
			//   "description": "The password of the user who has the permissions to access files and folders in the FSx for Windows file system.",
			//   "maxLength": 104,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The password of the user who has the permissions to access files and folders in the FSx for Windows file system.",
			Type:        types.StringType,
			Required:    true,
			// Password is a force-new attribute.
			// Password is a write-only attribute.
		},
		"security_group_arns": {
			// Property: SecurityGroupArns
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARNs of the security groups that are to use to configure the FSx for Windows file system.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 128,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The ARNs of the security groups that are to use to configure the FSx for Windows file system.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
			// SecurityGroupArns is a force-new attribute.
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "A subdirectory in the location's path.",
			//   "maxLength": 4096,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A subdirectory in the location's path.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Subdirectory is a force-new attribute.
			// Subdirectory is a write-only attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
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
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 256)},
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 256)},
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"user": {
			// Property: User
			// CloudFormation resource type schema:
			// {
			//   "description": "The user who has the permissions to access files and folders in the FSx for Windows file system.",
			//   "maxLength": 104,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The user who has the permissions to access files and folders in the FSx for Windows file system.",
			Type:        types.StringType,
			Required:    true,
			// User is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::DataSync::LocationFSxWindows.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationFSxWindows").WithTerraformTypeName("awscc_datasync_location_f_sx_windows")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"domain":              "Domain",
		"fsx_filesystem_arn":  "FsxFilesystemArn",
		"key":                 "Key",
		"location_arn":        "LocationArn",
		"location_uri":        "LocationUri",
		"password":            "Password",
		"security_group_arns": "SecurityGroupArns",
		"subdirectory":        "Subdirectory",
		"tags":                "Tags",
		"user":                "User",
		"value":               "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Password",
		"/properties/Subdirectory",
		"/properties/FsxFilesystemArn",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_datasync_location_f_sx_windows", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
