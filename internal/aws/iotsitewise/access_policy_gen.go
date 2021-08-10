// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise

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
	registry.AddResourceTypeFactory("aws_iotsitewise_access_policy", accessPolicy)
}

// accessPolicy returns the Terraform aws_iotsitewise_access_policy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTSiteWise::AccessPolicy resource type.
func accessPolicy(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"access_policy_arn": {
			// Property: AccessPolicyArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the access policy.",
			     "type": "string"
			   }
			*/
			Description: "The ARN of the access policy.",
			Type:        types.StringType,
			Computed:    true,
		},
		"access_policy_id": {
			// Property: AccessPolicyId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the access policy.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the access policy.",
			Type:        types.StringType,
			Computed:    true,
		},
		"access_policy_identity": {
			// Property: AccessPolicyIdentity
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The identity for this access policy. Choose either an SSO user or group or an IAM user or role.",
			     "properties": {
			       "IamRole": {
			         "additionalProperties": false,
			         "description": "Contains information for an IAM role identity in an access policy.",
			         "properties": {
			           "arn": {
			             "description": "The ARN of the IAM role.",
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/IamRole",
			         "type": "object"
			       },
			       "IamUser": {
			         "additionalProperties": false,
			         "description": "Contains information for an IAM user identity in an access policy.",
			         "properties": {
			           "arn": {
			             "description": "The ARN of the IAM user.",
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/IamUser",
			         "type": "object"
			       },
			       "User": {
			         "additionalProperties": false,
			         "description": "Contains information for a user identity in an access policy.",
			         "properties": {
			           "id": {
			             "description": "The AWS SSO ID of the user.",
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/User",
			         "type": "object"
			       }
			     },
			     "$ref": "#/definitions/AccessPolicyIdentity",
			     "type": "object"
			   }
			*/
			Description: "The identity for this access policy. Choose either an SSO user or group or an IAM user or role.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"iam_role": {
						// Property: IamRole
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Contains information for an IAM role identity in an access policy.",
						     "properties": {
						       "arn": {
						         "description": "The ARN of the IAM role.",
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/IamRole",
						     "type": "object"
						   }
						*/
						Description: "Contains information for an IAM role identity in an access policy.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"arn": {
									// Property: arn
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The ARN of the IAM role.",
									     "type": "string"
									   }
									*/
									Description: "The ARN of the IAM role.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"iam_user": {
						// Property: IamUser
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Contains information for an IAM user identity in an access policy.",
						     "properties": {
						       "arn": {
						         "description": "The ARN of the IAM user.",
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/IamUser",
						     "type": "object"
						   }
						*/
						Description: "Contains information for an IAM user identity in an access policy.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"arn": {
									// Property: arn
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The ARN of the IAM user.",
									     "type": "string"
									   }
									*/
									Description: "The ARN of the IAM user.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"user": {
						// Property: User
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Contains information for a user identity in an access policy.",
						     "properties": {
						       "id": {
						         "description": "The AWS SSO ID of the user.",
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/User",
						     "type": "object"
						   }
						*/
						Description: "Contains information for a user identity in an access policy.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"id": {
									// Property: id
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The AWS SSO ID of the user.",
									     "type": "string"
									   }
									*/
									Description: "The AWS SSO ID of the user.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"access_policy_permission": {
			// Property: AccessPolicyPermission
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The permission level for this access policy. Valid values are ADMINISTRATOR or VIEWER.",
			     "type": "string"
			   }
			*/
			Description: "The permission level for this access policy. Valid values are ADMINISTRATOR or VIEWER.",
			Type:        types.StringType,
			Required:    true,
		},
		"access_policy_resource": {
			// Property: AccessPolicyResource
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The AWS IoT SiteWise Monitor resource for this access policy. Choose either portal or project but not both.",
			     "properties": {
			       "Portal": {
			         "additionalProperties": false,
			         "description": "A portal resource.",
			         "properties": {
			           "id": {
			             "description": "The ID of the portal.",
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/Portal",
			         "type": "object"
			       },
			       "Project": {
			         "additionalProperties": false,
			         "description": "A project resource.",
			         "properties": {
			           "id": {
			             "description": "The ID of the project.",
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/Project",
			         "type": "object"
			       }
			     },
			     "$ref": "#/definitions/AccessPolicyResource",
			     "type": "object"
			   }
			*/
			Description: "The AWS IoT SiteWise Monitor resource for this access policy. Choose either portal or project but not both.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"portal": {
						// Property: Portal
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "A portal resource.",
						     "properties": {
						       "id": {
						         "description": "The ID of the portal.",
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/Portal",
						     "type": "object"
						   }
						*/
						Description: "A portal resource.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"id": {
									// Property: id
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The ID of the portal.",
									     "type": "string"
									   }
									*/
									Description: "The ID of the portal.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"project": {
						// Property: Project
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "A project resource.",
						     "properties": {
						       "id": {
						         "description": "The ID of the project.",
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/Project",
						     "type": "object"
						   }
						*/
						Description: "A project resource.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"id": {
									// Property: id
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The ID of the project.",
									     "type": "string"
									   }
									*/
									Description: "The ID of the project.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
			),
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
		Description: "Resource schema for AWS::IoTSiteWise::AccessPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::AccessPolicy").WithTerraformTypeName("aws_iotsitewise_access_policy").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_iotsitewise_access_policy", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
