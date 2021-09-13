// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package cassandra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_cassandra_tables", tablesDataSourceType)
}

// tablesDataSourceType returns the Terraform awscc_cassandra_tables data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Cassandra::Table resource type.
func tablesDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			Description: "Uniquely identifies the data source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ids": {
			Description: "Set of Resource Identifiers.",
			Type:        providertypes.SetType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Plural Data Source schema for AWS::Cassandra::Table",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Cassandra::Table").WithTerraformTypeName("awscc_cassandra_tables")
	opts = opts.WithTerraformSchema(schema)

	pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return pluralDataSourceType, nil
}
