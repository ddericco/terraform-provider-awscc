// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSIoTSiteWiseAccessPolicy_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoTSiteWise::AccessPolicy", "aws_iotsitewise_access_policy", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
