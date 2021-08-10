// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkfirewall

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
	registry.AddResourceTypeFactory("aws_networkfirewall_logging_configuration", loggingConfiguration)
}

// loggingConfiguration returns the Terraform aws_networkfirewall_logging_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkFirewall::LoggingConfiguration resource type.
func loggingConfiguration(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"firewall_arn": {
			// Property: FirewallArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A resource ARN.",
			     "maxLength": 256,
			     "minLength": 1,
			     "pattern": "",
			     "$ref": "#/definitions/ResourceArn",
			     "type": "string"
			   }
			*/
			Description: "A resource ARN.",
			Type:        types.StringType,
			Required:    true,
			// FirewallArn is a force-new attribute.
		},
		"firewall_name": {
			// Property: FirewallName
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 128,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// FirewallName is a force-new attribute.
		},
		"logging_configuration": {
			// Property: LoggingConfiguration
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "LogDestinationConfigs": {
			         "insertionOrder": false,
			         "items": {
			           "additionalProperties": false,
			           "properties": {
			             "LogDestination": {
			               "additionalProperties": false,
			               "description": "A key-value pair to configure the logDestinations.",
			               "minItems": 1,
			               "patternProperties": {
			                 "": {
			                   "maxLength": 1024,
			                   "minLength": 1,
			                   "type": "string"
			                 }
			               },
			               "type": "object"
			             },
			             "LogDestinationType": {
			               "enum": [
			                 "S3",
			                 "CloudWatchLogs",
			                 "KinesisDataFirehose"
			               ],
			               "type": "string"
			             },
			             "LogType": {
			               "enum": [
			                 "ALERT",
			                 "FLOW"
			               ],
			               "type": "string"
			             }
			           },
			           "$ref": "#/definitions/LogDestinationConfig",
			           "required": [
			             "LogType",
			             "LogDestinationType",
			             "LogDestination"
			           ],
			           "type": "object"
			         },
			         "minItems": 1,
			         "type": "array"
			       }
			     },
			     "$ref": "#/definitions/LoggingConfiguration",
			     "required": [
			       "LogDestinationConfigs"
			     ],
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"log_destination_configs": {
						// Property: LogDestinationConfigs
						// CloudFormation resource type schema:
						/*
						   {
						     "insertionOrder": false,
						     "items": {
						       "additionalProperties": false,
						       "properties": {
						         "LogDestination": {
						           "additionalProperties": false,
						           "description": "A key-value pair to configure the logDestinations.",
						           "minItems": 1,
						           "patternProperties": {
						             "": {
						               "maxLength": 1024,
						               "minLength": 1,
						               "type": "string"
						             }
						           },
						           "type": "object"
						         },
						         "LogDestinationType": {
						           "enum": [
						             "S3",
						             "CloudWatchLogs",
						             "KinesisDataFirehose"
						           ],
						           "type": "string"
						         },
						         "LogType": {
						           "enum": [
						             "ALERT",
						             "FLOW"
						           ],
						           "type": "string"
						         }
						       },
						       "$ref": "#/definitions/LogDestinationConfig",
						       "required": [
						         "LogType",
						         "LogDestinationType",
						         "LogDestination"
						       ],
						       "type": "object"
						     },
						     "minItems": 1,
						     "type": "array"
						   }
						*/
						// Multiset.
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"log_destination": {
									// Property: LogDestination
									// CloudFormation resource type schema:
									/*
									   {
									     "additionalProperties": false,
									     "description": "A key-value pair to configure the logDestinations.",
									     "minItems": 1,
									     "patternProperties": {
									       "": {
									         "maxLength": 1024,
									         "minLength": 1,
									         "type": "string"
									       }
									     },
									     "type": "object"
									   }
									*/
									Description: "A key-value pair to configure the logDestinations.",
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Required: true,
								},
								"log_destination_type": {
									// Property: LogDestinationType
									// CloudFormation resource type schema:
									/*
									   {
									     "enum": [
									       "S3",
									       "CloudWatchLogs",
									       "KinesisDataFirehose"
									     ],
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"log_type": {
									// Property: LogType
									// CloudFormation resource type schema:
									/*
									   {
									     "enum": [
									       "ALERT",
									       "FLOW"
									     ],
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{
								MinItems: 1,
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource type definition for AWS::NetworkFirewall::LoggingConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkFirewall::LoggingConfiguration").WithTerraformTypeName("aws_networkfirewall_logging_configuration").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_networkfirewall_logging_configuration", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
