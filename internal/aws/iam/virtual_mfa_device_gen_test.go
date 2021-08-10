// Code generated by generators/resource/main.go; DO NOT EDIT.

package iam_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSIAMVirtualMFADevice_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IAM::VirtualMFADevice", "aws_iam_virtual_mfa_device", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
