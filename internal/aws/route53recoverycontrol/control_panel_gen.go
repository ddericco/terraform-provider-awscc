// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoverycontrol

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_route53recoverycontrol_control_panel", controlPanelResourceType)
}

// controlPanelResourceType returns the Terraform awscc_route53recoverycontrol_control_panel resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53RecoveryControl::ControlPanel resource type.
func controlPanelResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cluster_arn": {
			// Property: ClusterArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Cluster to associate with the Control Panel",
			//   "type": "string"
			// }
			Description: "Cluster to associate with the Control Panel",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ClusterArn is a force-new attribute.
		},
		"control_panel_arn": {
			// Property: ControlPanelArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the cluster.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the cluster.",
			Type:        types.StringType,
			Computed:    true,
		},
		"default_control_panel": {
			// Property: DefaultControlPanel
			// CloudFormation resource type schema:
			// {
			//   "description": "A flag that Amazon Route 53 Application Recovery Controller sets to true to designate the default control panel for a cluster. When you create a cluster, Amazon Route 53 Application Recovery Controller creates a control panel, and sets this flag for that control panel. If you create a control panel yourself, this flag is set to false.",
			//   "type": "boolean"
			// }
			Description: "A flag that Amazon Route 53 Application Recovery Controller sets to true to designate the default control panel for a cluster. When you create a cluster, Amazon Route 53 Application Recovery Controller creates a control panel, and sets this flag for that control panel. If you create a control panel yourself, this flag is set to false.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the control panel. You can use any non-white space character in the name.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the control panel. You can use any non-white space character in the name.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLength(1, 64)},
			Required:    true,
		},
		"routing_control_count": {
			// Property: RoutingControlCount
			// CloudFormation resource type schema:
			// {
			//   "description": "Count of associated routing controls",
			//   "type": "integer"
			// }
			Description: "Count of associated routing controls",
			Type:        types.NumberType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
			//   "enum": [
			//     "PENDING",
			//     "DEPLOYED",
			//     "PENDING_DELETION"
			//   ],
			//   "type": "string"
			// }
			Description: "The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "AWS Route53 Recovery Control Control Panel resource schema .",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryControl::ControlPanel").WithTerraformTypeName("awscc_route53recoverycontrol_control_panel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cluster_arn":           "ClusterArn",
		"control_panel_arn":     "ControlPanelArn",
		"default_control_panel": "DefaultControlPanel",
		"name":                  "Name",
		"routing_control_count": "RoutingControlCount",
		"status":                "Status",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53recoverycontrol_control_panel", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
