// Code generated by generators/resource/main.go; DO NOT EDIT.

package eventschemas

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
	registry.AddResourceTypeFactory("aws_eventschemas_registry_policy", registryPolicy)
}

// registryPolicy returns the Terraform aws_eventschemas_registry_policy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EventSchemas::RegistryPolicy resource type.
func registryPolicy(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "object"
			   }
			*/
			Type:     types.MapType{ElemType: types.StringType},
			Required: true,
		},
		"registry_name": {
			// Property: RegistryName
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
		},
		"revision_id": {
			// Property: RevisionId
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
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
		Description: "Resource Type definition for AWS::EventSchemas::RegistryPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EventSchemas::RegistryPolicy").WithTerraformTypeName("aws_eventschemas_registry_policy").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_eventschemas_registry_policy", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
