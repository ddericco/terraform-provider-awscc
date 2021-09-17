// Code generated by generators/resource/main.go; DO NOT EDIT.

package backup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_backup_backup_plan", backupPlanResourceType)
}

// backupPlanResourceType returns the Terraform awscc_backup_backup_plan resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Backup::BackupPlan resource type.
func backupPlanResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"backup_plan": {
			// Property: BackupPlan
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AdvancedBackupSettings": {
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "BackupOptions": {
			//             "type": "object"
			//           },
			//           "ResourceType": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "BackupOptions",
			//           "ResourceType"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "BackupPlanName": {
			//       "type": "string"
			//     },
			//     "BackupPlanRule": {
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "CompletionWindowMinutes": {
			//             "type": "number"
			//           },
			//           "CopyActions": {
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "DestinationBackupVaultArn": {
			//                   "type": "string"
			//                 },
			//                 "Lifecycle": {
			//                   "additionalProperties": false,
			//                   "properties": {
			//                     "DeleteAfterDays": {
			//                       "type": "number"
			//                     },
			//                     "MoveToColdStorageAfterDays": {
			//                       "type": "number"
			//                     }
			//                   },
			//                   "type": "object"
			//                 }
			//               },
			//               "required": [
			//                 "DestinationBackupVaultArn"
			//               ],
			//               "type": "object"
			//             },
			//             "type": "array",
			//             "uniqueItems": false
			//           },
			//           "EnableContinuousBackup": {
			//             "type": "boolean"
			//           },
			//           "Lifecycle": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DeleteAfterDays": {
			//                 "type": "number"
			//               },
			//               "MoveToColdStorageAfterDays": {
			//                 "type": "number"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "RecoveryPointTags": {
			//             "additionalProperties": false,
			//             "patternProperties": {
			//               "": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "RuleName": {
			//             "type": "string"
			//           },
			//           "ScheduleExpression": {
			//             "type": "string"
			//           },
			//           "StartWindowMinutes": {
			//             "type": "number"
			//           },
			//           "TargetBackupVault": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "TargetBackupVault",
			//           "RuleName"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     }
			//   },
			//   "required": [
			//     "BackupPlanName",
			//     "BackupPlanRule"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"advanced_backup_settings": {
						// Property: AdvancedBackupSettings
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"backup_options": {
									// Property: BackupOptions
									Type:     types.MapType{ElemType: types.StringType},
									Required: true,
								},
								"resource_type": {
									// Property: ResourceType
									Type:     types.StringType,
									Required: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
					"backup_plan_name": {
						// Property: BackupPlanName
						Type:     types.StringType,
						Required: true,
					},
					"backup_plan_rule": {
						// Property: BackupPlanRule
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"completion_window_minutes": {
									// Property: CompletionWindowMinutes
									Type:     types.NumberType,
									Optional: true,
								},
								"copy_actions": {
									// Property: CopyActions
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"destination_backup_vault_arn": {
												// Property: DestinationBackupVaultArn
												Type:     types.StringType,
												Required: true,
											},
											"lifecycle": {
												// Property: Lifecycle
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"delete_after_days": {
															// Property: DeleteAfterDays
															Type:     types.NumberType,
															Optional: true,
														},
														"move_to_cold_storage_after_days": {
															// Property: MoveToColdStorageAfterDays
															Type:     types.NumberType,
															Optional: true,
														},
													},
												),
												Optional: true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
								"enable_continuous_backup": {
									// Property: EnableContinuousBackup
									Type:     types.BoolType,
									Optional: true,
								},
								"lifecycle": {
									// Property: Lifecycle
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"delete_after_days": {
												// Property: DeleteAfterDays
												Type:     types.NumberType,
												Optional: true,
											},
											"move_to_cold_storage_after_days": {
												// Property: MoveToColdStorageAfterDays
												Type:     types.NumberType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"recovery_point_tags": {
									// Property: RecoveryPointTags
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
								"rule_name": {
									// Property: RuleName
									Type:     types.StringType,
									Required: true,
								},
								"schedule_expression": {
									// Property: ScheduleExpression
									Type:     types.StringType,
									Optional: true,
								},
								"start_window_minutes": {
									// Property: StartWindowMinutes
									Type:     types.NumberType,
									Optional: true,
								},
								"target_backup_vault": {
									// Property: TargetBackupVault
									Type:     types.StringType,
									Required: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Required: true,
					},
				},
			),
			Required: true,
		},
		"backup_plan_arn": {
			// Property: BackupPlanArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"backup_plan_id": {
			// Property: BackupPlanId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"backup_plan_tags": {
			// Property: BackupPlanTags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
		"version_id": {
			// Property: VersionId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Backup::BackupPlan",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Backup::BackupPlan").WithTerraformTypeName("awscc_backup_backup_plan")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"advanced_backup_settings":        "AdvancedBackupSettings",
		"backup_options":                  "BackupOptions",
		"backup_plan":                     "BackupPlan",
		"backup_plan_arn":                 "BackupPlanArn",
		"backup_plan_id":                  "BackupPlanId",
		"backup_plan_name":                "BackupPlanName",
		"backup_plan_rule":                "BackupPlanRule",
		"backup_plan_tags":                "BackupPlanTags",
		"completion_window_minutes":       "CompletionWindowMinutes",
		"copy_actions":                    "CopyActions",
		"delete_after_days":               "DeleteAfterDays",
		"destination_backup_vault_arn":    "DestinationBackupVaultArn",
		"enable_continuous_backup":        "EnableContinuousBackup",
		"lifecycle":                       "Lifecycle",
		"move_to_cold_storage_after_days": "MoveToColdStorageAfterDays",
		"recovery_point_tags":             "RecoveryPointTags",
		"resource_type":                   "ResourceType",
		"rule_name":                       "RuleName",
		"schedule_expression":             "ScheduleExpression",
		"start_window_minutes":            "StartWindowMinutes",
		"target_backup_vault":             "TargetBackupVault",
		"version_id":                      "VersionId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
