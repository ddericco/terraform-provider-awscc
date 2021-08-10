// Code generated by generators/resource/main.go; DO NOT EDIT.

package lambda

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
	registry.AddResourceTypeFactory("aws_lambda_code_signing_config", codeSigningConfig)
}

// codeSigningConfig returns the Terraform aws_lambda_code_signing_config resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Lambda::CodeSigningConfig resource type.
func codeSigningConfig(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"allowed_publishers": {
			// Property: AllowedPublishers
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list",
			     "properties": {
			       "SigningProfileVersionArns": {
			         "description": "List of Signing profile version Arns",
			         "items": {
			           "maxLength": 1024,
			           "minLength": 12,
			           "pattern": "",
			           "type": "string"
			         },
			         "maxItems": 20,
			         "minItems": 1,
			         "type": "array"
			       }
			     },
			     "$ref": "#/definitions/AllowedPublishers",
			     "required": [
			       "SigningProfileVersionArns"
			     ],
			     "type": "object"
			   }
			*/
			Description: "When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"signing_profile_version_arns": {
						// Property: SigningProfileVersionArns
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "List of Signing profile version Arns",
						     "items": {
						       "maxLength": 1024,
						       "minLength": 12,
						       "pattern": "",
						       "type": "string"
						     },
						     "maxItems": 20,
						     "minItems": 1,
						     "type": "array"
						   }
						*/
						Description: "List of Signing profile version Arns",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
					},
				},
			),
			Required: true,
		},
		"code_signing_config_arn": {
			// Property: CodeSigningConfigArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A unique Arn for CodeSigningConfig resource",
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "A unique Arn for CodeSigningConfig resource",
			Type:        types.StringType,
			Computed:    true,
		},
		"code_signing_config_id": {
			// Property: CodeSigningConfigId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A unique identifier for CodeSigningConfig resource",
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "A unique identifier for CodeSigningConfig resource",
			Type:        types.StringType,
			Computed:    true,
		},
		"code_signing_policies": {
			// Property: CodeSigningPolicies
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "Policies to control how to act if a signature is invalid",
			     "properties": {
			       "UntrustedArtifactOnDeployment": {
			         "description": "Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided",
			         "enum": [
			           "Warn",
			           "Enforce"
			         ],
			         "type": "string"
			       }
			     },
			     "$ref": "#/definitions/CodeSigningPolicies",
			     "required": [
			       "UntrustedArtifactOnDeployment"
			     ],
			     "type": "object"
			   }
			*/
			Description: "Policies to control how to act if a signature is invalid",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"untrusted_artifact_on_deployment": {
						// Property: UntrustedArtifactOnDeployment
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided",
						     "enum": [
						       "Warn",
						       "Enforce"
						     ],
						     "type": "string"
						   }
						*/
						Description: "Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A description of the CodeSigningConfig",
			     "maxLength": 256,
			     "minLength": 0,
			     "type": "string"
			   }
			*/
			Description: "A description of the CodeSigningConfig",
			Type:        types.StringType,
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::Lambda::CodeSigningConfig.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::CodeSigningConfig").WithTerraformTypeName("aws_lambda_code_signing_config").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_lambda_code_signing_config", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
