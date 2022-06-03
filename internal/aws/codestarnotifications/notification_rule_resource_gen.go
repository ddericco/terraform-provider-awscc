// Code generated by generators/resource/main.go; DO NOT EDIT.

package codestarnotifications

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_codestarnotifications_notification_rule", notificationRuleResourceType)
}

// notificationRuleResourceType returns the Terraform awscc_codestarnotifications_notification_rule resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CodeStarNotifications::NotificationRule resource type.
func notificationRuleResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "^arn:aws[^:\\s]*:codestar-notifications:[^:\\s]+:\\d{12}:notificationrule\\/(.*\\S)?$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"created_by": {
			// Property: CreatedBy
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
		},
		"detail_type": {
			// Property: DetailType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "BASIC",
			//     "FULL"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"BASIC",
					"FULL",
				}),
			},
		},
		"event_type_id": {
			// Property: EventTypeId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
		},
		"event_type_ids": {
			// Property: EventTypeIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "maxLength": 200,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringLenBetween(1, 200)),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "[A-Za-z0-9\\-_ ]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
				validate.StringMatch(regexp.MustCompile("[A-Za-z0-9\\-_ ]+$"), ""),
			},
		},
		"resource": {
			// Property: Resource
			// CloudFormation resource type schema:
			// {
			//   "pattern": "^arn:aws[^:\\s]*:[^:\\s]*:[^:\\s]*:[0-9]{12}:[^\\s]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^arn:aws[^:\\s]*:[^:\\s]*:[^:\\s]*:[0-9]{12}:[^\\s]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ENABLED",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ENABLED",
					"DISABLED",
				}),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "type": "object"
			// }
			Type:     JSONStringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				JSONStringType.AttributePlanModifier(),
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"target_address": {
			// Property: TargetAddress
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
		},
		"targets": {
			// Property: Targets
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "TargetAddress": {
			//         "type": "string"
			//       },
			//       "TargetType": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "TargetType",
			//       "TargetAddress"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 10,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"target_address": {
						// Property: TargetAddress
						Type:     types.StringType,
						Required: true,
					},
					"target_type": {
						// Property: TargetType
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(10),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::CodeStarNotifications::NotificationRule",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeStarNotifications::NotificationRule").WithTerraformTypeName("awscc_codestarnotifications_notification_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":            "Arn",
		"created_by":     "CreatedBy",
		"detail_type":    "DetailType",
		"event_type_id":  "EventTypeId",
		"event_type_ids": "EventTypeIds",
		"name":           "Name",
		"resource":       "Resource",
		"status":         "Status",
		"tags":           "Tags",
		"target_address": "TargetAddress",
		"target_type":    "TargetType",
		"targets":        "Targets",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
