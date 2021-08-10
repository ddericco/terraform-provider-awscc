// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSIoTCustomMetric_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoT::CustomMetric", "aws_iot_custom_metric", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
