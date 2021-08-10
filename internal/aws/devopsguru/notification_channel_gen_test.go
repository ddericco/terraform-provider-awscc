// Code generated by generators/resource/main.go; DO NOT EDIT.

package devopsguru_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSDevOpsGuruNotificationChannel_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DevOpsGuru::NotificationChannel", "aws_devopsguru_notification_channel", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
