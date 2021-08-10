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
	providertypes "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("aws_ec2_local_gateway_route_table_vpc_association", localGatewayRouteTableVPCAssociation)
}

// localGatewayRouteTableVPCAssociation returns the Terraform aws_ec2_local_gateway_route_table_vpc_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::LocalGatewayRouteTableVPCAssociation resource type.
func localGatewayRouteTableVPCAssociation(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"local_gateway_id": {
			// Property: LocalGatewayId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the local gateway.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the local gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"local_gateway_route_table_id": {
			// Property: LocalGatewayRouteTableId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the local gateway route table.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the local gateway route table.",
			Type:        types.StringType,
			Required:    true,
			// LocalGatewayRouteTableId is a force-new attribute.
		},
		"local_gateway_route_table_vpc_association_id": {
			// Property: LocalGatewayRouteTableVpcAssociationId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the association.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the association.",
			Type:        types.StringType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The state of the association.",
			     "type": "string"
			   }
			*/
			Description: "The state of the association.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "maxLength": 127,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "Value": {
			           "maxLength": 255,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "type": "object"
			     },
			     "$ref": "#/definitions/Tags",
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Attributes: providertypes.SetNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 127,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 255,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the VPC.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the VPC.",
			Type:        types.StringType,
			Required:    true,
			// VpcId is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Describes an association between a local gateway route table and a VPC.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::LocalGatewayRouteTableVPCAssociation").WithTerraformTypeName("aws_ec2_local_gateway_route_table_vpc_association").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_ec2_local_gateway_route_table_vpc_association", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
