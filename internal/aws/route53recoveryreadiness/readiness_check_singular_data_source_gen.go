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
	registry.AddDataSourceTypeFactory("awscc_route53recoveryreadiness_readiness_check", readinessCheckDataSourceType)
}

// readinessCheckDataSourceType returns the Terraform awscc_route53recoveryreadiness_readiness_check data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Route53RecoveryReadiness::ReadinessCheck resource type.
func readinessCheckDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"readiness_check_arn": {
			// Property: ReadinessCheckArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the readiness check.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the readiness check.",
			Type:        types.StringType,
			Computed:    true,
		},
		"readiness_check_name": {
			// Property: ReadinessCheckName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the ReadinessCheck to create.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of the ReadinessCheck to create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_set_name": {
			// Property: ResourceSetName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the resource set to check.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the resource set to check.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource.",
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
			Description: "A collection of tags associated with a resource.",
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
		Description: "Data Source schema for AWS::Route53RecoveryReadiness::ReadinessCheck",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryReadiness::ReadinessCheck").WithTerraformTypeName("awscc_route53recoveryreadiness_readiness_check")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                  "Key",
		"readiness_check_arn":  "ReadinessCheckArn",
		"readiness_check_name": "ReadinessCheckName",
		"resource_set_name":    "ResourceSetName",
		"tags":                 "Tags",
		"value":                "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
