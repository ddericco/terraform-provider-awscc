// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_apigatewayv2_api_mapping", apiMappingResource)
}

// apiMappingResource returns the Terraform awscc_apigatewayv2_api_mapping resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGatewayV2::ApiMapping resource.
func apiMappingResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the API.",
		//	  "type": "string"
		//	}
		"api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the API.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ApiMappingId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"api_mapping_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ApiMappingKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The API mapping key.",
		//	  "type": "string"
		//	}
		"api_mapping_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The API mapping key.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DomainName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The domain name.",
		//	  "type": "string"
		//	}
		"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The domain name.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Stage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The API stage.",
		//	  "type": "string"
		//	}
		"stage": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The API stage.",
			Required:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The ``AWS::ApiGatewayV2::ApiMapping`` resource contains an API mapping. An API mapping relates a path of your custom domain name to a stage of your API. A custom domain name can have multiple API mappings, but the paths can't overlap. A custom domain can map only to APIs of the same protocol type. For more information, see [CreateApiMapping](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/domainnames-domainname-apimappings.html#CreateApiMapping) in the *Amazon API Gateway V2 API Reference*.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGatewayV2::ApiMapping").WithTerraformTypeName("awscc_apigatewayv2_api_mapping")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_id":          "ApiId",
		"api_mapping_id":  "ApiMappingId",
		"api_mapping_key": "ApiMappingKey",
		"domain_name":     "DomainName",
		"stage":           "Stage",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
