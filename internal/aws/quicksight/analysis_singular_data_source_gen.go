// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_quicksight_analysis", analysisDataSourceType)
}

// analysisDataSourceType returns the Terraform awscc_quicksight_analysis data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::QuickSight::Analysis resource type.
func analysisDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"analysis_id": {
			// Property: AnalysisId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the analysis.\u003c/p\u003e",
			//   "type": "string"
			// }
			Description: "<p>The Amazon Resource Name (ARN) of the analysis.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"aws_account_id": {
			// Property: AwsAccountId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 12,
			//   "minLength": 12,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe time that the analysis was created.\u003c/p\u003e",
			//   "format": "string",
			//   "type": "string"
			// }
			Description: "<p>The time that the analysis was created.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"data_set_arns": {
			// Property: DataSetArns
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe ARNs of the datasets of the analysis.\u003c/p\u003e",
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 100,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "<p>The ARNs of the datasets of the analysis.</p>",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"errors": {
			// Property: Errors
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eErrors associated with the analysis.\u003c/p\u003e",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "\u003cp\u003eA metadata error structure for an analysis.\u003c/p\u003e",
			//     "properties": {
			//       "Message": {
			//         "description": "\u003cp\u003eThe message associated with the analysis error.\u003c/p\u003e",
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Type": {
			//         "enum": [
			//           "ACCESS_DENIED",
			//           "SOURCE_NOT_FOUND",
			//           "DATA_SET_NOT_FOUND",
			//           "INTERNAL_FAILURE",
			//           "PARAMETER_VALUE_INCOMPATIBLE",
			//           "PARAMETER_TYPE_INVALID",
			//           "PARAMETER_NOT_FOUND",
			//           "COLUMN_TYPE_MISMATCH",
			//           "COLUMN_GEOGRAPHIC_ROLE_MISMATCH",
			//           "COLUMN_REPLACEMENT_MISSING"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "<p>Errors associated with the analysis.</p>",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"message": {
						// Property: Message
						Description: "<p>The message associated with the analysis error.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"last_updated_time": {
			// Property: LastUpdatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe time that the analysis was last updated.\u003c/p\u003e",
			//   "format": "string",
			//   "type": "string"
			// }
			Description: "<p>The time that the analysis was last updated.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe descriptive name of the analysis.\u003c/p\u003e",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "<p>The descriptive name of the analysis.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"parameters": {
			// Property: Parameters
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "\u003cp\u003eA list of QuickSight parameters and the list's override values.\u003c/p\u003e",
			//   "properties": {
			//     "DateTimeParameters": {
			//       "description": "\u003cp\u003eDate-time parameters.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eA date-time parameter.\u003c/p\u003e",
			//         "properties": {
			//           "Name": {
			//             "description": "\u003cp\u003eA display name for the date-time parameter.\u003c/p\u003e",
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Values": {
			//             "description": "\u003cp\u003eThe values for the date-time parameter.\u003c/p\u003e",
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "required": [
			//           "Name",
			//           "Values"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 100,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "DecimalParameters": {
			//       "description": "\u003cp\u003eDecimal parameters.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eA decimal parameter.\u003c/p\u003e",
			//         "properties": {
			//           "Name": {
			//             "description": "\u003cp\u003eA display name for the decimal parameter.\u003c/p\u003e",
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Values": {
			//             "description": "\u003cp\u003eThe values for the decimal parameter.\u003c/p\u003e",
			//             "items": {
			//               "type": "number"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "required": [
			//           "Name",
			//           "Values"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 100,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "IntegerParameters": {
			//       "description": "\u003cp\u003eInteger parameters.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eAn integer parameter.\u003c/p\u003e",
			//         "properties": {
			//           "Name": {
			//             "description": "\u003cp\u003eThe name of the integer parameter.\u003c/p\u003e",
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Values": {
			//             "description": "\u003cp\u003eThe values for the integer parameter.\u003c/p\u003e",
			//             "items": {
			//               "type": "number"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "required": [
			//           "Name",
			//           "Values"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 100,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "StringParameters": {
			//       "description": "\u003cp\u003eString parameters.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eA string parameter.\u003c/p\u003e",
			//         "properties": {
			//           "Name": {
			//             "description": "\u003cp\u003eA display name for a string parameter.\u003c/p\u003e",
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Values": {
			//             "description": "\u003cp\u003eThe values of a string parameter.\u003c/p\u003e",
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "required": [
			//           "Name",
			//           "Values"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 100,
			//       "minItems": 0,
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "<p>A list of QuickSight parameters and the list's override values.</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"date_time_parameters": {
						// Property: DateTimeParameters
						Description: "<p>Date-time parameters.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "<p>A display name for the date-time parameter.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"values": {
									// Property: Values
									Description: "<p>The values for the date-time parameter.</p>",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
					"decimal_parameters": {
						// Property: DecimalParameters
						Description: "<p>Decimal parameters.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "<p>A display name for the decimal parameter.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"values": {
									// Property: Values
									Description: "<p>The values for the decimal parameter.</p>",
									Type:        types.ListType{ElemType: types.NumberType},
									Computed:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
					"integer_parameters": {
						// Property: IntegerParameters
						Description: "<p>Integer parameters.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "<p>The name of the integer parameter.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"values": {
									// Property: Values
									Description: "<p>The values for the integer parameter.</p>",
									Type:        types.ListType{ElemType: types.NumberType},
									Computed:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
					"string_parameters": {
						// Property: StringParameters
						Description: "<p>String parameters.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "<p>A display name for a string parameter.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"values": {
									// Property: Values
									Description: "<p>The values of a string parameter.</p>",
									Type:        types.ListType{ElemType: types.StringType},
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
		"permissions": {
			// Property: Permissions
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eA structure that describes the principals and the resource-level permissions on an\n            analysis. You can use the \u003ccode\u003ePermissions\u003c/code\u003e structure to grant permissions by\n            providing a list of AWS Identity and Access Management (IAM) action information for each\n            principal listed by Amazon Resource Name (ARN). \u003c/p\u003e\n\n        \u003cp\u003eTo specify no permissions, omit \u003ccode\u003ePermissions\u003c/code\u003e.\u003c/p\u003e",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "\u003cp\u003ePermission for the resource.\u003c/p\u003e",
			//     "properties": {
			//       "Actions": {
			//         "description": "\u003cp\u003eThe IAM action to grant or revoke permissions on.\u003c/p\u003e",
			//         "items": {
			//           "type": "string"
			//         },
			//         "maxItems": 16,
			//         "minItems": 1,
			//         "type": "array"
			//       },
			//       "Principal": {
			//         "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:\u003c/p\u003e\n        \u003cul\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an AWS account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across AWS accounts.\n                    (This is less common.) \u003c/p\u003e\n            \u003c/li\u003e\n         \u003c/ul\u003e",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Actions",
			//       "Principal"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 64,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "<p>A structure that describes the principals and the resource-level permissions on an\n            analysis. You can use the <code>Permissions</code> structure to grant permissions by\n            providing a list of AWS Identity and Access Management (IAM) action information for each\n            principal listed by Amazon Resource Name (ARN). </p>\n\n        <p>To specify no permissions, omit <code>Permissions</code>.</p>",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"actions": {
						// Property: Actions
						Description: "<p>The IAM action to grant or revoke permissions on.</p>",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
					"principal": {
						// Property: Principal
						Description: "<p>The Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:</p>\n        <ul>\n            <li>\n                <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>\n            </li>\n            <li>\n                <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>\n            </li>\n            <li>\n                <p>The ARN of an AWS account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across AWS accounts.\n                    (This is less common.) </p>\n            </li>\n         </ul>",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"sheets": {
			// Property: Sheets
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eA list of the associated sheets with the unique identifier and name of each sheet.\u003c/p\u003e",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "\u003cp\u003eA \u003ci\u003esheet\u003c/i\u003e, which is an object that contains a set of visuals that\n            are viewed together on one page in the Amazon QuickSight console. Every analysis and dashboard\n            contains at least one sheet. Each sheet contains at least one visualization widget, for\n            example a chart, pivot table, or narrative insight. Sheets can be associated with other\n            components, such as controls, filters, and so on.\u003c/p\u003e",
			//     "properties": {
			//       "Name": {
			//         "description": "\u003cp\u003eThe name of a sheet. This name is displayed on the sheet's tab in the QuickSight\n            console.\u003c/p\u003e",
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "SheetId": {
			//         "description": "\u003cp\u003eThe unique identifier associated with a sheet.\u003c/p\u003e",
			//         "maxLength": 2048,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 20,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "<p>A list of the associated sheets with the unique identifier and name of each sheet.</p>",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Description: "<p>The name of a sheet. This name is displayed on the sheet's tab in the QuickSight\n            console.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
					"sheet_id": {
						// Property: SheetId
						Description: "<p>The unique identifier associated with a sheet.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"source_entity": {
			// Property: SourceEntity
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "\u003cp\u003eThe source entity of an analysis.\u003c/p\u003e",
			//   "properties": {
			//     "SourceTemplate": {
			//       "additionalProperties": false,
			//       "description": "\u003cp\u003eThe source template of an analysis.\u003c/p\u003e",
			//       "properties": {
			//         "Arn": {
			//           "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the source template of an analysis.\u003c/p\u003e",
			//           "type": "string"
			//         },
			//         "DataSetReferences": {
			//           "description": "\u003cp\u003eThe dataset references of the source template of an analysis.\u003c/p\u003e",
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "\u003cp\u003eDataset reference.\u003c/p\u003e",
			//             "properties": {
			//               "DataSetArn": {
			//                 "description": "\u003cp\u003eDataset Amazon Resource Name (ARN).\u003c/p\u003e",
			//                 "type": "string"
			//               },
			//               "DataSetPlaceholder": {
			//                 "description": "\u003cp\u003eDataset placeholder.\u003c/p\u003e",
			//                 "pattern": "",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "DataSetArn",
			//               "DataSetPlaceholder"
			//             ],
			//             "type": "object"
			//           },
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "required": [
			//         "Arn",
			//         "DataSetReferences"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "<p>The source entity of an analysis.</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"source_template": {
						// Property: SourceTemplate
						Description: "<p>The source template of an analysis.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"arn": {
									// Property: Arn
									Description: "<p>The Amazon Resource Name (ARN) of the source template of an analysis.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"data_set_references": {
									// Property: DataSetReferences
									Description: "<p>The dataset references of the source template of an analysis.</p>",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"data_set_arn": {
												// Property: DataSetArn
												Description: "<p>Dataset Amazon Resource Name (ARN).</p>",
												Type:        types.StringType,
												Computed:    true,
											},
											"data_set_placeholder": {
												// Property: DataSetPlaceholder
												Description: "<p>Dataset placeholder.</p>",
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
				},
			),
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "CREATION_IN_PROGRESS",
			//     "CREATION_SUCCESSFUL",
			//     "CREATION_FAILED",
			//     "UPDATE_IN_PROGRESS",
			//     "UPDATE_SUCCESSFUL",
			//     "UPDATE_FAILED",
			//     "DELETED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eContains a map of the key-value pairs for the resource tag or tags assigned to the\n            analysis.\u003c/p\u003e",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "\u003cp\u003eThe key or keys of the key-value pairs for the resource tag or tags assigned to the\n            resource.\u003c/p\u003e",
			//     "properties": {
			//       "Key": {
			//         "description": "\u003cp\u003eTag key.\u003c/p\u003e",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "\u003cp\u003eTag value.\u003c/p\u003e",
			//         "maxLength": 256,
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
			//   "maxItems": 200,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "<p>Contains a map of the key-value pairs for the resource tag or tags assigned to the\n            analysis.</p>",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "<p>Tag key.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "<p>Tag value.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"theme_arn": {
			// Property: ThemeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe ARN of the theme of the analysis.\u003c/p\u003e",
			//   "type": "string"
			// }
			Description: "<p>The ARN of the theme of the analysis.</p>",
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
		Description: "Data Source schema for AWS::QuickSight::Analysis",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::QuickSight::Analysis").WithTerraformTypeName("awscc_quicksight_analysis")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":              "Actions",
		"analysis_id":          "AnalysisId",
		"arn":                  "Arn",
		"aws_account_id":       "AwsAccountId",
		"created_time":         "CreatedTime",
		"data_set_arn":         "DataSetArn",
		"data_set_arns":        "DataSetArns",
		"data_set_placeholder": "DataSetPlaceholder",
		"data_set_references":  "DataSetReferences",
		"date_time_parameters": "DateTimeParameters",
		"decimal_parameters":   "DecimalParameters",
		"errors":               "Errors",
		"integer_parameters":   "IntegerParameters",
		"key":                  "Key",
		"last_updated_time":    "LastUpdatedTime",
		"message":              "Message",
		"name":                 "Name",
		"parameters":           "Parameters",
		"permissions":          "Permissions",
		"principal":            "Principal",
		"sheet_id":             "SheetId",
		"sheets":               "Sheets",
		"source_entity":        "SourceEntity",
		"source_template":      "SourceTemplate",
		"status":               "Status",
		"string_parameters":    "StringParameters",
		"tags":                 "Tags",
		"theme_arn":            "ThemeArn",
		"type":                 "Type",
		"value":                "Value",
		"values":               "Values",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
