// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("aws_datasync_location_s3", locationS3)
}

// locationS3 returns the Terraform aws_datasync_location_s3 resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DataSync::LocationS3 resource type.
func locationS3(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the Amazon S3 bucket location.",
			     "maxLength": 128,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the Amazon S3 bucket location.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The URL of the S3 location that was described.",
			     "maxLength": 4356,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The URL of the S3 location that was described.",
			Type:        types.StringType,
			Computed:    true,
		},
		"s3_bucket_arn": {
			// Property: S3BucketArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the Amazon S3 bucket.",
			     "maxLength": 156,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the Amazon S3 bucket.",
			Type:        types.StringType,
			Required:    true,
			// S3BucketArn is a force-new attribute.
			// S3BucketArn is a write-only attribute.
		},
		"s3_config": {
			// Property: S3Config
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The Amazon Resource Name (ARN) of the AWS IAM role that is used to access an Amazon S3 bucket.",
			     "properties": {
			       "BucketAccessRoleArn": {
			         "description": "The ARN of the IAM role of the Amazon S3 bucket.",
			         "maxLength": 2048,
			         "pattern": "",
			         "type": "string"
			       }
			     },
			     "$ref": "#/definitions/S3Config",
			     "required": [
			       "BucketAccessRoleArn"
			     ],
			     "type": "object"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the AWS IAM role that is used to access an Amazon S3 bucket.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"bucket_access_role_arn": {
						// Property: BucketAccessRoleArn
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The ARN of the IAM role of the Amazon S3 bucket.",
						     "maxLength": 2048,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Description: "The ARN of the IAM role of the Amazon S3 bucket.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Required: true,
			// S3Config is a force-new attribute.
		},
		"s3_storage_class": {
			// Property: S3StorageClass
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon S3 storage class you want to store your files in when this location is used as a task destination.",
			     "enum": [
			       "STANDARD",
			       "STANDARD_IA",
			       "ONEZONE_IA",
			       "INTELLIGENT_TIERING",
			       "GLACIER",
			       "DEEP_ARCHIVE"
			     ],
			     "type": "string"
			   }
			*/
			Description: "The Amazon S3 storage class you want to store your files in when this location is used as a task destination.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// S3StorageClass is a force-new attribute.
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A subdirectory in the Amazon S3 bucket. This subdirectory in Amazon S3 is used to read data from the S3 source location or write data to the S3 destination.",
			     "maxLength": 4096,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "A subdirectory in the Amazon S3 bucket. This subdirectory in Amazon S3 is used to read data from the S3 source location or write data to the S3 destination.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Subdirectory is a force-new attribute.
			// Subdirectory is a write-only attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "An array of key-value pairs to apply to this resource.",
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to associate with a resource.",
			       "properties": {
			         "Key": {
			           "description": "The key for an AWS resource tag.",
			           "maxLength": 256,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "Value": {
			           "description": "The value for an AWS resource tag.",
			           "maxLength": 256,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "maxItems": 50,
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The key for an AWS resource tag.",
						     "maxLength": 256,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The value for an AWS resource tag.",
						     "maxLength": 256,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::DataSync::LocationS3",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationS3").WithTerraformTypeName("aws_datasync_location_s3").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Subdirectory",
		"/properties/S3BucketArn",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_datasync_location_s3", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
