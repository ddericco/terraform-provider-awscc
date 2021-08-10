// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkmanager

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
	registry.AddResourceTypeFactory("aws_networkmanager_global_network", globalNetwork)
}

// globalNetwork returns the Terraform aws_networkmanager_global_network resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkManager::GlobalNetwork resource type.
func globalNetwork(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the global network.",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the global network.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The description of the global network.",
			     "type": "string"
			   }
			*/
			Description: "The description of the global network.",
			Type:        types.StringType,
			Optional:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the global network.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the global network.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The tags for the global network.",
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to associate with a global network resource.",
			       "properties": {
			         "Key": {
			           "type": "string"
			         },
			         "Value": {
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "The tags for the global network.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
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
		Description: "The AWS::NetworkManager::GlobalNetwork type specifies a global network of the user's account",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::GlobalNetwork").WithTerraformTypeName("aws_networkmanager_global_network").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_networkmanager_global_network", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
