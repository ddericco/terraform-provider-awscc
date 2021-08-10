// Code generated by generators/resource/main.go; DO NOT EDIT.

package frauddetector

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
	registry.AddResourceTypeFactory("aws_frauddetector_label", label)
}

// label returns the Terraform aws_frauddetector_label resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::FraudDetector::Label resource type.
func label(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The label ARN.",
			     "type": "string"
			   }
			*/
			Description: "The label ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The timestamp when the label was created.",
			     "type": "string"
			   }
			*/
			Description: "The timestamp when the label was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The label description.",
			     "maxLength": 128,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "The label description.",
			Type:        types.StringType,
			Optional:    true,
		},
		"last_updated_time": {
			// Property: LastUpdatedTime
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The timestamp when the label was last updated.",
			     "type": "string"
			   }
			*/
			Description: "The timestamp when the label was last updated.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the label.",
			     "maxLength": 64,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name of the label.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Tags associated with this label.",
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "maxLength": 256,
			           "minLength": 0,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "maxItems": 200,
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Description: "Tags associated with this label.",
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 128,
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
						     "maxLength": 256,
						     "minLength": 0,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 200,
				},
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
		Description: "An label for fraud detector.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::Label").WithTerraformTypeName("aws_frauddetector_label").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_frauddetector_label", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
