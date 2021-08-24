// Code generated by generators/resource/main.go; DO NOT EDIT.

package signer

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
	registry.AddResourceTypeFactory("awscc_signer_signing_profile", signingProfileResourceType)
}

// signingProfileResourceType returns the Terraform awscc_signer_signing_profile resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Signer::SigningProfile resource type.
func signingProfileResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"platform_id": {
			// Property: PlatformId
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "AWSLambda-SHA384-ECDSA"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// PlatformId is a force-new attribute.
		},
		"profile_name": {
			// Property: ProfileName
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the signing profile. AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. ",
			//   "type": "string"
			// }
			Description: "A name for the signing profile. AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"profile_version": {
			// Property: ProfileVersion
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"profile_version_arn": {
			// Property: ProfileVersionArn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"signature_validity_period": {
			// Property: SignatureValidityPeriod
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Type": {
			//       "enum": [
			//         "DAYS",
			//         "MONTHS",
			//         "YEARS"
			//       ],
			//       "type": "string"
			//     },
			//     "Value": {
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"type": {
						// Property: Type
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						Type:     types.NumberType,
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// SignatureValidityPeriod is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of tags associated with the signing profile.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A list of tags associated with the signing profile.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 127)},
						Optional:   true,
					},
					"value": {
						// Property: Value
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 255)},
						Optional:   true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
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
		Description: "A signing profile is a signing template that can be used to carry out a pre-defined signing job.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Signer::SigningProfile").WithTerraformTypeName("awscc_signer_signing_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"key":                       "Key",
		"platform_id":               "PlatformId",
		"profile_name":              "ProfileName",
		"profile_version":           "ProfileVersion",
		"profile_version_arn":       "ProfileVersionArn",
		"signature_validity_period": "SignatureValidityPeriod",
		"tags":                      "Tags",
		"type":                      "Type",
		"value":                     "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_signer_signing_profile", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
