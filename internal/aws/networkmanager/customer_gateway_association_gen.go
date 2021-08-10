// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkmanager

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
	registry.AddResourceTypeFactory("aws_networkmanager_customer_gateway_association", customerGatewayAssociation)
}

// customerGatewayAssociation returns the Terraform aws_networkmanager_customer_gateway_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkManager::CustomerGatewayAssociation resource type.
func customerGatewayAssociation(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"customer_gateway_arn": {
			// Property: CustomerGatewayArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the customer gateway.",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the customer gateway.",
			Type:        types.StringType,
			Required:    true,
			// CustomerGatewayArn is a force-new attribute.
		},
		"device_id": {
			// Property: DeviceId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the device",
			     "type": "string"
			   }
			*/
			Description: "The ID of the device",
			Type:        types.StringType,
			Required:    true,
			// DeviceId is a force-new attribute.
		},
		"global_network_id": {
			// Property: GlobalNetworkId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the global network.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the global network.",
			Type:        types.StringType,
			Required:    true,
			// GlobalNetworkId is a force-new attribute.
		},
		"link_id": {
			// Property: LinkId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the link",
			     "type": "string"
			   }
			*/
			Description: "The ID of the link",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// LinkId is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::NetworkManager::CustomerGatewayAssociation type associates a customer gateway with a device and optionally, with a link.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::CustomerGatewayAssociation").WithTerraformTypeName("aws_networkmanager_customer_gateway_association").WithTerraformSchema(schema)

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_networkmanager_customer_gateway_association", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
