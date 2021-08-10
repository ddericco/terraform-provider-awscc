// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSSageMakerPipeline_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SageMaker::Pipeline", "aws_sagemaker_pipeline", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
