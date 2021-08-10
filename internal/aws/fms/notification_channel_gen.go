// Code generated by generators/resource/main.go; DO NOT EDIT.

package fms

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_fms_notification_channel", notificationChannel)
}

// notificationChannel returns the Terraform aws_fms_notification_channel resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::FMS::NotificationChannel resource type.
func notificationChannel(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"sns_role_name": {
			// Property: SnsRoleName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A resource ARN.",
			     "maxLength": 1024,
			     "minLength": 1,
			     "pattern": "",
			     "$ref": "#/definitions/ResourceArn",
			     "type": "string"
			   }
			*/
			Description: "A resource ARN.",
			Type:        types.StringType,
			Required:    true,
		},
		"sns_topic_arn": {
			// Property: SnsTopicArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A resource ARN.",
			     "maxLength": 1024,
			     "minLength": 1,
			     "pattern": "",
			     "$ref": "#/definitions/ResourceArn",
			     "type": "string"
			   }
			*/
			Description: "A resource ARN.",
			Type:        types.StringType,
			Required:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Designates the IAM role and Amazon Simple Notification Service (SNS) topic that AWS Firewall Manager uses to record SNS logs.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FMS::NotificationChannel").WithTerraformTypeName("aws_fms_notification_channel").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_fms_notification_channel", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
