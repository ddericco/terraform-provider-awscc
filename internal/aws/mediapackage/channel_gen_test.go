// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediapackage_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSMediaPackageChannel_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MediaPackage::Channel", "aws_mediapackage_channel", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
