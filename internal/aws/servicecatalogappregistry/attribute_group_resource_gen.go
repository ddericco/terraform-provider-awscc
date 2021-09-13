// Code generated by generators/resource/main.go; DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_servicecatalogappregistry_attribute_group", attributeGroupResourceType)
}

// attributeGroupResourceType returns the Terraform awscc_servicecatalogappregistry_attribute_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ServiceCatalogAppRegistry::AttributeGroup resource type.
func attributeGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
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
		"attributes": {
			// Property: Attributes
			// CloudFormation resource type schema:
			// {
			//   "type": "object"
			// }
			Type:     types.MapType{ElemType: types.StringType},
			Required: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the attribute group. ",
			//   "maxLength": 1024,
			//   "type": "string"
			// }
			Description: "The description of the attribute group. ",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1024),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the attribute group. ",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the attribute group. ",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroup.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalogAppRegistry::AttributeGroup").WithTerraformTypeName("awscc_servicecatalogappregistry_attribute_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":         "Arn",
		"attributes":  "Attributes",
		"description": "Description",
		"id":          "Id",
		"name":        "Name",
		"tags":        "Tags",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
