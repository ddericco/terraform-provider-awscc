// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53resolver

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
	registry.AddResourceTypeFactory("awscc_route53resolver_resolver_dnssec_config", resolverDNSSECConfigResourceType)
}

// resolverDNSSECConfigResourceType returns the Terraform awscc_route53resolver_resolver_dnssec_config resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53Resolver::ResolverDNSSECConfig resource type.
func resolverDNSSECConfigResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Id",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Id",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 64)},
			Computed:    true,
		},
		"owner_id": {
			// Property: OwnerId
			// CloudFormation resource type schema:
			// {
			//   "description": "AccountId",
			//   "maxLength": 32,
			//   "minLength": 12,
			//   "type": "string"
			// }
			Description: "AccountId",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(12, 32)},
			Computed:    true,
		},
		"resource_id": {
			// Property: ResourceId
			// CloudFormation resource type schema:
			// {
			//   "description": "ResourceId",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "ResourceId",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 64)},
			Optional:    true,
			Computed:    true,
			// ResourceId is a force-new attribute.
		},
		"validation_status": {
			// Property: ValidationStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.",
			//   "enum": [
			//     "ENABLING",
			//     "ENABLED",
			//     "DISABLING",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Description: "ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::ResolverDNSSECConfig").WithTerraformTypeName("awscc_route53resolver_resolver_dnssec_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":                "Id",
		"owner_id":          "OwnerId",
		"resource_id":       "ResourceId",
		"validation_status": "ValidationStatus",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53resolver_resolver_dnssec_config", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
