// Code generated by generators/resource/main.go; DO NOT EDIT.

package frauddetector

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
	registry.AddResourceTypeFactory("awscc_frauddetector_detector", detectorResourceType)
}

// detectorResourceType returns the Terraform awscc_frauddetector_detector resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::FraudDetector::Detector resource type.
func detectorResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the detector.",
			//   "type": "string"
			// }
			Description: "The ARN of the detector.",
			Type:        types.StringType,
			Computed:    true,
		},
		"associated_models": {
			// Property: AssociatedModels
			// CloudFormation resource type schema:
			// {
			//   "description": "The models to associate with this detector.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A model to associate with a detector.",
			//     "properties": {
			//       "Arn": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 10,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "The models to associate with this detector.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MaxItems: 10,
				},
			),
			Optional: true,
		},
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time when the detector was created.",
			//   "type": "string"
			// }
			Description: "The time when the detector was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the detector.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The description of the detector.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
			Optional:    true,
		},
		"detector_id": {
			// Property: DetectorId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the detector",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the detector",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 64)},
			Required:    true,
			// DetectorId is a force-new attribute.
		},
		"detector_version_id": {
			// Property: DetectorVersionId
			// CloudFormation resource type schema:
			// {
			//   "description": "The active version ID of the detector",
			//   "type": "string"
			// }
			Description: "The active version ID of the detector",
			Type:        types.StringType,
			Computed:    true,
		},
		"detector_version_status": {
			// Property: DetectorVersionStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "The desired detector version status for the detector",
			//   "enum": [
			//     "DRAFT",
			//     "ACTIVE"
			//   ],
			//   "type": "string"
			// }
			Description: "The desired detector version status for the detector",
			Type:        types.StringType,
			Optional:    true,
		},
		"event_type": {
			// Property: EventType
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Arn": {
			//       "description": "The ARN of the event type.",
			//       "type": "string"
			//     },
			//     "CreatedTime": {
			//       "description": "The time when the event type was created.",
			//       "type": "string"
			//     },
			//     "Description": {
			//       "description": "The description of the event type.",
			//       "maxLength": 128,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "EntityTypes": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Arn": {
			//             "type": "string"
			//           },
			//           "CreatedTime": {
			//             "description": "The time when the entity type was created.",
			//             "type": "string"
			//           },
			//           "Description": {
			//             "description": "The description.",
			//             "maxLength": 256,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Inline": {
			//             "type": "boolean"
			//           },
			//           "LastUpdatedTime": {
			//             "description": "The time when the entity type was last updated.",
			//             "type": "string"
			//           },
			//           "Name": {
			//             "type": "string"
			//           },
			//           "Tags": {
			//             "description": "Tags associated with this entity type.",
			//             "insertionOrder": false,
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Key": {
			//                   "maxLength": 128,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "maxLength": 256,
			//                   "minLength": 0,
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Key",
			//                 "Value"
			//               ],
			//               "type": "object"
			//             },
			//             "maxItems": 200,
			//             "type": "array",
			//             "uniqueItems": false
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "minItems": 1,
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "EventVariables": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Arn": {
			//             "type": "string"
			//           },
			//           "CreatedTime": {
			//             "description": "The time when the event variable was created.",
			//             "type": "string"
			//           },
			//           "DataSource": {
			//             "enum": [
			//               "EVENT"
			//             ],
			//             "type": "string"
			//           },
			//           "DataType": {
			//             "enum": [
			//               "STRING",
			//               "INTEGER",
			//               "FLOAT",
			//               "BOOLEAN"
			//             ],
			//             "type": "string"
			//           },
			//           "DefaultValue": {
			//             "type": "string"
			//           },
			//           "Description": {
			//             "description": "The description.",
			//             "maxLength": 256,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Inline": {
			//             "type": "boolean"
			//           },
			//           "LastUpdatedTime": {
			//             "description": "The time when the event variable was last updated.",
			//             "type": "string"
			//           },
			//           "Name": {
			//             "type": "string"
			//           },
			//           "Tags": {
			//             "description": "Tags associated with this event variable.",
			//             "insertionOrder": false,
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Key": {
			//                   "maxLength": 128,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "maxLength": 256,
			//                   "minLength": 0,
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Key",
			//                 "Value"
			//               ],
			//               "type": "object"
			//             },
			//             "maxItems": 200,
			//             "type": "array",
			//             "uniqueItems": false
			//           },
			//           "VariableType": {
			//             "enum": [
			//               "AUTH_CODE",
			//               "AVS",
			//               "BILLING_ADDRESS_L1",
			//               "BILLING_ADDRESS_L2",
			//               "BILLING_CITY",
			//               "BILLING_COUNTRY",
			//               "BILLING_NAME",
			//               "BILLING_PHONE",
			//               "BILLING_STATE",
			//               "BILLING_ZIP",
			//               "CARD_BIN",
			//               "CATEGORICAL",
			//               "CURRENCY_CODE",
			//               "EMAIL_ADDRESS",
			//               "FINGERPRINT",
			//               "FRAUD_LABEL",
			//               "FREE_FORM_TEXT",
			//               "IP_ADDRESS",
			//               "NUMERIC",
			//               "ORDER_ID",
			//               "PAYMENT_TYPE",
			//               "PHONE_NUMBER",
			//               "PRICE",
			//               "PRODUCT_CATEGORY",
			//               "SHIPPING_ADDRESS_L1",
			//               "SHIPPING_ADDRESS_L2",
			//               "SHIPPING_CITY",
			//               "SHIPPING_COUNTRY",
			//               "SHIPPING_NAME",
			//               "SHIPPING_PHONE",
			//               "SHIPPING_STATE",
			//               "SHIPPING_ZIP",
			//               "USERAGENT"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "minItems": 1,
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "Inline": {
			//       "type": "boolean"
			//     },
			//     "Labels": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Arn": {
			//             "type": "string"
			//           },
			//           "CreatedTime": {
			//             "description": "The time when the label was created.",
			//             "type": "string"
			//           },
			//           "Description": {
			//             "description": "The description.",
			//             "maxLength": 256,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Inline": {
			//             "type": "boolean"
			//           },
			//           "LastUpdatedTime": {
			//             "description": "The time when the label was last updated.",
			//             "type": "string"
			//           },
			//           "Name": {
			//             "type": "string"
			//           },
			//           "Tags": {
			//             "description": "Tags associated with this label.",
			//             "insertionOrder": false,
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Key": {
			//                   "maxLength": 128,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "maxLength": 256,
			//                   "minLength": 0,
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Key",
			//                 "Value"
			//               ],
			//               "type": "object"
			//             },
			//             "maxItems": 200,
			//             "type": "array",
			//             "uniqueItems": false
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "minItems": 2,
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "LastUpdatedTime": {
			//       "description": "The time when the event type was last updated.",
			//       "type": "string"
			//     },
			//     "Name": {
			//       "description": "The name for the event type",
			//       "maxLength": 64,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "Tags": {
			//       "description": "Tags associated with this event type.",
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Key": {
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Value": {
			//             "maxLength": 256,
			//             "minLength": 0,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Key",
			//           "Value"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 200,
			//       "type": "array",
			//       "uniqueItems": false
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Description: "The ARN of the event type.",
						Type:        types.StringType,
						Computed:    true,
					},
					"created_time": {
						// Property: CreatedTime
						Description: "The time when the event type was created.",
						Type:        types.StringType,
						Computed:    true,
					},
					"description": {
						// Property: Description
						Description: "The description of the event type.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
						Optional:    true,
					},
					"entity_types": {
						// Property: EntityTypes
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"arn": {
									// Property: Arn
									Type:     types.StringType,
									Optional: true,
								},
								"created_time": {
									// Property: CreatedTime
									Description: "The time when the entity type was created.",
									Type:        types.StringType,
									Optional:    true,
								},
								"description": {
									// Property: Description
									Description: "The description.",
									Type:        types.StringType,
									Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 256)},
									Optional:    true,
								},
								"inline": {
									// Property: Inline
									Type:     types.BoolType,
									Optional: true,
								},
								"last_updated_time": {
									// Property: LastUpdatedTime
									Description: "The time when the entity type was last updated.",
									Type:        types.StringType,
									Optional:    true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Optional: true,
								},
								"tags": {
									// Property: Tags
									Description: "Tags associated with this entity type.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Type:       types.StringType,
												Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
												Required:   true,
											},
											"value": {
												// Property: Value
												Type:       types.StringType,
												Validators: []tfsdk.AttributeValidator{validate.StringLength(0, 256)},
												Required:   true,
											},
										},
										tfsdk.ListNestedAttributesOptions{
											MaxItems: 200,
										},
									),
									Optional: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{
								MinItems: 1,
							},
						),
						Optional: true,
					},
					"event_variables": {
						// Property: EventVariables
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"arn": {
									// Property: Arn
									Type:     types.StringType,
									Optional: true,
								},
								"created_time": {
									// Property: CreatedTime
									Description: "The time when the event variable was created.",
									Type:        types.StringType,
									Optional:    true,
								},
								"data_source": {
									// Property: DataSource
									Type:     types.StringType,
									Optional: true,
								},
								"data_type": {
									// Property: DataType
									Type:     types.StringType,
									Optional: true,
								},
								"default_value": {
									// Property: DefaultValue
									Type:     types.StringType,
									Optional: true,
								},
								"description": {
									// Property: Description
									Description: "The description.",
									Type:        types.StringType,
									Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 256)},
									Optional:    true,
								},
								"inline": {
									// Property: Inline
									Type:     types.BoolType,
									Optional: true,
								},
								"last_updated_time": {
									// Property: LastUpdatedTime
									Description: "The time when the event variable was last updated.",
									Type:        types.StringType,
									Optional:    true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Optional: true,
								},
								"tags": {
									// Property: Tags
									Description: "Tags associated with this event variable.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Type:       types.StringType,
												Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
												Required:   true,
											},
											"value": {
												// Property: Value
												Type:       types.StringType,
												Validators: []tfsdk.AttributeValidator{validate.StringLength(0, 256)},
												Required:   true,
											},
										},
										tfsdk.ListNestedAttributesOptions{
											MaxItems: 200,
										},
									),
									Optional: true,
								},
								"variable_type": {
									// Property: VariableType
									Type:     types.StringType,
									Optional: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{
								MinItems: 1,
							},
						),
						Optional: true,
					},
					"inline": {
						// Property: Inline
						Type:     types.BoolType,
						Optional: true,
					},
					"labels": {
						// Property: Labels
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"arn": {
									// Property: Arn
									Type:     types.StringType,
									Optional: true,
								},
								"created_time": {
									// Property: CreatedTime
									Description: "The time when the label was created.",
									Type:        types.StringType,
									Optional:    true,
								},
								"description": {
									// Property: Description
									Description: "The description.",
									Type:        types.StringType,
									Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 256)},
									Optional:    true,
								},
								"inline": {
									// Property: Inline
									Type:     types.BoolType,
									Optional: true,
								},
								"last_updated_time": {
									// Property: LastUpdatedTime
									Description: "The time when the label was last updated.",
									Type:        types.StringType,
									Optional:    true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Optional: true,
								},
								"tags": {
									// Property: Tags
									Description: "Tags associated with this label.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Type:       types.StringType,
												Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
												Required:   true,
											},
											"value": {
												// Property: Value
												Type:       types.StringType,
												Validators: []tfsdk.AttributeValidator{validate.StringLength(0, 256)},
												Required:   true,
											},
										},
										tfsdk.ListNestedAttributesOptions{
											MaxItems: 200,
										},
									),
									Optional: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{
								MinItems: 2,
							},
						),
						Optional: true,
					},
					"last_updated_time": {
						// Property: LastUpdatedTime
						Description: "The time when the event type was last updated.",
						Type:        types.StringType,
						Computed:    true,
					},
					"name": {
						// Property: Name
						Description: "The name for the event type",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 64)},
						Optional:    true,
					},
					"tags": {
						// Property: Tags
						Description: "Tags associated with this event type.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Type:       types.StringType,
									Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
									Required:   true,
								},
								"value": {
									// Property: Value
									Type:       types.StringType,
									Validators: []tfsdk.AttributeValidator{validate.StringLength(0, 256)},
									Required:   true,
								},
							},
							tfsdk.ListNestedAttributesOptions{
								MaxItems: 200,
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"last_updated_time": {
			// Property: LastUpdatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time when the detector was last updated.",
			//   "type": "string"
			// }
			Description: "The time when the detector was last updated.",
			Type:        types.StringType,
			Computed:    true,
		},
		"rule_execution_mode": {
			// Property: RuleExecutionMode
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "FIRST_MATCHED",
			//     "ALL_MATCHED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"rules": {
			// Property: Rules
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Arn": {
			//         "type": "string"
			//       },
			//       "CreatedTime": {
			//         "description": "The time when the event type was created.",
			//         "type": "string"
			//       },
			//       "Description": {
			//         "description": "The description.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "DetectorId": {
			//         "type": "string"
			//       },
			//       "Expression": {
			//         "type": "string"
			//       },
			//       "Language": {
			//         "enum": [
			//           "DETECTORPL"
			//         ],
			//         "type": "string"
			//       },
			//       "LastUpdatedTime": {
			//         "description": "The time when the event type was last updated.",
			//         "type": "string"
			//       },
			//       "Outcomes": {
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Arn": {
			//               "type": "string"
			//             },
			//             "CreatedTime": {
			//               "description": "The time when the outcome was created.",
			//               "type": "string"
			//             },
			//             "Description": {
			//               "description": "The description.",
			//               "maxLength": 256,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Inline": {
			//               "type": "boolean"
			//             },
			//             "LastUpdatedTime": {
			//               "description": "The time when the outcome was last updated.",
			//               "type": "string"
			//             },
			//             "Name": {
			//               "type": "string"
			//             },
			//             "Tags": {
			//               "description": "Tags associated with this outcome.",
			//               "insertionOrder": false,
			//               "items": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "Key": {
			//                     "maxLength": 128,
			//                     "minLength": 1,
			//                     "type": "string"
			//                   },
			//                   "Value": {
			//                     "maxLength": 256,
			//                     "minLength": 0,
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "Key",
			//                   "Value"
			//                 ],
			//                 "type": "object"
			//               },
			//               "maxItems": 200,
			//               "type": "array",
			//               "uniqueItems": false
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "minItems": 1,
			//         "type": "array",
			//         "uniqueItems": false
			//       },
			//       "RuleId": {
			//         "type": "string"
			//       },
			//       "RuleVersion": {
			//         "type": "string"
			//       },
			//       "Tags": {
			//         "description": "Tags associated with this event type.",
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Key": {
			//               "maxLength": 128,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Value": {
			//               "maxLength": 256,
			//               "minLength": 0,
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Key",
			//             "Value"
			//           ],
			//           "type": "object"
			//         },
			//         "maxItems": 200,
			//         "type": "array",
			//         "uniqueItems": false
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Type:     types.StringType,
						Optional: true,
					},
					"created_time": {
						// Property: CreatedTime
						Description: "The time when the event type was created.",
						Type:        types.StringType,
						Optional:    true,
					},
					"description": {
						// Property: Description
						Description: "The description.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 256)},
						Optional:    true,
					},
					"detector_id": {
						// Property: DetectorId
						Type:     types.StringType,
						Optional: true,
					},
					"expression": {
						// Property: Expression
						Type:     types.StringType,
						Optional: true,
					},
					"language": {
						// Property: Language
						Type:     types.StringType,
						Optional: true,
					},
					"last_updated_time": {
						// Property: LastUpdatedTime
						Description: "The time when the event type was last updated.",
						Type:        types.StringType,
						Optional:    true,
					},
					"outcomes": {
						// Property: Outcomes
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"arn": {
									// Property: Arn
									Type:     types.StringType,
									Optional: true,
								},
								"created_time": {
									// Property: CreatedTime
									Description: "The time when the outcome was created.",
									Type:        types.StringType,
									Optional:    true,
								},
								"description": {
									// Property: Description
									Description: "The description.",
									Type:        types.StringType,
									Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 256)},
									Optional:    true,
								},
								"inline": {
									// Property: Inline
									Type:     types.BoolType,
									Optional: true,
								},
								"last_updated_time": {
									// Property: LastUpdatedTime
									Description: "The time when the outcome was last updated.",
									Type:        types.StringType,
									Optional:    true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Optional: true,
								},
								"tags": {
									// Property: Tags
									Description: "Tags associated with this outcome.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Type:       types.StringType,
												Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
												Required:   true,
											},
											"value": {
												// Property: Value
												Type:       types.StringType,
												Validators: []tfsdk.AttributeValidator{validate.StringLength(0, 256)},
												Required:   true,
											},
										},
										tfsdk.ListNestedAttributesOptions{
											MaxItems: 200,
										},
									),
									Optional: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{
								MinItems: 1,
							},
						),
						Optional: true,
					},
					"rule_id": {
						// Property: RuleId
						Type:     types.StringType,
						Optional: true,
					},
					"rule_version": {
						// Property: RuleVersion
						Type:     types.StringType,
						Optional: true,
					},
					"tags": {
						// Property: Tags
						Description: "Tags associated with this event type.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Type:       types.StringType,
									Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
									Required:   true,
								},
								"value": {
									// Property: Value
									Type:       types.StringType,
									Validators: []tfsdk.AttributeValidator{validate.StringLength(0, 256)},
									Required:   true,
								},
							},
							tfsdk.ListNestedAttributesOptions{
								MaxItems: 200,
							},
						),
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Required: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags associated with this detector.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
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
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "Tags associated with this detector.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 128)},
						Required:   true,
					},
					"value": {
						// Property: Value
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLength(0, 256)},
						Required:   true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MaxItems: 200,
				},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "A resource schema for a Detector in Amazon Fraud Detector.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::Detector").WithTerraformTypeName("awscc_frauddetector_detector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"associated_models":       "AssociatedModels",
		"created_time":            "CreatedTime",
		"data_source":             "DataSource",
		"data_type":               "DataType",
		"default_value":           "DefaultValue",
		"description":             "Description",
		"detector_id":             "DetectorId",
		"detector_version_id":     "DetectorVersionId",
		"detector_version_status": "DetectorVersionStatus",
		"entity_types":            "EntityTypes",
		"event_type":              "EventType",
		"event_variables":         "EventVariables",
		"expression":              "Expression",
		"inline":                  "Inline",
		"key":                     "Key",
		"labels":                  "Labels",
		"language":                "Language",
		"last_updated_time":       "LastUpdatedTime",
		"name":                    "Name",
		"outcomes":                "Outcomes",
		"rule_execution_mode":     "RuleExecutionMode",
		"rule_id":                 "RuleId",
		"rule_version":            "RuleVersion",
		"rules":                   "Rules",
		"tags":                    "Tags",
		"value":                   "Value",
		"variable_type":           "VariableType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_frauddetector_detector", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
