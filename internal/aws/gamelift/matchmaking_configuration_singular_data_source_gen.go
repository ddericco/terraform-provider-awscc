// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_gamelift_matchmaking_configuration", matchmakingConfigurationDataSource)
}

// matchmakingConfigurationDataSource returns the Terraform awscc_gamelift_matchmaking_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::GameLift::MatchmakingConfiguration resource.
func matchmakingConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AcceptanceRequired
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A flag that indicates whether a match that was created with this configuration must be accepted by the matched players",
		//	  "type": "boolean"
		//	}
		"acceptance_required": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A flag that indicates whether a match that was created with this configuration must be accepted by the matched players",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AcceptanceTimeoutSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.",
		//	  "maximum": 600,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"acceptance_timeout_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AdditionalPlayerCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of player slots in a match to keep open for future players.",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"additional_player_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of player slots in a match to keep open for future players.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift matchmaking configuration resource and uniquely identifies it.",
		//	  "pattern": "^arn:.*:matchmakingconfiguration\\/[a-zA-Z0-9-\\.]*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift matchmaking configuration resource and uniquely identifies it.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BackfillMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The method used to backfill game sessions created with this matchmaking configuration.",
		//	  "enum": [
		//	    "AUTOMATIC",
		//	    "MANUAL"
		//	  ],
		//	  "type": "string"
		//	}
		"backfill_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The method used to backfill game sessions created with this matchmaking configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A time stamp indicating when this data object was created.",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A time stamp indicating when this data object was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CustomEventData
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Information to attach to all events related to the matchmaking configuration.",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"custom_event_data": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Information to attach to all events related to the matchmaking configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A descriptive label that is associated with matchmaking configuration.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A descriptive label that is associated with matchmaking configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FlexMatchMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether this matchmaking configuration is being used with Amazon GameLift hosting or as a standalone matchmaking solution.",
		//	  "enum": [
		//	    "STANDALONE",
		//	    "WITH_QUEUE"
		//	  ],
		//	  "type": "string"
		//	}
		"flex_match_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether this matchmaking configuration is being used with Amazon GameLift hosting or as a standalone matchmaking solution.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GameProperties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A set of custom properties for a game session, formatted as key:value pairs.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair that contains information about a game session.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The game property identifier.",
		//	        "maxLength": 32,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The game property value.",
		//	        "maxLength": 96,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 16,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"game_properties": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The game property identifier.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The game property value.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A set of custom properties for a game session, formatted as key:value pairs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GameSessionData
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A set of custom game session properties, formatted as a single string value.",
		//	  "maxLength": 4096,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"game_session_data": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A set of custom game session properties, formatted as a single string value.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GameSessionQueueArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 256,
		//	    "minLength": 1,
		//	    "pattern": "[a-zA-Z0-9:/-]+",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"game_session_queue_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for the matchmaking configuration.",
		//	  "maxLength": 128,
		//	  "pattern": "[a-zA-Z0-9-\\.]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for the matchmaking configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NotificationTarget
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An SNS topic ARN that is set up to receive matchmaking notifications.",
		//	  "maxLength": 300,
		//	  "minLength": 0,
		//	  "pattern": "[a-zA-Z0-9:_/-]*(.fifo)?",
		//	  "type": "string"
		//	}
		"notification_target": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An SNS topic ARN that is set up to receive matchmaking notifications.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RequestTimeoutSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.",
		//	  "maximum": 43200,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"request_timeout_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RuleSetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) associated with the GameLift matchmaking rule set resource that this configuration uses.",
		//	  "pattern": "^arn:.*:matchmakingruleset\\/[a-zA-Z0-9-\\.]*",
		//	  "type": "string"
		//	}
		"rule_set_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) associated with the GameLift matchmaking rule set resource that this configuration uses.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RuleSetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for the matchmaking rule set to use with this configuration.",
		//	  "maxLength": 128,
		//	  "pattern": "[a-zA-Z0-9-\\.]*",
		//	  "type": "string"
		//	}
		"rule_set_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for the matchmaking rule set to use with this configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::GameLift::MatchmakingConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GameLift::MatchmakingConfiguration").WithTerraformTypeName("awscc_gamelift_matchmaking_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"acceptance_required":        "AcceptanceRequired",
		"acceptance_timeout_seconds": "AcceptanceTimeoutSeconds",
		"additional_player_count":    "AdditionalPlayerCount",
		"arn":                        "Arn",
		"backfill_mode":              "BackfillMode",
		"creation_time":              "CreationTime",
		"custom_event_data":          "CustomEventData",
		"description":                "Description",
		"flex_match_mode":            "FlexMatchMode",
		"game_properties":            "GameProperties",
		"game_session_data":          "GameSessionData",
		"game_session_queue_arns":    "GameSessionQueueArns",
		"key":                        "Key",
		"name":                       "Name",
		"notification_target":        "NotificationTarget",
		"request_timeout_seconds":    "RequestTimeoutSeconds",
		"rule_set_arn":               "RuleSetArn",
		"rule_set_name":              "RuleSetName",
		"tags":                       "Tags",
		"value":                      "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}