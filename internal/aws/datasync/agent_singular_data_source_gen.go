// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_datasync_agent", agentDataSourceType)
}

// agentDataSourceType returns the Terraform awscc_datasync_agent data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::DataSync::Agent resource type.
func agentDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"activation_key": {
			// Property: ActivationKey
			// CloudFormation resource type schema:
			// {
			//   "description": "Activation key of the Agent.",
			//   "maxLength": 29,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Activation key of the Agent.",
			Type:        types.StringType,
			Computed:    true,
		},
		"agent_arn": {
			// Property: AgentArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The DataSync Agent ARN.",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The DataSync Agent ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"agent_name": {
			// Property: AgentName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name configured for the agent. Text reference used to identify the agent in the console.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name configured for the agent. Text reference used to identify the agent in the console.",
			Type:        types.StringType,
			Computed:    true,
		},
		"endpoint_type": {
			// Property: EndpointType
			// CloudFormation resource type schema:
			// {
			//   "description": "The service endpoints that the agent will connect to.",
			//   "enum": [
			//     "FIPS",
			//     "PUBLIC",
			//     "PRIVATE_LINK"
			//   ],
			//   "type": "string"
			// }
			Description: "The service endpoints that the agent will connect to.",
			Type:        types.StringType,
			Computed:    true,
		},
		"security_group_arns": {
			// Property: SecurityGroupArns
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARNs of the security group used to protect your data transfer task subnets.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 128,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The ARNs of the security group used to protect your data transfer task subnets.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"subnet_arns": {
			// Property: SubnetArns
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARNs of the subnets in which DataSync will create elastic network interfaces for each data transfer task.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 128,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The ARNs of the subnets in which DataSync will create elastic network interfaces for each data transfer task.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"vpc_endpoint_id": {
			// Property: VpcEndpointId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the VPC endpoint that the agent has access to.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the VPC endpoint that the agent has access to.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::DataSync::Agent",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::Agent").WithTerraformTypeName("awscc_datasync_agent")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"activation_key":      "ActivationKey",
		"agent_arn":           "AgentArn",
		"agent_name":          "AgentName",
		"endpoint_type":       "EndpointType",
		"key":                 "Key",
		"security_group_arns": "SecurityGroupArns",
		"subnet_arns":         "SubnetArns",
		"tags":                "Tags",
		"value":               "Value",
		"vpc_endpoint_id":     "VpcEndpointId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
