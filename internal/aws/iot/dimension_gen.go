// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

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
	registry.AddResourceTypeFactory("aws_iot_dimension", dimension)
}

// dimension returns the Terraform aws_iot_dimension resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::Dimension resource type.
func dimension(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN (Amazon resource name) of the created dimension.",
			     "type": "string"
			   }
			*/
			Description: "The ARN (Amazon resource name) of the created dimension.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A unique identifier for the dimension.",
			     "maxLength": 128,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "A unique identifier for the dimension.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"string_values": {
			// Property: StringValues
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Specifies the value or list of values for the dimension.",
			     "insertionOrder": false,
			     "items": {
			       "maxLength": 256,
			       "minLength": 1,
			       "type": "string"
			     },
			     "maxItems": 5,
			     "minItems": 1,
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "Specifies the value or list of values for the dimension.",
			Type:        providertypes.SetType{ElemType: types.StringType},
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Metadata that can be used to manage the dimension.",
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to associate with a resource.",
			       "properties": {
			         "Key": {
			           "description": "The tag's key.",
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "description": "The tag's value.",
			           "maxLength": 256,
			           "minLength": 1,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Value",
			         "Key"
			       ],
			       "type": "object"
			     },
			     "maxItems": 50,
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "Metadata that can be used to manage the dimension.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The tag's key.",
						     "maxLength": 128,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The tag's key.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The tag's value.",
						     "maxLength": 256,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The tag's value.",
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
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Specifies the type of the dimension.",
			     "enum": [
			       "TOPIC_FILTER"
			     ],
			     "type": "string"
			   }
			*/
			Description: "Specifies the type of the dimension.",
			Type:        types.StringType,
			Required:    true,
			// Type is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "A dimension can be used to limit the scope of a metric used in a security profile for AWS IoT Device Defender.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::Dimension").WithTerraformTypeName("aws_iot_dimension").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_iot_dimension", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
