// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_route53recoveryreadiness_cell", cellDataSourceType)
}

// cellDataSourceType returns the Terraform awscc_route53recoveryreadiness_cell data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Route53RecoveryReadiness::Cell resource type.
func cellDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cell_arn": {
			// Property: CellArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the cell.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the cell.",
			Type:        types.StringType,
			Computed:    true,
		},
		"cell_name": {
			// Property: CellName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the cell to create.",
			//   "maxLength": 64,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the cell to create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"cells": {
			// Property: Cells
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells. For example, Availability Zones within specific Regions.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells. For example, Availability Zones within specific Regions.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"parent_readiness_scopes": {
			// Property: ParentReadinessScopes
			// CloudFormation resource type schema:
			// {
			//   "description": "The readiness scope for the cell, which can be a cell Amazon Resource Name (ARN) or a recovery group ARN. This is a list but currently can have only one element.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "The readiness scope for the cell, which can be a cell Amazon Resource Name (ARN) or a recovery group ARN. This is a list but currently can have only one element.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "insertionOrder": false,
			//         "items": {
			//           "maxItems": 50,
			//           "type": "string"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A collection of tags associated with a resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
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
		Description: "Data Source schema for AWS::Route53RecoveryReadiness::Cell",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryReadiness::Cell").WithTerraformTypeName("awscc_route53recoveryreadiness_cell")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cell_arn":                "CellArn",
		"cell_name":               "CellName",
		"cells":                   "Cells",
		"key":                     "Key",
		"parent_readiness_scopes": "ParentReadinessScopes",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
