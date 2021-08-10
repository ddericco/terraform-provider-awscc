// Code generated by generators/resource/main.go; DO NOT EDIT.

package xray

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
	registry.AddResourceTypeFactory("aws_xray_group", group)
}

// group returns the Terraform aws_xray_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::XRay::Group resource type.
func group(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"filter_expression": {
			// Property: FilterExpression
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The filter expression defining criteria by which to group traces.",
			     "type": "string"
			   }
			*/
			Description: "The filter expression defining criteria by which to group traces.",
			Type:        types.StringType,
			Optional:    true,
		},
		"group_arn": {
			// Property: GroupARN
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the group that was generated on creation.",
			     "maxLength": 400,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "The ARN of the group that was generated on creation.",
			Type:        types.StringType,
			Computed:    true,
		},
		"group_name": {
			// Property: GroupName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The case-sensitive name of the new group. Names must be unique.",
			     "maxLength": 32,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "The case-sensitive name of the new group. Names must be unique.",
			Type:        types.StringType,
			Optional:    true,
		},
		"insights_configuration": {
			// Property: InsightsConfiguration
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "InsightsEnabled": {
			         "description": "Set the InsightsEnabled value to true to enable insights or false to disable insights.",
			         "type": "boolean"
			       },
			       "NotificationsEnabled": {
			         "description": "Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be enabled on a group with InsightsEnabled set to true.",
			         "type": "boolean"
			       }
			     },
			     "$ref": "#/definitions/InsightsConfiguration",
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"insights_enabled": {
						// Property: InsightsEnabled
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Set the InsightsEnabled value to true to enable insights or false to disable insights.",
						     "type": "boolean"
						   }
						*/
						Description: "Set the InsightsEnabled value to true to enable insights or false to disable insights.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"notifications_enabled": {
						// Property: NotificationsEnabled
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be enabled on a group with InsightsEnabled set to true.",
						     "type": "boolean"
						   }
						*/
						Description: "Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be enabled on a group with InsightsEnabled set to true.",
						Type:        types.BoolType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "type": "string"
			         },
			         "Value": {
			           "type": "string"
			         }
			       },
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "$ref": "#/definitions/Tags",
			     "type": "array"
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
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
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
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
		Description: "This schema provides construct and validation rules for AWS-XRay Group resource parameters.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::XRay::Group").WithTerraformTypeName("aws_xray_group").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_xray_group", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
