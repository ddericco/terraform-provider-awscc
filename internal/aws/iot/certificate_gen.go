// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iot_certificate", certificateResourceType)
}

// certificateResourceType returns the Terraform awscc_iot_certificate resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::Certificate resource type.
func certificateResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"ca_certificate_pem": {
			// Property: CACertificatePem
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 65536,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 65536)},
			Optional:   true,
			Computed:   true,
			// CACertificatePem is a force-new attribute.
			// CACertificatePem is a write-only attribute.
		},
		"certificate_mode": {
			// Property: CertificateMode
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DEFAULT",
			//     "SNI_ONLY"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// CertificateMode is a force-new attribute.
		},
		"certificate_pem": {
			// Property: CertificatePem
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 65536,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLength(1, 65536)},
			Optional:   true,
			Computed:   true,
			// CertificatePem is a force-new attribute.
		},
		"certificate_signing_request": {
			// Property: CertificateSigningRequest
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// CertificateSigningRequest is a force-new attribute.
			// CertificateSigningRequest is a write-only attribute.
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ACTIVE",
			//     "INACTIVE",
			//     "REVOKED",
			//     "PENDING_TRANSFER",
			//     "PENDING_ACTIVATION"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Use the AWS::IoT::Certificate resource to declare an AWS IoT X.509 certificate.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::Certificate").WithTerraformTypeName("awscc_iot_certificate")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"ca_certificate_pem":          "CACertificatePem",
		"certificate_mode":            "CertificateMode",
		"certificate_pem":             "CertificatePem",
		"certificate_signing_request": "CertificateSigningRequest",
		"id":                          "Id",
		"status":                      "Status",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/CertificateSigningRequest",
		"/properties/CACertificatePem",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iot_certificate", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
