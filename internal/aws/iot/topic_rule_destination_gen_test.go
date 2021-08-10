// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSIoTTopicRuleDestination_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoT::TopicRuleDestination", "aws_iot_topic_rule_destination", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
	})
}

func TestAccAWSIoTTopicRuleDestination_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoT::TopicRuleDestination", "aws_iot_topic_rule_destination", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
