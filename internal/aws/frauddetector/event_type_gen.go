// Code generated by generators/resource/main.go; DO NOT EDIT.

package frauddetector

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
	registry.AddResourceTypeFactory("aws_frauddetector_event_type", eventType)
}

// eventType returns the Terraform aws_frauddetector_event_type resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::FraudDetector::EventType resource type.
func eventType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the event type.",
			     "type": "string"
			   }
			*/
			Description: "The ARN of the event type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The time when the event type was created.",
			     "type": "string"
			   }
			*/
			Description: "The time when the event type was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The description of the event type.",
			     "maxLength": 128,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "The description of the event type.",
			Type:        types.StringType,
			Optional:    true,
		},
		"entity_types": {
			// Property: EntityTypes
			// CloudFormation resource type schema:
			/*
			   {
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Arn": {
			           "type": "string"
			         },
			         "CreatedTime": {
			           "description": "The time when the event type was created.",
			           "type": "string"
			         },
			         "Description": {
			           "description": "The description.",
			           "maxLength": 256,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Inline": {
			           "type": "boolean"
			         },
			         "LastUpdatedTime": {
			           "description": "The time when the event type was last updated.",
			           "type": "string"
			         },
			         "Name": {
			           "type": "string"
			         },
			         "Tags": {
			           "description": "Tags associated with this event type.",
			           "insertionOrder": false,
			           "items": {
			             "additionalProperties": false,
			             "properties": {
			               "Key": {
			                 "maxLength": 128,
			                 "minLength": 1,
			                 "type": "string"
			               },
			               "Value": {
			                 "maxLength": 256,
			                 "minLength": 0,
			                 "type": "string"
			               }
			             },
			             "$ref": "#/definitions/Tag",
			             "required": [
			               "Key",
			               "Value"
			             ],
			             "type": "object"
			           },
			           "maxItems": 200,
			           "type": "array",
			           "uniqueItems": false
			         }
			       },
			       "$ref": "#/definitions/EntityType",
			       "type": "object"
			     },
			     "minItems": 1,
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"arn": {
						// Property: Arn
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"created_time": {
						// Property: CreatedTime
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The time when the event type was created.",
						     "type": "string"
						   }
						*/
						Description: "The time when the event type was created.",
						Type:        types.StringType,
						Optional:    true,
					},
					"description": {
						// Property: Description
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The description.",
						     "maxLength": 256,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The description.",
						Type:        types.StringType,
						Optional:    true,
					},
					"inline": {
						// Property: Inline
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
					"last_updated_time": {
						// Property: LastUpdatedTime
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The time when the event type was last updated.",
						     "type": "string"
						   }
						*/
						Description: "The time when the event type was last updated.",
						Type:        types.StringType,
						Optional:    true,
					},
					"name": {
						// Property: Name
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"tags": {
						// Property: Tags
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Tags associated with this event type.",
						     "insertionOrder": false,
						     "items": {
						       "additionalProperties": false,
						       "properties": {
						         "Key": {
						           "maxLength": 128,
						           "minLength": 1,
						           "type": "string"
						         },
						         "Value": {
						           "maxLength": 256,
						           "minLength": 0,
						           "type": "string"
						         }
						       },
						       "$ref": "#/definitions/Tag",
						       "required": [
						         "Key",
						         "Value"
						       ],
						       "type": "object"
						     },
						     "maxItems": 200,
						     "type": "array",
						     "uniqueItems": false
						   }
						*/
						Description: "Tags associated with this event type.",
						// Multiset.
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"key": {
									// Property: Key
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 128,
									     "minLength": 1,
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"value": {
									// Property: Value
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 256,
									     "minLength": 0,
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{
								MaxItems: 200,
							},
						),
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Required: true,
		},
		"event_variables": {
			// Property: EventVariables
			// CloudFormation resource type schema:
			/*
			   {
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Arn": {
			           "type": "string"
			         },
			         "CreatedTime": {
			           "description": "The time when the event type was created.",
			           "type": "string"
			         },
			         "DataSource": {
			           "enum": [
			             "EVENT"
			           ],
			           "type": "string"
			         },
			         "DataType": {
			           "enum": [
			             "STRING",
			             "INTEGER",
			             "FLOAT",
			             "BOOLEAN"
			           ],
			           "type": "string"
			         },
			         "DefaultValue": {
			           "type": "string"
			         },
			         "Description": {
			           "description": "The description.",
			           "maxLength": 256,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Inline": {
			           "type": "boolean"
			         },
			         "LastUpdatedTime": {
			           "description": "The time when the event type was last updated.",
			           "type": "string"
			         },
			         "Name": {
			           "type": "string"
			         },
			         "Tags": {
			           "description": "Tags associated with this event type.",
			           "insertionOrder": false,
			           "items": {
			             "additionalProperties": false,
			             "properties": {
			               "Key": {
			                 "maxLength": 128,
			                 "minLength": 1,
			                 "type": "string"
			               },
			               "Value": {
			                 "maxLength": 256,
			                 "minLength": 0,
			                 "type": "string"
			               }
			             },
			             "$ref": "#/definitions/Tag",
			             "required": [
			               "Key",
			               "Value"
			             ],
			             "type": "object"
			           },
			           "maxItems": 200,
			           "type": "array",
			           "uniqueItems": false
			         },
			         "VariableType": {
			           "enum": [
			             "AUTH_CODE",
			             "AVS",
			             "BILLING_ADDRESS_L1",
			             "BILLING_ADDRESS_L2",
			             "BILLING_CITY",
			             "BILLING_COUNTRY",
			             "BILLING_NAME",
			             "BILLING_PHONE",
			             "BILLING_STATE",
			             "BILLING_ZIP",
			             "CARD_BIN",
			             "CATEGORICAL",
			             "CURRENCY_CODE",
			             "EMAIL_ADDRESS",
			             "FINGERPRINT",
			             "FRAUD_LABEL",
			             "FREE_FORM_TEXT",
			             "IP_ADDRESS",
			             "NUMERIC",
			             "ORDER_ID",
			             "PAYMENT_TYPE",
			             "PHONE_NUMBER",
			             "PRICE",
			             "PRODUCT_CATEGORY",
			             "SHIPPING_ADDRESS_L1",
			             "SHIPPING_ADDRESS_L2",
			             "SHIPPING_CITY",
			             "SHIPPING_COUNTRY",
			             "SHIPPING_NAME",
			             "SHIPPING_PHONE",
			             "SHIPPING_STATE",
			             "SHIPPING_ZIP",
			             "USERAGENT"
			           ],
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/EventVariable",
			       "type": "object"
			     },
			     "minItems": 1,
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"arn": {
						// Property: Arn
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"created_time": {
						// Property: CreatedTime
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The time when the event type was created.",
						     "type": "string"
						   }
						*/
						Description: "The time when the event type was created.",
						Type:        types.StringType,
						Optional:    true,
					},
					"data_source": {
						// Property: DataSource
						// CloudFormation resource type schema:
						/*
						   {
						     "enum": [
						       "EVENT"
						     ],
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"data_type": {
						// Property: DataType
						// CloudFormation resource type schema:
						/*
						   {
						     "enum": [
						       "STRING",
						       "INTEGER",
						       "FLOAT",
						       "BOOLEAN"
						     ],
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"default_value": {
						// Property: DefaultValue
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"description": {
						// Property: Description
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The description.",
						     "maxLength": 256,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The description.",
						Type:        types.StringType,
						Optional:    true,
					},
					"inline": {
						// Property: Inline
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
					"last_updated_time": {
						// Property: LastUpdatedTime
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The time when the event type was last updated.",
						     "type": "string"
						   }
						*/
						Description: "The time when the event type was last updated.",
						Type:        types.StringType,
						Optional:    true,
					},
					"name": {
						// Property: Name
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"tags": {
						// Property: Tags
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Tags associated with this event type.",
						     "insertionOrder": false,
						     "items": {
						       "additionalProperties": false,
						       "properties": {
						         "Key": {
						           "maxLength": 128,
						           "minLength": 1,
						           "type": "string"
						         },
						         "Value": {
						           "maxLength": 256,
						           "minLength": 0,
						           "type": "string"
						         }
						       },
						       "$ref": "#/definitions/Tag",
						       "required": [
						         "Key",
						         "Value"
						       ],
						       "type": "object"
						     },
						     "maxItems": 200,
						     "type": "array",
						     "uniqueItems": false
						   }
						*/
						Description: "Tags associated with this event type.",
						// Multiset.
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"key": {
									// Property: Key
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 128,
									     "minLength": 1,
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"value": {
									// Property: Value
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 256,
									     "minLength": 0,
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{
								MaxItems: 200,
							},
						),
						Optional: true,
					},
					"variable_type": {
						// Property: VariableType
						// CloudFormation resource type schema:
						/*
						   {
						     "enum": [
						       "AUTH_CODE",
						       "AVS",
						       "BILLING_ADDRESS_L1",
						       "BILLING_ADDRESS_L2",
						       "BILLING_CITY",
						       "BILLING_COUNTRY",
						       "BILLING_NAME",
						       "BILLING_PHONE",
						       "BILLING_STATE",
						       "BILLING_ZIP",
						       "CARD_BIN",
						       "CATEGORICAL",
						       "CURRENCY_CODE",
						       "EMAIL_ADDRESS",
						       "FINGERPRINT",
						       "FRAUD_LABEL",
						       "FREE_FORM_TEXT",
						       "IP_ADDRESS",
						       "NUMERIC",
						       "ORDER_ID",
						       "PAYMENT_TYPE",
						       "PHONE_NUMBER",
						       "PRICE",
						       "PRODUCT_CATEGORY",
						       "SHIPPING_ADDRESS_L1",
						       "SHIPPING_ADDRESS_L2",
						       "SHIPPING_CITY",
						       "SHIPPING_COUNTRY",
						       "SHIPPING_NAME",
						       "SHIPPING_PHONE",
						       "SHIPPING_STATE",
						       "SHIPPING_ZIP",
						       "USERAGENT"
						     ],
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Required: true,
		},
		"labels": {
			// Property: Labels
			// CloudFormation resource type schema:
			/*
			   {
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Arn": {
			           "type": "string"
			         },
			         "CreatedTime": {
			           "description": "The time when the event type was created.",
			           "type": "string"
			         },
			         "Description": {
			           "description": "The description.",
			           "maxLength": 256,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Inline": {
			           "type": "boolean"
			         },
			         "LastUpdatedTime": {
			           "description": "The time when the event type was last updated.",
			           "type": "string"
			         },
			         "Name": {
			           "type": "string"
			         },
			         "Tags": {
			           "description": "Tags associated with this event type.",
			           "insertionOrder": false,
			           "items": {
			             "additionalProperties": false,
			             "properties": {
			               "Key": {
			                 "maxLength": 128,
			                 "minLength": 1,
			                 "type": "string"
			               },
			               "Value": {
			                 "maxLength": 256,
			                 "minLength": 0,
			                 "type": "string"
			               }
			             },
			             "$ref": "#/definitions/Tag",
			             "required": [
			               "Key",
			               "Value"
			             ],
			             "type": "object"
			           },
			           "maxItems": 200,
			           "type": "array",
			           "uniqueItems": false
			         }
			       },
			       "$ref": "#/definitions/Label",
			       "type": "object"
			     },
			     "minItems": 2,
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"arn": {
						// Property: Arn
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"created_time": {
						// Property: CreatedTime
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The time when the event type was created.",
						     "type": "string"
						   }
						*/
						Description: "The time when the event type was created.",
						Type:        types.StringType,
						Optional:    true,
					},
					"description": {
						// Property: Description
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The description.",
						     "maxLength": 256,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The description.",
						Type:        types.StringType,
						Optional:    true,
					},
					"inline": {
						// Property: Inline
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
					"last_updated_time": {
						// Property: LastUpdatedTime
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The time when the event type was last updated.",
						     "type": "string"
						   }
						*/
						Description: "The time when the event type was last updated.",
						Type:        types.StringType,
						Optional:    true,
					},
					"name": {
						// Property: Name
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"tags": {
						// Property: Tags
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Tags associated with this event type.",
						     "insertionOrder": false,
						     "items": {
						       "additionalProperties": false,
						       "properties": {
						         "Key": {
						           "maxLength": 128,
						           "minLength": 1,
						           "type": "string"
						         },
						         "Value": {
						           "maxLength": 256,
						           "minLength": 0,
						           "type": "string"
						         }
						       },
						       "$ref": "#/definitions/Tag",
						       "required": [
						         "Key",
						         "Value"
						       ],
						       "type": "object"
						     },
						     "maxItems": 200,
						     "type": "array",
						     "uniqueItems": false
						   }
						*/
						Description: "Tags associated with this event type.",
						// Multiset.
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"key": {
									// Property: Key
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 128,
									     "minLength": 1,
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"value": {
									// Property: Value
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 256,
									     "minLength": 0,
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{
								MaxItems: 200,
							},
						),
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 2,
				},
			),
			Required: true,
		},
		"last_updated_time": {
			// Property: LastUpdatedTime
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The time when the event type was last updated.",
			     "type": "string"
			   }
			*/
			Description: "The time when the event type was last updated.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name for the event type",
			     "maxLength": 64,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name for the event type",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Tags associated with this event type.",
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "maxLength": 256,
			           "minLength": 0,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "maxItems": 200,
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Description: "Tags associated with this event type.",
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 128,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 256,
						     "minLength": 0,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 200,
				},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "A resource schema for an EventType in Amazon Fraud Detector.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::EventType").WithTerraformTypeName("aws_frauddetector_event_type").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_frauddetector_event_type", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
