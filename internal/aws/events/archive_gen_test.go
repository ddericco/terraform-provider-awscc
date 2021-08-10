// Code generated by generators/resource/main.go; DO NOT EDIT.

package events_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSEventsArchive_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Events::Archive", "aws_events_archive", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
