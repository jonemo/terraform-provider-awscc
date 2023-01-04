// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudwatch_composite_alarm", compositeAlarmDataSource)
}

// compositeAlarmDataSource returns the Terraform awscc_cloudwatch_composite_alarm data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudWatch::CompositeAlarm resource.
func compositeAlarmDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ActionsEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
		//	  "type": "boolean"
		//	}
		"actions_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ActionsSuppressor
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Actions will be suppressed if the suppressor alarm is in the ALARM state. ActionsSuppressor can be an AlarmName or an Amazon Resource Name (ARN) from an existing alarm. ",
		//	  "maxLength": 1600,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"actions_suppressor": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Actions will be suppressed if the suppressor alarm is in the ALARM state. ActionsSuppressor can be an AlarmName or an Amazon Resource Name (ARN) from an existing alarm. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ActionsSuppressorExtensionPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Actions will be suppressed if WaitPeriod is active. The length of time that actions are suppressed is in seconds.",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"actions_suppressor_extension_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Actions will be suppressed if WaitPeriod is active. The length of time that actions are suppressed is in seconds.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ActionsSuppressorWaitPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Actions will be suppressed if ExtensionPeriod is active. The length of time that actions are suppressed is in seconds.",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"actions_suppressor_wait_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Actions will be suppressed if ExtensionPeriod is active. The length of time that actions are suppressed is in seconds.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AlarmActions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).",
		//	  "items": {
		//	    "description": "Amazon Resource Name (ARN) of the action",
		//	    "maxLength": 1024,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 5,
		//	  "type": "array"
		//	}
		"alarm_actions": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AlarmDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the alarm",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"alarm_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the alarm",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AlarmName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Composite Alarm",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"alarm_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Composite Alarm",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AlarmRule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Expression which aggregates the state of other Alarms (Metric or Composite Alarms)",
		//	  "maxLength": 10240,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"alarm_rule": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Expression which aggregates the state of other Alarms (Metric or Composite Alarms)",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the alarm",
		//	  "maxLength": 1600,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the alarm",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InsufficientDataActions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
		//	  "items": {
		//	    "description": "Amazon Resource Name (ARN) of the action",
		//	    "maxLength": 1024,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 5,
		//	  "type": "array"
		//	}
		"insufficient_data_actions": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OKActions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
		//	  "items": {
		//	    "description": "Amazon Resource Name (ARN) of the action",
		//	    "maxLength": 1024,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 5,
		//	  "type": "array"
		//	}
		"ok_actions": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudWatch::CompositeAlarm",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudWatch::CompositeAlarm").WithTerraformTypeName("awscc_cloudwatch_composite_alarm")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions_enabled":                     "ActionsEnabled",
		"actions_suppressor":                  "ActionsSuppressor",
		"actions_suppressor_extension_period": "ActionsSuppressorExtensionPeriod",
		"actions_suppressor_wait_period":      "ActionsSuppressorWaitPeriod",
		"alarm_actions":                       "AlarmActions",
		"alarm_description":                   "AlarmDescription",
		"alarm_name":                          "AlarmName",
		"alarm_rule":                          "AlarmRule",
		"arn":                                 "Arn",
		"insufficient_data_actions":           "InsufficientDataActions",
		"ok_actions":                          "OKActions",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
