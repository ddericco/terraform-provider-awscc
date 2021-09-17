// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotevents

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iotevents_input", inputDataSourceType)
}

// inputDataSourceType returns the Terraform awscc_iotevents_input data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoTEvents::Input resource type.
func inputDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"input_definition": {
			// Property: InputDefinition
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The definition of the input.",
			//   "properties": {
			//     "Attributes": {
			//       "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
			//         "properties": {
			//           "JsonPath": {
			//             "description": "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events (`BatchPutMessage`). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.\n\n_Syntax_: `\u003cfield-name\u003e.\u003cfield-name\u003e...`",
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "JsonPath"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 200,
			//       "minItems": 1,
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "required": [
			//     "Attributes"
			//   ],
			//   "type": "object"
			// }
			Description: "The definition of the input.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"attributes": {
						// Property: Attributes
						Description: "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"json_path": {
									// Property: JsonPath
									Description: "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events (`BatchPutMessage`). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.\n\n_Syntax_: `<field-name>.<field-name>...`",
									Type:        types.StringType,
									Computed:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"input_description": {
			// Property: InputDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A brief description of the input.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A brief description of the input.",
			Type:        types.StringType,
			Computed:    true,
		},
		"input_name": {
			// Property: InputName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the input.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the input.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.\n\nFor more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Tags to be applied to Input.",
			//     "properties": {
			//       "Key": {
			//         "description": "Key of the Tag.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Value of the Tag.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An array of key-value pairs to apply to this resource.\n\nFor more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "Key of the Tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "Value of the Tag.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoTEvents::Input",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTEvents::Input").WithTerraformTypeName("awscc_iotevents_input")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attributes":        "Attributes",
		"input_definition":  "InputDefinition",
		"input_description": "InputDescription",
		"input_name":        "InputName",
		"json_path":         "JsonPath",
		"key":               "Key",
		"tags":              "Tags",
		"value":             "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
