// Code generated by generators/resource/main.go; DO NOT EDIT.

package lambda

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_lambda_event_source_mapping", eventSourceMappingResourceType)
}

// eventSourceMappingResourceType returns the Terraform awscc_lambda_event_source_mapping resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Lambda::EventSourceMapping resource type.
func eventSourceMappingResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"batch_size": {
			// Property: BatchSize
			// CloudFormation resource type schema:
			// {
			//   "description": "The maximum number of items to retrieve in a single batch.",
			//   "type": "integer"
			// }
			Description: "The maximum number of items to retrieve in a single batch.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"bisect_batch_on_function_error": {
			// Property: BisectBatchOnFunctionError
			// CloudFormation resource type schema:
			// {
			//   "description": "(Streams) If the function returns an error, split the batch in two and retry.",
			//   "type": "boolean"
			// }
			Description: "(Streams) If the function returns an error, split the batch in two and retry.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"destination_config": {
			// Property: DestinationConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "(Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.",
			//   "properties": {
			//     "OnFailure": {
			//       "additionalProperties": false,
			//       "description": "A destination for events that failed processing.",
			//       "properties": {
			//         "Destination": {
			//           "description": "The Amazon Resource Name (ARN) of the destination resource.",
			//           "maxLength": 1024,
			//           "minLength": 12,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "(Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"on_failure": {
						// Property: OnFailure
						Description: "A destination for events that failed processing.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"destination": {
									// Property: Destination
									Description: "The Amazon Resource Name (ARN) of the destination resource.",
									Type:        types.StringType,
									Validators:  []tfsdk.AttributeValidator{validate.StringLength(12, 1024)},
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Disables the event source mapping to pause polling and invocation.",
			//   "type": "boolean"
			// }
			Description: "Disables the event source mapping to pause polling and invocation.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"event_source_arn": {
			// Property: EventSourceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the event source.",
			//   "maxLength": 1024,
			//   "minLength": 12,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the event source.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(12, 1024)},
			Optional:    true,
			Computed:    true,
			// EventSourceArn is a force-new attribute.
		},
		"function_name": {
			// Property: FunctionName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Lambda function.",
			//   "maxLength": 140,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the Lambda function.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 140)},
			Required:    true,
		},
		"function_response_types": {
			// Property: FunctionResponseTypes
			// CloudFormation resource type schema:
			// {
			//   "description": "(Streams) A list of response types supported by the function.",
			//   "items": {
			//     "enum": [
			//       "ReportBatchItemFailures"
			//     ],
			//     "type": "string"
			//   },
			//   "maxLength": 1,
			//   "minLength": 0,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "(Streams) A list of response types supported by the function.",
			Type:        types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			Optional: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Event Source Mapping Identifier UUID.",
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Event Source Mapping Identifier UUID.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(36, 36)},
			Computed:    true,
		},
		"maximum_batching_window_in_seconds": {
			// Property: MaximumBatchingWindowInSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "(Streams) The maximum amount of time to gather records before invoking the function, in seconds.",
			//   "type": "integer"
			// }
			Description: "(Streams) The maximum amount of time to gather records before invoking the function, in seconds.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"maximum_record_age_in_seconds": {
			// Property: MaximumRecordAgeInSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "(Streams) The maximum age of a record that Lambda sends to a function for processing.",
			//   "type": "integer"
			// }
			Description: "(Streams) The maximum age of a record that Lambda sends to a function for processing.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"maximum_retry_attempts": {
			// Property: MaximumRetryAttempts
			// CloudFormation resource type schema:
			// {
			//   "description": "(Streams) The maximum number of times to retry when the function returns an error.",
			//   "type": "integer"
			// }
			Description: "(Streams) The maximum number of times to retry when the function returns an error.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"parallelization_factor": {
			// Property: ParallelizationFactor
			// CloudFormation resource type schema:
			// {
			//   "description": "(Streams) The number of batches to process from each shard concurrently.",
			//   "type": "integer"
			// }
			Description: "(Streams) The number of batches to process from each shard concurrently.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"queues": {
			// Property: Queues
			// CloudFormation resource type schema:
			// {
			//   "description": "(ActiveMQ) A list of ActiveMQ queues.",
			//   "items": {
			//     "maxLength": 1000,
			//     "minLength": 1,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "maxItems": 1,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "(ActiveMQ) A list of ActiveMQ queues.",
			Type:        types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLength(1, 1),
				validate.UniqueItems(),
			},
			Optional: true,
		},
		"self_managed_event_source": {
			// Property: SelfManagedEventSource
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The configuration used by AWS Lambda to access a self-managed event source.",
			//   "properties": {
			//     "Endpoints": {
			//       "additionalProperties": false,
			//       "description": "The endpoints used by AWS Lambda to access a self-managed event source.",
			//       "properties": {
			//         "KafkaBootstrapServers": {
			//           "description": "A list of Kafka server endpoints.",
			//           "items": {
			//             "description": "The URL of a Kafka server.",
			//             "maxLength": 300,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "maxItems": 10,
			//           "minItems": 1,
			//           "type": "array",
			//           "uniqueItems": true
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The configuration used by AWS Lambda to access a self-managed event source.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"endpoints": {
						// Property: Endpoints
						Description: "The endpoints used by AWS Lambda to access a self-managed event source.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"kafka_bootstrap_servers": {
									// Property: KafkaBootstrapServers
									Description: "A list of Kafka server endpoints.",
									Type:        types.ListType{ElemType: types.StringType},
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLength(1, 10),
										validate.UniqueItems(),
									},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// SelfManagedEventSource is a force-new attribute.
		},
		"source_access_configurations": {
			// Property: SourceAccessConfigurations
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of SourceAccessConfiguration.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The configuration used by AWS Lambda to access event source",
			//     "properties": {
			//       "Type": {
			//         "description": "The type of source access configuration.",
			//         "enum": [
			//           "BASIC_AUTH",
			//           "VPC_SUBNET",
			//           "VPC_SECURITY_GROUP",
			//           "SASL_SCRAM_512_AUTH",
			//           "SASL_SCRAM_256_AUTH",
			//           "VIRTUAL_HOST"
			//         ],
			//         "type": "string"
			//       },
			//       "URI": {
			//         "description": "The URI for the source access configuration resource.",
			//         "maxLength": 200,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 22,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of SourceAccessConfiguration.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"type": {
						// Property: Type
						Description: "The type of source access configuration.",
						Type:        types.StringType,
						Optional:    true,
					},
					"uri": {
						// Property: URI
						Description: "The URI for the source access configuration resource.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 200)},
						Optional:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 1,
					MaxItems: 22,
				},
			),
			Validators: []tfsdk.AttributeValidator{validate.UniqueItems()},
			Optional:   true,
		},
		"starting_position": {
			// Property: StartingPosition
			// CloudFormation resource type schema:
			// {
			//   "description": "The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB Streams sources.",
			//   "maxLength": 12,
			//   "minLength": 6,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB Streams sources.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(6, 12)},
			Optional:    true,
			Computed:    true,
			// StartingPosition is a force-new attribute.
		},
		"starting_position_timestamp": {
			// Property: StartingPositionTimestamp
			// CloudFormation resource type schema:
			// {
			//   "description": "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.",
			//   "type": "number"
			// }
			Description: "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"topics": {
			// Property: Topics
			// CloudFormation resource type schema:
			// {
			//   "description": "(Kafka) A list of Kafka topics.",
			//   "items": {
			//     "maxLength": 249,
			//     "minLength": 1,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "maxItems": 1,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "(Kafka) A list of Kafka topics.",
			Type:        types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLength(1, 1),
				validate.UniqueItems(),
			},
			Optional: true,
		},
		"tumbling_window_in_seconds": {
			// Property: TumblingWindowInSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "(Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.",
			//   "type": "integer"
			// }
			Description: "(Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.",
			Type:        types.NumberType,
			Optional:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Lambda::EventSourceMapping",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::EventSourceMapping").WithTerraformTypeName("awscc_lambda_event_source_mapping")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"batch_size":                         "BatchSize",
		"bisect_batch_on_function_error":     "BisectBatchOnFunctionError",
		"destination":                        "Destination",
		"destination_config":                 "DestinationConfig",
		"enabled":                            "Enabled",
		"endpoints":                          "Endpoints",
		"event_source_arn":                   "EventSourceArn",
		"function_name":                      "FunctionName",
		"function_response_types":            "FunctionResponseTypes",
		"id":                                 "Id",
		"kafka_bootstrap_servers":            "KafkaBootstrapServers",
		"maximum_batching_window_in_seconds": "MaximumBatchingWindowInSeconds",
		"maximum_record_age_in_seconds":      "MaximumRecordAgeInSeconds",
		"maximum_retry_attempts":             "MaximumRetryAttempts",
		"on_failure":                         "OnFailure",
		"parallelization_factor":             "ParallelizationFactor",
		"queues":                             "Queues",
		"self_managed_event_source":          "SelfManagedEventSource",
		"source_access_configurations":       "SourceAccessConfigurations",
		"starting_position":                  "StartingPosition",
		"starting_position_timestamp":        "StartingPositionTimestamp",
		"topics":                             "Topics",
		"tumbling_window_in_seconds":         "TumblingWindowInSeconds",
		"type":                               "Type",
		"uri":                                "URI",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_lambda_event_source_mapping", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
