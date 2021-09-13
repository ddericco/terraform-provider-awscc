// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_apigateway_documentation_version", documentationVersionResourceType)
}

// documentationVersionResourceType returns the Terraform awscc_apigateway_documentation_version resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ApiGateway::DocumentationVersion resource type.
func documentationVersionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the API documentation snapshot.",
			//   "type": "string"
			// }
			Description: "The description of the API documentation snapshot.",
			Type:        types.StringType,
			Optional:    true,
		},
		"documentation_version": {
			// Property: DocumentationVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The version identifier of the API documentation snapshot.",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The version identifier of the API documentation snapshot.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"rest_api_id": {
			// Property: RestApiId
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the API.",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The identifier of the API.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "A snapshot of the documentation of an API.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::DocumentationVersion").WithTerraformTypeName("awscc_apigateway_documentation_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":           "Description",
		"documentation_version": "DocumentationVersion",
		"rest_api_id":           "RestApiId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
