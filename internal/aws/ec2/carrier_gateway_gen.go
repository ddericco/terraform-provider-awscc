// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ec2_carrier_gateway", carrierGatewayResourceType)
}

// carrierGatewayResourceType returns the Terraform awscc_ec2_carrier_gateway resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::CarrierGateway resource type.
func carrierGatewayResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"carrier_gateway_id": {
			// Property: CarrierGatewayId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the carrier gateway.",
			//   "type": "string"
			// }
			Description: "The ID of the carrier gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"owner_id": {
			// Property: OwnerId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the owner.",
			//   "type": "string"
			// }
			Description: "The ID of the owner.",
			Type:        types.StringType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of the carrier gateway.",
			//   "type": "string"
			// }
			Description: "The state of the carrier gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 127)},
						Optional:   true,
					},
					"value": {
						// Property: Value
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 255)},
						Optional:   true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the VPC.",
			//   "type": "string"
			// }
			Description: "The ID of the VPC.",
			Type:        types.StringType,
			Required:    true,
			// VpcId is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::CarrierGateway").WithTerraformTypeName("awscc_ec2_carrier_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"carrier_gateway_id": "CarrierGatewayId",
		"key":                "Key",
		"owner_id":           "OwnerId",
		"state":              "State",
		"tags":               "Tags",
		"value":              "Value",
		"vpc_id":             "VpcId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ec2_carrier_gateway", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
