// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs

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
	registry.AddResourceTypeFactory("aws_logs_resource_policy", resourcePolicy)
}

// resourcePolicy returns the Terraform aws_logs_resource_policy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Logs::ResourcePolicy resource type.
func resourcePolicy(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"policy_document": {
			// Property: PolicyDocument
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The policy document",
			     "maxLength": 5120,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The policy document",
			Type:        types.StringType,
			Required:    true,
		},
		"policy_name": {
			// Property: PolicyName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A name for resource policy",
			     "maxLength": 255,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "A name for resource policy",
			Type:        types.StringType,
			Required:    true,
			// PolicyName is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The resource schema for AWSLogs ResourcePolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::ResourcePolicy").WithTerraformTypeName("aws_logs_resource_policy").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_logs_resource_policy", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
