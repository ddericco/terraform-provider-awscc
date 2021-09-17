// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_cloudformation_public_type_versions", publicTypeVersionsDataSourceType)
}

// publicTypeVersionsDataSourceType returns the Terraform awscc_cloudformation_public_type_versions data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CloudFormation::PublicTypeVersion resource type.
func publicTypeVersionsDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			Description: "Uniquely identifies the data source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ids": {
			Description: "Set of Resource Identifiers.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Plural Data Source schema for AWS::CloudFormation::PublicTypeVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::PublicTypeVersion").WithTerraformTypeName("awscc_cloudformation_public_type_versions")
	opts = opts.WithTerraformSchema(schema)

	pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return pluralDataSourceType, nil
}
