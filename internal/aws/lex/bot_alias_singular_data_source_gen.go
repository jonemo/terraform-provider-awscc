// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lex_bot_alias", botAliasDataSource)
}

// botAliasDataSource returns the Terraform awscc_lex_bot_alias data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lex::BotAlias resource.
func botAliasDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BotAliasId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique ID of resource",
		//	  "maxLength": 10,
		//	  "minLength": 10,
		//	  "pattern": "^[0-9a-zA-Z]+$",
		//	  "type": "string"
		//	}
		"bot_alias_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique ID of resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BotAliasLocaleSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of bot alias locale settings to add to the bot alias.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A locale setting in alias",
		//	    "properties": {
		//	      "BotAliasLocaleSetting": {
		//	        "additionalProperties": false,
		//	        "description": "You can use this parameter to specify a specific Lambda function to run different functions in different locales.",
		//	        "properties": {
		//	          "CodeHookSpecification": {
		//	            "additionalProperties": false,
		//	            "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
		//	            "properties": {
		//	              "LambdaCodeHook": {
		//	                "additionalProperties": false,
		//	                "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
		//	                "properties": {
		//	                  "CodeHookInterfaceVersion": {
		//	                    "description": "The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.",
		//	                    "maxLength": 5,
		//	                    "minLength": 1,
		//	                    "type": "string"
		//	                  },
		//	                  "LambdaArn": {
		//	                    "description": "The Amazon Resource Name (ARN) of the Lambda function.",
		//	                    "maxLength": 2048,
		//	                    "minLength": 20,
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "CodeHookInterfaceVersion",
		//	                  "LambdaArn"
		//	                ],
		//	                "type": "object"
		//	              }
		//	            },
		//	            "required": [
		//	              "LambdaCodeHook"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "Enabled": {
		//	            "description": "Whether the Lambda code hook is enabled",
		//	            "type": "boolean"
		//	          }
		//	        },
		//	        "required": [
		//	          "Enabled"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "LocaleId": {
		//	        "description": "A string used to identify the locale",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "LocaleId",
		//	      "BotAliasLocaleSetting"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"bot_alias_locale_settings": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: BotAliasLocaleSetting
					"bot_alias_locale_setting": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: CodeHookSpecification
							"code_hook_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: LambdaCodeHook
									"lambda_code_hook": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: CodeHookInterfaceVersion
											"code_hook_interface_version": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: LambdaArn
											"lambda_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The Amazon Resource Name (ARN) of the Lambda function.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "Contains information about code hooks that Amazon Lex calls during a conversation.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Contains information about code hooks that Amazon Lex calls during a conversation.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Enabled
							"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Whether the Lambda code hook is enabled",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "You can use this parameter to specify a specific Lambda function to run different functions in different locales.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: LocaleId
					"locale_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string used to identify the locale",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of bot alias locale settings to add to the bot alias.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BotAliasName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for a resource.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^([0-9a-zA-Z][_-]?)+$",
		//	  "type": "string"
		//	}
		"bot_alias_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BotAliasStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "Creating",
		//	    "Available",
		//	    "Deleting",
		//	    "Failed"
		//	  ],
		//	  "type": "string"
		//	}
		"bot_alias_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BotAliasTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags to add to the bot alias.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A label for tagging Lex resources",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A string used to identify this tag",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A string containing the value for the tag",
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
		"bot_alias_tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string used to identify this tag",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string containing the value for the tag",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of tags to add to the bot alias.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BotId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique ID of resource",
		//	  "maxLength": 10,
		//	  "minLength": 10,
		//	  "pattern": "^[0-9a-zA-Z]+$",
		//	  "type": "string"
		//	}
		"bot_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique ID of resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BotVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of a bot.",
		//	  "maxLength": 5,
		//	  "minLength": 1,
		//	  "pattern": "^(DRAFT|[0-9]+)$",
		//	  "type": "string"
		//	}
		"bot_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version of a bot.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConversationLogSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
		//	  "properties": {
		//	    "AudioLogSettings": {
		//	      "description": "List of audio log settings",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Settings for logging audio of conversations between Amazon Lex and a user. You specify whether to log audio and the Amazon S3 bucket where the audio file is stored.",
		//	        "properties": {
		//	          "Destination": {
		//	            "additionalProperties": false,
		//	            "description": "The location of audio log files collected when conversation logging is enabled for a bot.",
		//	            "properties": {
		//	              "S3Bucket": {
		//	                "additionalProperties": false,
		//	                "description": "Specifies an Amazon S3 bucket for logging audio conversations",
		//	                "properties": {
		//	                  "KmsKeyArn": {
		//	                    "description": "The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an S3 bucket.",
		//	                    "maxLength": 2048,
		//	                    "minLength": 20,
		//	                    "pattern": "^arn:[\\w\\-]+:kms:[\\w\\-]+:[\\d]{12}:(?:key\\/[\\w\\-]+|alias\\/[a-zA-Z0-9:\\/_\\-]{1,256})$",
		//	                    "type": "string"
		//	                  },
		//	                  "LogPrefix": {
		//	                    "description": "The Amazon S3 key of the deployment package.",
		//	                    "maxLength": 1024,
		//	                    "minLength": 0,
		//	                    "type": "string"
		//	                  },
		//	                  "S3BucketArn": {
		//	                    "description": "The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.",
		//	                    "maxLength": 2048,
		//	                    "minLength": 1,
		//	                    "pattern": "^arn:[\\w\\-]+:s3:::[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$",
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "LogPrefix",
		//	                  "S3BucketArn"
		//	                ],
		//	                "type": "object"
		//	              }
		//	            },
		//	            "required": [
		//	              "S3Bucket"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "Enabled": {
		//	            "description": "",
		//	            "type": "boolean"
		//	          }
		//	        },
		//	        "required": [
		//	          "Destination",
		//	          "Enabled"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "TextLogSettings": {
		//	      "description": "List of text log settings",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
		//	        "properties": {
		//	          "Destination": {
		//	            "additionalProperties": false,
		//	            "description": "Defines the Amazon CloudWatch Logs destination log group for conversation text logs.",
		//	            "properties": {
		//	              "CloudWatch": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "CloudWatchLogGroupArn": {
		//	                    "description": "A string used to identify the groupArn for the Cloudwatch Log Group",
		//	                    "maxLength": 2048,
		//	                    "minLength": 1,
		//	                    "type": "string"
		//	                  },
		//	                  "LogPrefix": {
		//	                    "description": "A string containing the value for the Log Prefix",
		//	                    "maxLength": 1024,
		//	                    "minLength": 0,
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "CloudWatchLogGroupArn",
		//	                  "LogPrefix"
		//	                ],
		//	                "type": "object"
		//	              }
		//	            },
		//	            "required": [
		//	              "CloudWatch"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "Enabled": {
		//	            "description": "",
		//	            "type": "boolean"
		//	          }
		//	        },
		//	        "required": [
		//	          "Destination",
		//	          "Enabled"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"conversation_log_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AudioLogSettings
				"audio_log_settings": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Destination
							"destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: S3Bucket
									"s3_bucket": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: KmsKeyArn
											"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an S3 bucket.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: LogPrefix
											"log_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The Amazon S3 key of the deployment package.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: S3BucketArn
											"s3_bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "Specifies an Amazon S3 bucket for logging audio conversations",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The location of audio log files collected when conversation logging is enabled for a bot.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Enabled
							"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of audio log settings",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: TextLogSettings
				"text_log_settings": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Destination
							"destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: CloudWatch
									"cloudwatch": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: CloudWatchLogGroupArn
											"cloudwatch_log_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "A string used to identify the groupArn for the Cloudwatch Log Group",
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: LogPrefix
											"log_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "A string containing the value for the Log Prefix",
												Computed:    true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Defines the Amazon CloudWatch Logs destination log group for conversation text logs.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Enabled
							"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of text log settings",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Contains information about code hooks that Amazon Lex calls during a conversation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the bot alias. Use the description to help identify the bot alias in lists.",
		//	  "maxLength": 200,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the bot alias. Use the description to help identify the bot alias in lists.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SentimentAnalysisSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.",
		//	  "properties": {
		//	    "DetectSentiment": {
		//	      "description": "Enable to call Amazon Comprehend for Sentiment natively within Lex",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "required": [
		//	    "DetectSentiment"
		//	  ],
		//	  "type": "object"
		//	}
		"sentiment_analysis_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DetectSentiment
				"detect_sentiment": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Enable to call Amazon Comprehend for Sentiment natively within Lex",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Lex::BotAlias",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lex::BotAlias").WithTerraformTypeName("awscc_lex_bot_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"audio_log_settings":          "AudioLogSettings",
		"bot_alias_id":                "BotAliasId",
		"bot_alias_locale_setting":    "BotAliasLocaleSetting",
		"bot_alias_locale_settings":   "BotAliasLocaleSettings",
		"bot_alias_name":              "BotAliasName",
		"bot_alias_status":            "BotAliasStatus",
		"bot_alias_tags":              "BotAliasTags",
		"bot_id":                      "BotId",
		"bot_version":                 "BotVersion",
		"cloudwatch":                  "CloudWatch",
		"cloudwatch_log_group_arn":    "CloudWatchLogGroupArn",
		"code_hook_interface_version": "CodeHookInterfaceVersion",
		"code_hook_specification":     "CodeHookSpecification",
		"conversation_log_settings":   "ConversationLogSettings",
		"description":                 "Description",
		"destination":                 "Destination",
		"detect_sentiment":            "DetectSentiment",
		"enabled":                     "Enabled",
		"key":                         "Key",
		"kms_key_arn":                 "KmsKeyArn",
		"lambda_arn":                  "LambdaArn",
		"lambda_code_hook":            "LambdaCodeHook",
		"locale_id":                   "LocaleId",
		"log_prefix":                  "LogPrefix",
		"s3_bucket":                   "S3Bucket",
		"s3_bucket_arn":               "S3BucketArn",
		"sentiment_analysis_settings": "SentimentAnalysisSettings",
		"text_log_settings":           "TextLogSettings",
		"value":                       "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
