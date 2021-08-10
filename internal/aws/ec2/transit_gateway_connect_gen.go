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
	registry.AddResourceTypeFactory("aws_ec2_transit_gateway_connect", transitGatewayConnect)
}

// transitGatewayConnect returns the Terraform aws_ec2_transit_gateway_connect resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::TransitGatewayConnect resource type.
func transitGatewayConnect(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The creation time.",
			     "type": "string"
			   }
			*/
			Description: "The creation time.",
			Type:        types.StringType,
			Computed:    true,
		},
		"options": {
			// Property: Options
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "Protocol": {
			         "description": "The tunnel protocol.",
			         "type": "string"
			       }
			     },
			     "$ref": "#/definitions/TransitGatewayConnectOptions",
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"protocol": {
						// Property: Protocol
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The tunnel protocol.",
						     "type": "string"
						   }
						*/
						Description: "The tunnel protocol.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Required: true,
			// Options is a force-new attribute.
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The state of the attachment.",
			     "type": "string"
			   }
			*/
			Description: "The state of the attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The tags for the attachment.",
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
			           "type": "string"
			         },
			         "Value": {
			           "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "The tags for the attachment.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
						     "type": "string"
						   }
						*/
						Description: "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
						Type:        types.StringType,
						Optional:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
						     "type": "string"
						   }
						*/
						Description: "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"transit_gateway_attachment_id": {
			// Property: TransitGatewayAttachmentId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the Connect attachment.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the Connect attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"transit_gateway_id": {
			// Property: TransitGatewayId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the transit gateway.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the transit gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"transport_transit_gateway_attachment_id": {
			// Property: TransportTransitGatewayAttachmentId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the attachment from which the Connect attachment was created.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the attachment from which the Connect attachment was created.",
			Type:        types.StringType,
			Required:    true,
			// TransportTransitGatewayAttachmentId is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::EC2::TransitGatewayConnect type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayConnect").WithTerraformTypeName("aws_ec2_transit_gateway_connect").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_ec2_transit_gateway_connect", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
