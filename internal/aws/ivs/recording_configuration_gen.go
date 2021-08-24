// Code generated by generators/resource/main.go; DO NOT EDIT.

package ivs

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
	registry.AddResourceTypeFactory("awscc_ivs_recording_configuration", recordingConfigurationResourceType)
}

// recordingConfigurationResourceType returns the Terraform awscc_ivs_recording_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IVS::RecordingConfiguration resource type.
func recordingConfigurationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
			Computed:    true,
		},
		"destination_configuration": {
			// Property: DestinationConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Recording Destination Configuration.",
			//   "properties": {
			//     "S3": {
			//       "additionalProperties": false,
			//       "description": "Recording S3 Destination Configuration.",
			//       "properties": {
			//         "BucketName": {
			//           "maxLength": 63,
			//           "minLength": 3,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "BucketName"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "S3"
			//   ],
			//   "type": "object"
			// }
			Description: "Recording Destination Configuration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"s3": {
						// Property: S3
						Description: "Recording S3 Destination Configuration.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket_name": {
									// Property: BucketName
									Type:       types.StringType,
									Validators: []tfsdk.AttributeValidator{validate.StringLength(3, 63)},
									Required:   true,
									// BucketName is a force-new attribute.
								},
							},
						),
						Required: true,
						// S3 is a force-new attribute.
					},
				},
			),
			Required: true,
			// DestinationConfiguration is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Recording Configuration Name.",
			//   "maxLength": 128,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Recording Configuration Name.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(0, 128)},
			Optional:    true,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "Recording Configuration State.",
			//   "enum": [
			//     "CREATING",
			//     "CREATE_FAILED",
			//     "ACTIVE"
			//   ],
			//   "type": "string"
			// }
			Description: "Recording Configuration State.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the asset model.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the asset model.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
						Required:   true,
					},
					"value": {
						// Property: Value
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 256)},
						Required:   true,
					},
				},
				providertypes.SetNestedAttributesOptions{
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
		Description: "Resource Type definition for AWS::IVS::RecordingConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVS::RecordingConfiguration").WithTerraformTypeName("awscc_ivs_recording_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"bucket_name":               "BucketName",
		"destination_configuration": "DestinationConfiguration",
		"key":                       "Key",
		"name":                      "Name",
		"s3":                        "S3",
		"state":                     "State",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ivs_recording_configuration", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
