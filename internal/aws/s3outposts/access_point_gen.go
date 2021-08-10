// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3outposts

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_s3outposts_access_point", accessPoint)
}

// accessPoint returns the Terraform aws_s3outposts_access_point resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::S3Outposts::AccessPoint resource type.
func accessPoint(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the specified AccessPoint.",
			     "maxLength": 2048,
			     "minLength": 20,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the specified AccessPoint.",
			Type:        types.StringType,
			Computed:    true,
		},
		"bucket": {
			// Property: Bucket
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.",
			     "maxLength": 2048,
			     "minLength": 20,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.",
			Type:        types.StringType,
			Required:    true,
			// Bucket is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A name for the AccessPoint.",
			     "maxLength": 50,
			     "minLength": 3,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "A name for the AccessPoint.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The access point policy associated with this access point.",
			     "type": "object"
			   }
			*/
			Description: "The access point policy associated with this access point.",
			Type:        types.MapType{ElemType: types.StringType},
			Optional:    true,
		},
		"vpc_configuration": {
			// Property: VpcConfiguration
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "VpcId": {
			         "description": "Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.",
			         "maxLength": 1024,
			         "minLength": 1,
			         "type": "string"
			       }
			     },
			     "$ref": "#/definitions/VpcConfiguration",
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"vpc_id": {
						// Property: VpcId
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.",
						     "maxLength": 1024,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Required: true,
			// VpcConfiguration is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type Definition for AWS::S3Outposts::AccessPoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Outposts::AccessPoint").WithTerraformTypeName("aws_s3outposts_access_point").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_s3outposts_access_point", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
