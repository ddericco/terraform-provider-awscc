// Code generated by generators/resource/main.go; DO NOT EDIT.

package globalaccelerator

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
	registry.AddResourceTypeFactory("aws_globalaccelerator_endpoint_group", endpointGroup)
}

// endpointGroup returns the Terraform aws_globalaccelerator_endpoint_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::GlobalAccelerator::EndpointGroup resource type.
func endpointGroup(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"endpoint_configurations": {
			// Property: EndpointConfigurations
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The list of endpoint objects.",
			     "items": {
			       "description": "The configuration for a given endpoint",
			       "properties": {
			         "ClientIPPreservationEnabled": {
			           "description": "true if client ip should be preserved",
			           "type": "boolean"
			         },
			         "EndpointId": {
			           "description": "Id of the endpoint. For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID",
			           "type": "string"
			         },
			         "Weight": {
			           "description": "The weight for the endpoint.",
			           "type": "integer"
			         }
			       },
			       "$ref": "#/definitions/EndpointConfiguration",
			       "required": [
			         "EndpointId"
			       ],
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "The list of endpoint objects.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"client_ip_preservation_enabled": {
						// Property: ClientIPPreservationEnabled
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "true if client ip should be preserved",
						     "type": "boolean"
						   }
						*/
						Description: "true if client ip should be preserved",
						Type:        types.BoolType,
						Optional:    true,
					},
					"endpoint_id": {
						// Property: EndpointId
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Id of the endpoint. For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID",
						     "type": "string"
						   }
						*/
						Description: "Id of the endpoint. For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID",
						Type:        types.StringType,
						Required:    true,
					},
					"weight": {
						// Property: Weight
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The weight for the endpoint.",
						     "type": "integer"
						   }
						*/
						Description: "The weight for the endpoint.",
						Type:        types.NumberType,
						Optional:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"endpoint_group_arn": {
			// Property: EndpointGroupArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the endpoint group",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the endpoint group",
			Type:        types.StringType,
			Computed:    true,
		},
		"endpoint_group_region": {
			// Property: EndpointGroupRegion
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the AWS Region where the endpoint group is located",
			     "type": "string"
			   }
			*/
			Description: "The name of the AWS Region where the endpoint group is located",
			Type:        types.StringType,
			Required:    true,
			// EndpointGroupRegion is a force-new attribute.
		},
		"health_check_interval_seconds": {
			// Property: HealthCheckIntervalSeconds
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The time in seconds between each health check for an endpoint. Must be a value of 10 or 30",
			     "type": "integer"
			   }
			*/
			Description: "The time in seconds between each health check for an endpoint. Must be a value of 10 or 30",
			Type:        types.NumberType,
			Optional:    true,
		},
		"health_check_path": {
			// Property: HealthCheckPath
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "",
			     "type": "string"
			   }
			*/
			Description: "",
			Type:        types.StringType,
			Optional:    true,
		},
		"health_check_port": {
			// Property: HealthCheckPort
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
			     "type": "integer"
			   }
			*/
			Description: "The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"health_check_protocol": {
			// Property: HealthCheckProtocol
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
			     "enum": [
			       "TCP",
			       "HTTP",
			       "HTTPS"
			     ],
			     "type": "string"
			   }
			*/
			Description: "The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
			Type:        types.StringType,
			Optional:    true,
		},
		"listener_arn": {
			// Property: ListenerArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the listener",
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the listener",
			Type:        types.StringType,
			Required:    true,
			// ListenerArn is a force-new attribute.
		},
		"port_overrides": {
			// Property: PortOverrides
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "description": "listener to endpoint port mapping.",
			       "properties": {
			         "EndpointPort": {
			           "description": "A network port number",
			           "$ref": "#/definitions/Port",
			           "type": "integer"
			         },
			         "ListenerPort": {
			           "description": "A network port number",
			           "$ref": "#/definitions/Port",
			           "type": "integer"
			         }
			       },
			       "$ref": "#/definitions/PortOverride",
			       "required": [
			         "ListenerPort",
			         "EndpointPort"
			       ],
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"endpoint_port": {
						// Property: EndpointPort
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "A network port number",
						     "$ref": "#/definitions/Port",
						     "type": "integer"
						   }
						*/
						Description: "A network port number",
						Type:        types.NumberType,
						Required:    true,
					},
					"listener_port": {
						// Property: ListenerPort
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "A network port number",
						     "$ref": "#/definitions/Port",
						     "type": "integer"
						   }
						*/
						Description: "A network port number",
						Type:        types.NumberType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"threshold_count": {
			// Property: ThresholdCount
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The number of consecutive health checks required to set the state of the endpoint to unhealthy.",
			     "type": "integer"
			   }
			*/
			Description: "The number of consecutive health checks required to set the state of the endpoint to unhealthy.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"traffic_dial_percentage": {
			// Property: TrafficDialPercentage
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The percentage of traffic to sent to an AWS Region",
			     "type": "number"
			   }
			*/
			Description: "The percentage of traffic to sent to an AWS Region",
			Type:        types.NumberType,
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::GlobalAccelerator::EndpointGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GlobalAccelerator::EndpointGroup").WithTerraformTypeName("aws_globalaccelerator_endpoint_group").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_globalaccelerator_endpoint_group", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
