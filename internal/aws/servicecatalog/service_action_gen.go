// Code generated by generators/resource/main.go; DO NOT EDIT.

package servicecatalog

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
	registry.AddResourceTypeFactory("aws_servicecatalog_service_action", serviceAction)
}

// serviceAction returns the Terraform aws_servicecatalog_service_action resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ServiceCatalog::ServiceAction resource type.
func serviceAction(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"accept_language": {
			// Property: AcceptLanguage
			// CloudFormation resource type schema:
			/*
			   {
			     "enum": [
			       "en",
			       "jp",
			       "zh"
			     ],
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
		},
		"definition": {
			// Property: Definition
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "maxLength": 1000,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "maxLength": 4096,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/DefinitionParameter",
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 1000,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 4096,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Required: true,
		},
		"definition_type": {
			// Property: DefinitionType
			// CloudFormation resource type schema:
			/*
			   {
			     "enum": [
			       "SSM_AUTOMATION"
			     ],
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 1024,
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 100,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 256,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Schema for AWS::ServiceCatalog::ServiceAction",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalog::ServiceAction").WithTerraformTypeName("aws_servicecatalog_service_action").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_servicecatalog_service_action", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
