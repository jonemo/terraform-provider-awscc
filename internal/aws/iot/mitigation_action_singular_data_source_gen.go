// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_mitigation_action", mitigationActionDataSource)
}

// mitigationActionDataSource returns the Terraform awscc_iot_mitigation_action data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::MitigationAction resource.
func mitigationActionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ActionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for the mitigation action.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9:_-]+",
		//	  "type": "string"
		//	}
		"action_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for the mitigation action.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ActionParams
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The set of parameters for this mitigation action. You can specify only one type of parameter (in other words, you can apply only one action for each defined mitigation action).",
		//	  "properties": {
		//	    "AddThingsToThingGroupParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters to define a mitigation action that moves devices associated with a certificate to one or more specified thing groups, typically for quarantine.",
		//	      "properties": {
		//	        "OverrideDynamicGroups": {
		//	          "description": "Specifies if this mitigation action can move the things that triggered the mitigation action out of one or more dynamic thing groups.",
		//	          "type": "boolean"
		//	        },
		//	        "ThingGroupNames": {
		//	          "description": "The list of groups to which you want to add the things that triggered the mitigation action.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "maxItems": 10,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "required": [
		//	        "ThingGroupNames"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "EnableIoTLoggingParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters to define a mitigation action that enables AWS IoT logging at a specified level of detail.",
		//	      "properties": {
		//	        "LogLevel": {
		//	          "description": " Specifies which types of information are logged.",
		//	          "enum": [
		//	            "DEBUG",
		//	            "INFO",
		//	            "ERROR",
		//	            "WARN"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "RoleArnForLogging": {
		//	          "description": " The ARN of the IAM role used for logging.",
		//	          "maxLength": 2048,
		//	          "minLength": 20,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "LogLevel",
		//	        "RoleArnForLogging"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "PublishFindingToSnsParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters, to define a mitigation action that publishes findings to Amazon SNS. You can implement your own custom actions in response to the Amazon SNS messages.",
		//	      "properties": {
		//	        "TopicArn": {
		//	          "description": "The ARN of the topic to which you want to publish the findings.",
		//	          "maxLength": 2048,
		//	          "minLength": 20,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "TopicArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ReplaceDefaultPolicyVersionParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters to define a mitigation action that adds a blank policy to restrict permissions.",
		//	      "properties": {
		//	        "TemplateName": {
		//	          "enum": [
		//	            "BLANK_POLICY"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "TemplateName"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "UpdateCACertificateParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters to define a mitigation action that changes the state of the CA certificate to inactive.",
		//	      "properties": {
		//	        "Action": {
		//	          "enum": [
		//	            "DEACTIVATE"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Action"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "UpdateDeviceCertificateParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters to define a mitigation action that changes the state of the device certificate to inactive.",
		//	      "properties": {
		//	        "Action": {
		//	          "enum": [
		//	            "DEACTIVATE"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Action"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"action_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AddThingsToThingGroupParams
				"add_things_to_thing_group_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: OverrideDynamicGroups
						"override_dynamic_groups": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies if this mitigation action can move the things that triggered the mitigation action out of one or more dynamic thing groups.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ThingGroupNames
						"thing_group_names": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The list of groups to which you want to add the things that triggered the mitigation action.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters to define a mitigation action that moves devices associated with a certificate to one or more specified thing groups, typically for quarantine.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: EnableIoTLoggingParams
				"enable_io_t_logging_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: LogLevel
						"log_level": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: " Specifies which types of information are logged.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: RoleArnForLogging
						"role_arn_for_logging": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: " The ARN of the IAM role used for logging.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters to define a mitigation action that enables AWS IoT logging at a specified level of detail.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PublishFindingToSnsParams
				"publish_finding_to_sns_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: TopicArn
						"topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the topic to which you want to publish the findings.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters, to define a mitigation action that publishes findings to Amazon SNS. You can implement your own custom actions in response to the Amazon SNS messages.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ReplaceDefaultPolicyVersionParams
				"replace_default_policy_version_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: TemplateName
						"template_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters to define a mitigation action that adds a blank policy to restrict permissions.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UpdateCACertificateParams
				"update_ca_certificate_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Action
						"action": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters to define a mitigation action that changes the state of the CA certificate to inactive.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UpdateDeviceCertificateParams
				"update_device_certificate_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Action
						"action": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters to define a mitigation action that changes the state of the device certificate to inactive.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The set of parameters for this mitigation action. You can specify only one type of parameter (in other words, you can apply only one action for each defined mitigation action).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MitigationActionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"mitigation_action_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MitigationActionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"mitigation_action_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
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
		//	        "description": "The tag's key.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag's value.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's key.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's value.",
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
		Description: "Data Source schema for AWS::IoT::MitigationAction",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::MitigationAction").WithTerraformTypeName("awscc_iot_mitigation_action")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                                "Action",
		"action_name":                           "ActionName",
		"action_params":                         "ActionParams",
		"add_things_to_thing_group_params":      "AddThingsToThingGroupParams",
		"enable_io_t_logging_params":            "EnableIoTLoggingParams",
		"key":                                   "Key",
		"log_level":                             "LogLevel",
		"mitigation_action_arn":                 "MitigationActionArn",
		"mitigation_action_id":                  "MitigationActionId",
		"override_dynamic_groups":               "OverrideDynamicGroups",
		"publish_finding_to_sns_params":         "PublishFindingToSnsParams",
		"replace_default_policy_version_params": "ReplaceDefaultPolicyVersionParams",
		"role_arn":                              "RoleArn",
		"role_arn_for_logging":                  "RoleArnForLogging",
		"tags":                                  "Tags",
		"template_name":                         "TemplateName",
		"thing_group_names":                     "ThingGroupNames",
		"topic_arn":                             "TopicArn",
		"update_ca_certificate_params":          "UpdateCACertificateParams",
		"update_device_certificate_params":      "UpdateDeviceCertificateParams",
		"value":                                 "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
