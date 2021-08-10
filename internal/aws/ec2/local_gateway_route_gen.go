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
	registry.AddResourceTypeFactory("aws_ec2_local_gateway_route", localGatewayRoute)
}

// localGatewayRoute returns the Terraform aws_ec2_local_gateway_route resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::LocalGatewayRoute resource type.
func localGatewayRoute(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"destination_cidr_block": {
			// Property: DestinationCidrBlock
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The CIDR block used for destination matches.",
			     "type": "string"
			   }
			*/
			Description: "The CIDR block used for destination matches.",
			Type:        types.StringType,
			Required:    true,
			// DestinationCidrBlock is a force-new attribute.
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
		"local_gateway_virtual_interface_group_id": {
			// Property: LocalGatewayVirtualInterfaceGroupId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the virtual interface group.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the virtual interface group.",
			Type:        types.StringType,
			Required:    true,
			// LocalGatewayVirtualInterfaceGroupId is a force-new attribute.
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The state of the route.",
			     "type": "string"
			   }
			*/
			Description: "The state of the route.",
			Type:        types.StringType,
			Computed:    true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The route type.",
			     "type": "string"
			   }
			*/
			Description: "The route type.",
			Type:        types.StringType,
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
		Description: "Describes a route for a local gateway route table.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::LocalGatewayRoute").WithTerraformTypeName("aws_ec2_local_gateway_route").WithTerraformSchema(schema)

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_ec2_local_gateway_route", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
