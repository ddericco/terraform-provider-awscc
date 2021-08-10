// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediapackage_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSMediaPackagePackagingConfiguration_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MediaPackage::PackagingConfiguration", "aws_mediapackage_packaging_configuration", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
