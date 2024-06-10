// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package securityhub_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSSecurityHubFindingAggregatorDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SecurityHub::FindingAggregator", "awscc_securityhub_finding_aggregator", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSSecurityHubFindingAggregatorDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SecurityHub::FindingAggregator", "awscc_securityhub_finding_aggregator", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}