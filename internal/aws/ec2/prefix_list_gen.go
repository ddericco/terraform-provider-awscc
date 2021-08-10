// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

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
	registry.AddResourceTypeFactory("aws_ec2_prefix_list", prefixList)
}

// prefixList returns the Terraform aws_ec2_prefix_list resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::PrefixList resource type.
func prefixList(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"address_family": {
			// Property: AddressFamily
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Ip Version of Prefix List.",
			     "enum": [
			       "IPv4",
			       "IPv6"
			     ],
			     "type": "string"
			   }
			*/
			Description: "Ip Version of Prefix List.",
			Type:        types.StringType,
			Required:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the Prefix List.",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the Prefix List.",
			Type:        types.StringType,
			Computed:    true,
		},
		"entries": {
			// Property: Entries
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Entries of Prefix List.",
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Cidr": {
			           "maxLength": 46,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Description": {
			           "maxLength": 255,
			           "minLength": 0,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Entry",
			       "required": [
			         "Cidr"
			       ],
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "Entries of Prefix List.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"cidr": {
						// Property: Cidr
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 46,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"description": {
						// Property: Description
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 255,
						     "minLength": 0,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"max_entries": {
			// Property: MaxEntries
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Max Entries of Prefix List.",
			     "type": "integer"
			   }
			*/
			Description: "Max Entries of Prefix List.",
			Type:        types.NumberType,
			Required:    true,
		},
		"owner_id": {
			// Property: OwnerId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Owner Id of Prefix List.",
			     "type": "string"
			   }
			*/
			Description: "Owner Id of Prefix List.",
			Type:        types.StringType,
			Computed:    true,
		},
		"prefix_list_id": {
			// Property: PrefixListId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Id of Prefix List.",
			     "type": "string"
			   }
			*/
			Description: "Id of Prefix List.",
			Type:        types.StringType,
			Computed:    true,
		},
		"prefix_list_name": {
			// Property: PrefixListName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Name of Prefix List.",
			     "maxLength": 255,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Name of Prefix List.",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Tags for Prefix List",
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "maxLength": 256,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Key"
			       ],
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "Tags for Prefix List",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 128,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 256,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Version of Prefix List.",
			     "type": "integer"
			   }
			*/
			Description: "Version of Prefix List.",
			Type:        types.NumberType,
			Computed:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema of AWS::EC2::PrefixList Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::PrefixList").WithTerraformTypeName("aws_ec2_prefix_list").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_ec2_prefix_list", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
