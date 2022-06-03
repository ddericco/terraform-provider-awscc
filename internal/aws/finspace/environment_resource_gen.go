// Code generated by generators/resource/main.go; DO NOT EDIT.

package finspace

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_finspace_environment", environmentResourceType)
}

// environmentResourceType returns the Terraform awscc_finspace_environment resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::FinSpace::Environment resource type.
func environmentResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"aws_account_id": {
			// Property: AwsAccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "AWS account ID associated with the Environment",
			//   "pattern": "^[a-zA-Z0-9]{1,26}$",
			//   "type": "string"
			// }
			Description: "AWS account ID associated with the Environment",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"data_bundles": {
			// Property: DataBundles
			// CloudFormation resource type schema:
			// {
			//   "description": "ARNs of FinSpace Data Bundles to install",
			//   "items": {
			//     "pattern": "^arn:aws:finspace:[A-Za-z0-9_/.-]{0,63}:\\d*:data-bundle/[0-9A-Za-z_-]{1,128}$",
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "ARNs of FinSpace Data Bundles to install",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^arn:aws:finspace:[A-Za-z0-9_/.-]{0,63}:\\d*:data-bundle/[0-9A-Za-z_-]{1,128}$"), "")),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"dedicated_service_account_id": {
			// Property: DedicatedServiceAccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "ID for FinSpace created account used to store Environment artifacts",
			//   "pattern": "^[a-zA-Z0-9]{1,26}$",
			//   "type": "string"
			// }
			Description: "ID for FinSpace created account used to store Environment artifacts",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of the Environment",
			//   "pattern": "^[a-zA-Z0-9. ]{1,1000}$",
			//   "type": "string"
			// }
			Description: "Description of the Environment",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9. ]{1,1000}$"), ""),
			},
		},
		"environment_arn": {
			// Property: EnvironmentArn
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN of the Environment",
			//   "pattern": "^arn:aws:finspace:[A-Za-z0-9_/.-]{0,63}:\\d+:environment/[0-9A-Za-z_-]{1,128}$",
			//   "type": "string"
			// }
			Description: "ARN of the Environment",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"environment_id": {
			// Property: EnvironmentId
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique identifier for representing FinSpace Environment",
			//   "pattern": "^[a-zA-Z0-9]{1,26}$",
			//   "type": "string"
			// }
			Description: "Unique identifier for representing FinSpace Environment",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"environment_url": {
			// Property: EnvironmentUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "URL used to login to the Environment",
			//   "pattern": "^[-a-zA-Z0-9+\u0026amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\u0026amp;@#/%=~_|]{1,1000}",
			//   "type": "string"
			// }
			Description: "URL used to login to the Environment",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"federation_mode": {
			// Property: FederationMode
			// CloudFormation resource type schema:
			// {
			//   "description": "Federation mode used with the Environment",
			//   "enum": [
			//     "LOCAL",
			//     "FEDERATED"
			//   ],
			//   "type": "string"
			// }
			Description: "Federation mode used with the Environment",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"LOCAL",
					"FEDERATED",
				}),
			},
		},
		"federation_parameters": {
			// Property: FederationParameters
			// CloudFormation resource type schema:
			// {
			//   "description": "Additional parameters to identify Federation mode",
			//   "properties": {
			//     "ApplicationCallBackURL": {
			//       "description": "SAML metadata URL to link with the Environment",
			//       "pattern": "^https?://[-a-zA-Z0-9+\u0026amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\u0026amp;@#/%=~_|]{1,1000}",
			//       "type": "string"
			//     },
			//     "AttributeMap": {
			//       "description": "Attribute map for SAML configuration",
			//       "type": "object"
			//     },
			//     "FederationProviderName": {
			//       "description": "Federation provider name to link with the Environment",
			//       "maxLength": 32,
			//       "minLength": 1,
			//       "pattern": "[^_\\p{Z}][\\p{L}\\p{M}\\p{S}\\p{N}\\p{P}][^_\\p{Z}]+",
			//       "type": "string"
			//     },
			//     "FederationURN": {
			//       "description": "SAML metadata URL to link with the Environment",
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "SamlMetadataDocument": {
			//       "description": "SAML metadata document to link the federation provider to the Environment",
			//       "maxLength": 10000000,
			//       "minLength": 1000,
			//       "pattern": ".*",
			//       "type": "string"
			//     },
			//     "SamlMetadataURL": {
			//       "description": "SAML metadata URL to link with the Environment",
			//       "pattern": "^https?://[-a-zA-Z0-9+\u0026amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\u0026amp;@#/%=~_|]{1,1000}",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Additional parameters to identify Federation mode",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"application_call_back_url": {
						// Property: ApplicationCallBackURL
						Description: "SAML metadata URL to link with the Environment",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^https?://[-a-zA-Z0-9+&amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&amp;@#/%=~_|]{1,1000}"), ""),
						},
					},
					"attribute_map": {
						// Property: AttributeMap
						Description: "Attribute map for SAML configuration",
						Type:        JSONStringType,
						Optional:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							JSONStringType.AttributePlanModifier(),
						},
					},
					"federation_provider_name": {
						// Property: FederationProviderName
						Description: "Federation provider name to link with the Environment",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 32),
							validate.StringMatch(regexp.MustCompile("[^_\\p{Z}][\\p{L}\\p{M}\\p{S}\\p{N}\\p{P}][^_\\p{Z}]+"), ""),
						},
					},
					"federation_urn": {
						// Property: FederationURN
						Description: "SAML metadata URL to link with the Environment",
						Type:        types.StringType,
						Optional:    true,
					},
					"saml_metadata_document": {
						// Property: SamlMetadataDocument
						Description: "SAML metadata document to link the federation provider to the Environment",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1000, 10000000),
							validate.StringMatch(regexp.MustCompile(".*"), ""),
						},
					},
					"saml_metadata_url": {
						// Property: SamlMetadataURL
						Description: "SAML metadata URL to link with the Environment",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^https?://[-a-zA-Z0-9+&amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&amp;@#/%=~_|]{1,1000}"), ""),
						},
					},
				},
			),
			Optional: true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "KMS key used to encrypt customer data within FinSpace Environment infrastructure",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "KMS key used to encrypt customer data within FinSpace Environment infrastructure",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the Environment",
			//   "pattern": "^[a-zA-Z0-9]+[a-zA-Z0-9-]*[a-zA-Z0-9]{1,255}$",
			//   "type": "string"
			// }
			Description: "Name of the Environment",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]+[a-zA-Z0-9-]*[a-zA-Z0-9]{1,255}$"), ""),
			},
		},
		"sage_maker_studio_domain_url": {
			// Property: SageMakerStudioDomainUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "SageMaker Studio Domain URL associated with the Environment",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "SageMaker Studio Domain URL associated with the Environment",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "State of the Environment",
			//   "enum": [
			//     "CREATE_REQUESTED",
			//     "CREATING",
			//     "CREATED",
			//     "DELETE_REQUESTED",
			//     "DELETING",
			//     "DELETED",
			//     "FAILED_CREATION",
			//     "FAILED_DELETION",
			//     "RETRY_DELETION",
			//     "SUSPENDED"
			//   ],
			//   "type": "string"
			// }
			Description: "State of the Environment",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"superuser_parameters": {
			// Property: SuperuserParameters
			// CloudFormation resource type schema:
			// {
			//   "description": "Parameters of the first Superuser for the FinSpace Environment",
			//   "properties": {
			//     "EmailAddress": {
			//       "description": "Email address",
			//       "maxLength": 128,
			//       "minLength": 1,
			//       "pattern": "[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+[.]+[A-Za-z]+",
			//       "type": "string"
			//     },
			//     "FirstName": {
			//       "description": "First name",
			//       "maxLength": 50,
			//       "minLength": 1,
			//       "pattern": "^[a-zA-Z0-9]{1,50}$",
			//       "type": "string"
			//     },
			//     "LastName": {
			//       "description": "Last name",
			//       "maxLength": 50,
			//       "minLength": 1,
			//       "pattern": "^[a-zA-Z0-9]{1,50}$",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Parameters of the first Superuser for the FinSpace Environment",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"email_address": {
						// Property: EmailAddress
						Description: "Email address",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
							validate.StringMatch(regexp.MustCompile("[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+[.]+[A-Za-z]+"), ""),
						},
					},
					"first_name": {
						// Property: FirstName
						Description: "First name",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 50),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]{1,50}$"), ""),
						},
					},
					"last_name": {
						// Property: LastName
						Description: "Last name",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 50),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]{1,50}$"), ""),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FinSpace::Environment").WithTerraformTypeName("awscc_finspace_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_call_back_url":    "ApplicationCallBackURL",
		"attribute_map":                "AttributeMap",
		"aws_account_id":               "AwsAccountId",
		"data_bundles":                 "DataBundles",
		"dedicated_service_account_id": "DedicatedServiceAccountId",
		"description":                  "Description",
		"email_address":                "EmailAddress",
		"environment_arn":              "EnvironmentArn",
		"environment_id":               "EnvironmentId",
		"environment_url":              "EnvironmentUrl",
		"federation_mode":              "FederationMode",
		"federation_parameters":        "FederationParameters",
		"federation_provider_name":     "FederationProviderName",
		"federation_urn":               "FederationURN",
		"first_name":                   "FirstName",
		"kms_key_id":                   "KmsKeyId",
		"last_name":                    "LastName",
		"name":                         "Name",
		"sage_maker_studio_domain_url": "SageMakerStudioDomainUrl",
		"saml_metadata_document":       "SamlMetadataDocument",
		"saml_metadata_url":            "SamlMetadataURL",
		"status":                       "Status",
		"superuser_parameters":         "SuperuserParameters",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
