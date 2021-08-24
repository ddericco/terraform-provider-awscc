// Code generated by generators/resource/main.go; DO NOT EDIT.

package kinesis

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_kinesis_stream", streamResourceType)
}

// streamResourceType returns the Terraform awscc_kinesis_stream resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Kinesis::Stream resource type.
func streamResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon resource name (ARN) of the Kinesis stream",
			//   "type": "string"
			// }
			Description: "The Amazon resource name (ARN) of the Kinesis stream",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Kinesis stream.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the Kinesis stream.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
			Optional:    true,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"retention_period_hours": {
			// Property: RetentionPeriodHours
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of hours for the data records that are stored in shards to remain accessible.",
			//   "type": "integer"
			// }
			Description: "The number of hours for the data records that are stored in shards to remain accessible.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"shard_count": {
			// Property: ShardCount
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of shards that the stream uses.",
			//   "type": "integer"
			// }
			Description: "The number of shards that the stream uses.",
			Type:        types.NumberType,
			Required:    true,
		},
		"stream_encryption": {
			// Property: StreamEncryption
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream. Removing this property from your stack template and updating your stack disables encryption.",
			//   "properties": {
			//     "EncryptionType": {
			//       "description": "The encryption type to use. The only valid value is KMS. ",
			//       "enum": [
			//         "KMS"
			//       ],
			//       "type": "string"
			//     },
			//     "KeyId": {
			//       "description": "The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by \"alias/\".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.",
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "EncryptionType",
			//     "KeyId"
			//   ],
			//   "type": "object"
			// }
			Description: "When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream. Removing this property from your stack template and updating your stack disables encryption.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"encryption_type": {
						// Property: EncryptionType
						Description: "The encryption type to use. The only valid value is KMS. ",
						Type:        types.StringType,
						Required:    true,
					},
					"key_id": {
						// Property: KeyId
						Description: "The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by \"alias/\".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 2048)},
						Required:    true,
					},
				},
			),
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "An arbitrary set of tags (key-value pairs) to associate with the Kinesis stream.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.",
			Attributes: tfsdk.ListNestedAttributes(
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
						Description: "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLength(0, 255)},
						Required:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
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
		Description: "Resource Type definition for AWS::Kinesis::Stream",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Kinesis::Stream").WithTerraformTypeName("awscc_kinesis_stream")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"encryption_type":        "EncryptionType",
		"key":                    "Key",
		"key_id":                 "KeyId",
		"name":                   "Name",
		"retention_period_hours": "RetentionPeriodHours",
		"shard_count":            "ShardCount",
		"stream_encryption":      "StreamEncryption",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_kinesis_stream", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
