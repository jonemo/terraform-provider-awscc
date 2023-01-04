// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mwaa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_mwaa_environment", environmentDataSource)
}

// environmentDataSource returns the Terraform awscc_mwaa_environment data source.
// This Terraform data source corresponds to the CloudFormation AWS::MWAA::Environment resource.
func environmentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AirflowConfigurationOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Key/value pairs representing Airflow configuration variables.\n    Keys are prefixed by their section:\n\n    [core]\n    dags_folder={AIRFLOW_HOME}/dags\n\n    Would be represented as\n\n    \"core.dags_folder\": \"{AIRFLOW_HOME}/dags\"",
		//	  "type": "object"
		//	}
		"airflow_configuration_options": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Key/value pairs representing Airflow configuration variables.\n    Keys are prefixed by their section:\n\n    [core]\n    dags_folder={AIRFLOW_HOME}/dags\n\n    Would be represented as\n\n    \"core.dags_folder\": \"{AIRFLOW_HOME}/dags\"",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AirflowVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Version of airflow to deploy to the environment.",
		//	  "maxLength": 32,
		//	  "pattern": "^[0-9a-z.]+$",
		//	  "type": "string"
		//	}
		"airflow_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Version of airflow to deploy to the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN for the MWAA environment.",
		//	  "maxLength": 1224,
		//	  "minLength": 1,
		//	  "pattern": "^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:airflow:[a-z0-9\\-]+:\\d{12}:environment/\\w+",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN for the MWAA environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DagS3Path
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents an S3 prefix relative to the root of an S3 bucket.",
		//	  "maxLength": 1024,
		//	  "pattern": ".*",
		//	  "type": "string"
		//	}
		"dag_s3_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents an S3 prefix relative to the root of an S3 bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentClass
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Templated configuration for airflow processes and backing infrastructure.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"environment_class": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Templated configuration for airflow processes and backing infrastructure.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ExecutionRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "IAM role to be used by tasks.",
		//	  "maxLength": 1224,
		//	  "pattern": "^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$",
		//	  "type": "string"
		//	}
		"execution_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "IAM role to be used by tasks.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for MWAA data encryption.\n\n    You can specify the CMK using any of the following:\n\n    Key ID. For example, key/1234abcd-12ab-34cd-56ef-1234567890ab.\n\n    Key alias. For example, alias/ExampleAlias.\n\n    Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.\n\n    Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.\n\n    AWS authenticates the CMK asynchronously. Therefore, if you specify an ID, alias, or ARN that is not valid, the action can appear to complete, but eventually fails.",
		//	  "maxLength": 1224,
		//	  "pattern": "^(((arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:kms:[a-z]{2}-[a-z]+-\\d:\\d+:)?key\\/)?[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}|(arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):kms:[a-z]{2}-[a-z]+-\\d:\\d+:)?alias/.+)$",
		//	  "type": "string"
		//	}
		"kms_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for MWAA data encryption.\n\n    You can specify the CMK using any of the following:\n\n    Key ID. For example, key/1234abcd-12ab-34cd-56ef-1234567890ab.\n\n    Key alias. For example, alias/ExampleAlias.\n\n    Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.\n\n    Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.\n\n    AWS authenticates the CMK asynchronously. Therefore, if you specify an ID, alias, or ARN that is not valid, the action can appear to complete, but eventually fails.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoggingConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Logging configuration for the environment.",
		//	  "properties": {
		//	    "DagProcessingLogs": {
		//	      "additionalProperties": false,
		//	      "description": "Logging configuration for a specific airflow component.",
		//	      "properties": {
		//	        "CloudWatchLogGroupArn": {
		//	          "description": "",
		//	          "maxLength": 1224,
		//	          "pattern": "^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:logs:[a-z0-9\\-]+:\\d{12}:log-group:\\w+",
		//	          "type": "string"
		//	        },
		//	        "Enabled": {
		//	          "description": "",
		//	          "type": "boolean"
		//	        },
		//	        "LogLevel": {
		//	          "description": "",
		//	          "enum": [
		//	            "CRITICAL",
		//	            "ERROR",
		//	            "WARNING",
		//	            "INFO",
		//	            "DEBUG"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "SchedulerLogs": {
		//	      "additionalProperties": false,
		//	      "description": "Logging configuration for a specific airflow component.",
		//	      "properties": {
		//	        "CloudWatchLogGroupArn": {
		//	          "description": "",
		//	          "maxLength": 1224,
		//	          "pattern": "^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:logs:[a-z0-9\\-]+:\\d{12}:log-group:\\w+",
		//	          "type": "string"
		//	        },
		//	        "Enabled": {
		//	          "description": "",
		//	          "type": "boolean"
		//	        },
		//	        "LogLevel": {
		//	          "description": "",
		//	          "enum": [
		//	            "CRITICAL",
		//	            "ERROR",
		//	            "WARNING",
		//	            "INFO",
		//	            "DEBUG"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "TaskLogs": {
		//	      "additionalProperties": false,
		//	      "description": "Logging configuration for a specific airflow component.",
		//	      "properties": {
		//	        "CloudWatchLogGroupArn": {
		//	          "description": "",
		//	          "maxLength": 1224,
		//	          "pattern": "^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:logs:[a-z0-9\\-]+:\\d{12}:log-group:\\w+",
		//	          "type": "string"
		//	        },
		//	        "Enabled": {
		//	          "description": "",
		//	          "type": "boolean"
		//	        },
		//	        "LogLevel": {
		//	          "description": "",
		//	          "enum": [
		//	            "CRITICAL",
		//	            "ERROR",
		//	            "WARNING",
		//	            "INFO",
		//	            "DEBUG"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "WebserverLogs": {
		//	      "additionalProperties": false,
		//	      "description": "Logging configuration for a specific airflow component.",
		//	      "properties": {
		//	        "CloudWatchLogGroupArn": {
		//	          "description": "",
		//	          "maxLength": 1224,
		//	          "pattern": "^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:logs:[a-z0-9\\-]+:\\d{12}:log-group:\\w+",
		//	          "type": "string"
		//	        },
		//	        "Enabled": {
		//	          "description": "",
		//	          "type": "boolean"
		//	        },
		//	        "LogLevel": {
		//	          "description": "",
		//	          "enum": [
		//	            "CRITICAL",
		//	            "ERROR",
		//	            "WARNING",
		//	            "INFO",
		//	            "DEBUG"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "WorkerLogs": {
		//	      "additionalProperties": false,
		//	      "description": "Logging configuration for a specific airflow component.",
		//	      "properties": {
		//	        "CloudWatchLogGroupArn": {
		//	          "description": "",
		//	          "maxLength": 1224,
		//	          "pattern": "^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:logs:[a-z0-9\\-]+:\\d{12}:log-group:\\w+",
		//	          "type": "string"
		//	        },
		//	        "Enabled": {
		//	          "description": "",
		//	          "type": "boolean"
		//	        },
		//	        "LogLevel": {
		//	          "description": "",
		//	          "enum": [
		//	            "CRITICAL",
		//	            "ERROR",
		//	            "WARNING",
		//	            "INFO",
		//	            "DEBUG"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"logging_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DagProcessingLogs
				"dag_processing_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CloudWatchLogGroupArn
						"cloudwatch_log_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: LogLevel
						"log_level": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Logging configuration for a specific airflow component.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SchedulerLogs
				"scheduler_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CloudWatchLogGroupArn
						"cloudwatch_log_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: LogLevel
						"log_level": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Logging configuration for a specific airflow component.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: TaskLogs
				"task_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CloudWatchLogGroupArn
						"cloudwatch_log_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: LogLevel
						"log_level": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Logging configuration for a specific airflow component.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: WebserverLogs
				"webserver_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CloudWatchLogGroupArn
						"cloudwatch_log_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: LogLevel
						"log_level": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Logging configuration for a specific airflow component.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: WorkerLogs
				"worker_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CloudWatchLogGroupArn
						"cloudwatch_log_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: LogLevel
						"log_level": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Logging configuration for a specific airflow component.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Logging configuration for the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaxWorkers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Maximum worker compute units.",
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"max_workers": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Maximum worker compute units.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MinWorkers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Minimum worker compute units.",
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"min_workers": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Minimum worker compute units.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Customer-defined identifier for the environment, unique per customer region.",
		//	  "maxLength": 80,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z][0-9a-zA-Z\\-_]*$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Customer-defined identifier for the environment, unique per customer region.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configures the network resources of the environment.",
		//	  "properties": {
		//	    "SecurityGroupIds": {
		//	      "description": "A list of security groups to use for the environment.",
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "description": "",
		//	        "maxLength": 1024,
		//	        "minLength": 1,
		//	        "pattern": "^sg-[a-zA-Z0-9\\-._]+$",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 5,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    },
		//	    "SubnetIds": {
		//	      "description": "A list of subnets to use for the environment. These must be private subnets, in the same VPC, in two different availability zones.",
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "description": "",
		//	        "maxLength": 1024,
		//	        "pattern": "^subnet-[a-zA-Z0-9\\-._]+$",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 2,
		//	      "minItems": 2,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"network_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SecurityGroupIds
				"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of security groups to use for the environment.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SubnetIds
				"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of subnets to use for the environment. These must be private subnets, in the same VPC, in two different availability zones.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configures the network resources of the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PluginsS3ObjectVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents an version ID for an S3 object.",
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"plugins_s3_object_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents an version ID for an S3 object.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PluginsS3Path
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents an S3 prefix relative to the root of an S3 bucket.",
		//	  "maxLength": 1024,
		//	  "pattern": ".*",
		//	  "type": "string"
		//	}
		"plugins_s3_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents an S3 prefix relative to the root of an S3 bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RequirementsS3ObjectVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents an version ID for an S3 object.",
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"requirements_s3_object_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents an version ID for an S3 object.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RequirementsS3Path
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents an S3 prefix relative to the root of an S3 bucket.",
		//	  "maxLength": 1024,
		//	  "pattern": ".*",
		//	  "type": "string"
		//	}
		"requirements_s3_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents an S3 prefix relative to the root of an S3 bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Schedulers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Scheduler compute units.",
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"schedulers": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Scheduler compute units.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceBucketArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN for the AWS S3 bucket to use as the source of DAGs and plugins for the environment.",
		//	  "maxLength": 1224,
		//	  "minLength": 1,
		//	  "pattern": "^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:s3:::[a-z0-9.\\-]+$",
		//	  "type": "string"
		//	}
		"source_bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN for the AWS S3 bucket to use as the source of DAGs and plugins for the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A map of tags for the environment.",
		//	  "type": "object"
		//	}
		"tags": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A map of tags for the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WebserverAccessMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Choice for mode of webserver access including over public internet or via private VPC endpoint.",
		//	  "enum": [
		//	    "PRIVATE_ONLY",
		//	    "PUBLIC_ONLY"
		//	  ],
		//	  "type": "string"
		//	}
		"webserver_access_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Choice for mode of webserver access including over public internet or via private VPC endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WebserverUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Url endpoint for the environment's Airflow UI.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^https://.+$",
		//	  "type": "string"
		//	}
		"webserver_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Url endpoint for the environment's Airflow UI.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WeeklyMaintenanceWindowStart
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Start time for the weekly maintenance window.",
		//	  "maxLength": 9,
		//	  "pattern": "(MON|TUE|WED|THU|FRI|SAT|SUN):([01]\\d|2[0-3]):(00|30)",
		//	  "type": "string"
		//	}
		"weekly_maintenance_window_start": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Start time for the weekly maintenance window.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MWAA::Environment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MWAA::Environment").WithTerraformTypeName("awscc_mwaa_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"airflow_configuration_options":   "AirflowConfigurationOptions",
		"airflow_version":                 "AirflowVersion",
		"arn":                             "Arn",
		"cloudwatch_log_group_arn":        "CloudWatchLogGroupArn",
		"dag_processing_logs":             "DagProcessingLogs",
		"dag_s3_path":                     "DagS3Path",
		"enabled":                         "Enabled",
		"environment_class":               "EnvironmentClass",
		"execution_role_arn":              "ExecutionRoleArn",
		"kms_key":                         "KmsKey",
		"log_level":                       "LogLevel",
		"logging_configuration":           "LoggingConfiguration",
		"max_workers":                     "MaxWorkers",
		"min_workers":                     "MinWorkers",
		"name":                            "Name",
		"network_configuration":           "NetworkConfiguration",
		"plugins_s3_object_version":       "PluginsS3ObjectVersion",
		"plugins_s3_path":                 "PluginsS3Path",
		"requirements_s3_object_version":  "RequirementsS3ObjectVersion",
		"requirements_s3_path":            "RequirementsS3Path",
		"scheduler_logs":                  "SchedulerLogs",
		"schedulers":                      "Schedulers",
		"security_group_ids":              "SecurityGroupIds",
		"source_bucket_arn":               "SourceBucketArn",
		"subnet_ids":                      "SubnetIds",
		"tags":                            "Tags",
		"task_logs":                       "TaskLogs",
		"webserver_access_mode":           "WebserverAccessMode",
		"webserver_logs":                  "WebserverLogs",
		"webserver_url":                   "WebserverUrl",
		"weekly_maintenance_window_start": "WeeklyMaintenanceWindowStart",
		"worker_logs":                     "WorkerLogs",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
