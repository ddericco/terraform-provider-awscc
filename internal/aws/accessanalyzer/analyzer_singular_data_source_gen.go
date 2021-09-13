// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package accessanalyzer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_accessanalyzer_analyzer", analyzerDataSourceType)
}

// analyzerDataSourceType returns the Terraform awscc_accessanalyzer_analyzer data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::AccessAnalyzer::Analyzer resource type.
func analyzerDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"analyzer_name": {
			// Property: AnalyzerName
			// CloudFormation resource type schema:
			// {
			//   "description": "Analyzer name",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Analyzer name",
			Type:        types.StringType,
			Computed:    true,
		},
		"archive_rules": {
			// Property: ArchiveRules
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "description": "An Access Analyzer archive rule. Archive rules automatically archive new findings that meet the criteria you define when you create the rule.",
			//     "properties": {
			//       "Filter": {
			//         "insertionOrder": false,
			//         "items": {
			//           "properties": {
			//             "Contains": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             },
			//             "Eq": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             },
			//             "Exists": {
			//               "type": "boolean"
			//             },
			//             "Neq": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             },
			//             "Property": {
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Property"
			//           ],
			//           "type": "object"
			//         },
			//         "minItems": 1,
			//         "type": "array"
			//       },
			//       "RuleName": {
			//         "description": "The archive rule name",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Filter",
			//       "RuleName"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"filter": {
						// Property: Filter
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"contains": {
									// Property: Contains
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"eq": {
									// Property: Eq
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"exists": {
									// Property: Exists
									Type:     types.BoolType,
									Computed: true,
								},
								"neq": {
									// Property: Neq
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"property": {
									// Property: Property
									Type:     types.StringType,
									Computed: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
					"rule_name": {
						// Property: RuleName
						Description: "The archive rule name",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the analyzer",
			//   "maxLength": 1600,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the analyzer",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 255,
			//         "minLength": 1,
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
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of the analyzer, must be ACCOUNT or ORGANIZATION",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The type of the analyzer, must be ACCOUNT or ORGANIZATION",
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
		Description: "Data Source schema for AWS::AccessAnalyzer::Analyzer",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AccessAnalyzer::Analyzer").WithTerraformTypeName("awscc_accessanalyzer_analyzer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"analyzer_name": "AnalyzerName",
		"archive_rules": "ArchiveRules",
		"arn":           "Arn",
		"contains":      "Contains",
		"eq":            "Eq",
		"exists":        "Exists",
		"filter":        "Filter",
		"key":           "Key",
		"neq":           "Neq",
		"property":      "Property",
		"rule_name":     "RuleName",
		"tags":          "Tags",
		"type":          "Type",
		"value":         "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
