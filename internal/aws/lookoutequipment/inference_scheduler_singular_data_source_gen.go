// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lookoutequipment

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lookoutequipment_inference_scheduler", inferenceSchedulerDataSource)
}

// inferenceSchedulerDataSource returns the Terraform awscc_lookoutequipment_inference_scheduler data source.
// This Terraform data source corresponds to the CloudFormation AWS::LookoutEquipment::InferenceScheduler resource.
func inferenceSchedulerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DataDelayOffsetInMinutes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A period of time (in minutes) by which inference on the data is delayed after the data starts.",
		//	  "maximum": 60,
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"data_delay_offset_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "A period of time (in minutes) by which inference on the data is delayed after the data starts.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataInputConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.",
		//	  "properties": {
		//	    "InferenceInputNameConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies configuration information for the input data for the inference, including timestamp format and delimiter.",
		//	      "properties": {
		//	        "ComponentTimestampDelimiter": {
		//	          "description": "Indicates the delimiter character used between items in the data.",
		//	          "maxLength": 1,
		//	          "minLength": 0,
		//	          "pattern": "^(\\-|\\_|\\s)?$",
		//	          "type": "string"
		//	        },
		//	        "TimestampFormat": {
		//	          "description": "The format of the timestamp, whether Epoch time, or standard, with or without hyphens (-).",
		//	          "pattern": "^EPOCH|yyyy-MM-dd-HH-mm-ss|yyyyMMddHHmmss$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "InputTimeZoneOffset": {
		//	      "description": "Indicates the difference between your time zone and Greenwich Mean Time (GMT).",
		//	      "pattern": "^(\\+|\\-)[0-9]{2}\\:[0-9]{2}$",
		//	      "type": "string"
		//	    },
		//	    "S3InputConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies configuration information for the input data for the inference, including input data S3 location.",
		//	      "properties": {
		//	        "Bucket": {
		//	          "maxLength": 63,
		//	          "minLength": 3,
		//	          "pattern": "^[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$",
		//	          "type": "string"
		//	        },
		//	        "Prefix": {
		//	          "maxLength": 1024,
		//	          "minLength": 0,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Bucket"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3InputConfiguration"
		//	  ],
		//	  "type": "object"
		//	}
		"data_input_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: InferenceInputNameConfiguration
				"inference_input_name_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ComponentTimestampDelimiter
						"component_timestamp_delimiter": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Indicates the delimiter character used between items in the data.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: TimestampFormat
						"timestamp_format": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The format of the timestamp, whether Epoch time, or standard, with or without hyphens (-).",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies configuration information for the input data for the inference, including timestamp format and delimiter.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: InputTimeZoneOffset
				"input_time_zone_offset": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates the difference between your time zone and Greenwich Mean Time (GMT).",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: S3InputConfiguration
				"s3_input_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Bucket
						"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Prefix
						"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies configuration information for the input data for the inference, including input data S3 location.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataOutputConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.",
		//	  "properties": {
		//	    "KmsKeyId": {
		//	      "description": "The ID number for the AWS KMS key used to encrypt the inference output.",
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "pattern": "",
		//	      "type": "string"
		//	    },
		//	    "S3OutputConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies configuration information for the output results from the inference, including output S3 location.",
		//	      "properties": {
		//	        "Bucket": {
		//	          "maxLength": 63,
		//	          "minLength": 3,
		//	          "pattern": "^[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$",
		//	          "type": "string"
		//	        },
		//	        "Prefix": {
		//	          "maxLength": 1024,
		//	          "minLength": 0,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Bucket"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3OutputConfiguration"
		//	  ],
		//	  "type": "object"
		//	}
		"data_output_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KmsKeyId
				"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID number for the AWS KMS key used to encrypt the inference output.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: S3OutputConfiguration
				"s3_output_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Bucket
						"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Prefix
						"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies configuration information for the output results from the inference, including output S3 location.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataUploadFrequency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "How often data is uploaded to the source S3 bucket for the input data.",
		//	  "enum": [
		//	    "PT5M",
		//	    "PT10M",
		//	    "PT15M",
		//	    "PT30M",
		//	    "PT1H"
		//	  ],
		//	  "type": "string"
		//	}
		"data_upload_frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "How often data is uploaded to the source S3 bucket for the input data.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InferenceSchedulerArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the inference scheduler being created.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "arn:aws(-[^:]+)?:lookoutequipment:[a-zA-Z0-9\\-]*:[0-9]{12}:inference-scheduler\\/.+",
		//	  "type": "string"
		//	}
		"inference_scheduler_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the inference scheduler being created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InferenceSchedulerName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the inference scheduler being created.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9a-zA-Z_-]{1,200}$",
		//	  "type": "string"
		//	}
		"inference_scheduler_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the inference scheduler being created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModelName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the previously trained ML model being used to create the inference scheduler.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9a-zA-Z_-]{1,200}$",
		//	  "type": "string"
		//	}
		"model_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the previously trained ML model being used to create the inference scheduler.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "arn:aws(-[^:]+)?:iam::[0-9]{12}:role/.+",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServerSideKmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt inference scheduler data by Amazon Lookout for Equipment.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"server_side_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt inference scheduler data by Amazon Lookout for Equipment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Any tags associated with the inference scheduler.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A tag is a key-value pair that can be added to a resource as metadata.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key for the specified tag.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the specified tag.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "[\\s\\w+-=\\.:/@]*",
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
						Description: "The key for the specified tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the specified tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Any tags associated with the inference scheduler.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::LookoutEquipment::InferenceScheduler",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LookoutEquipment::InferenceScheduler").WithTerraformTypeName("awscc_lookoutequipment_inference_scheduler")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket":                             "Bucket",
		"component_timestamp_delimiter":      "ComponentTimestampDelimiter",
		"data_delay_offset_in_minutes":       "DataDelayOffsetInMinutes",
		"data_input_configuration":           "DataInputConfiguration",
		"data_output_configuration":          "DataOutputConfiguration",
		"data_upload_frequency":              "DataUploadFrequency",
		"inference_input_name_configuration": "InferenceInputNameConfiguration",
		"inference_scheduler_arn":            "InferenceSchedulerArn",
		"inference_scheduler_name":           "InferenceSchedulerName",
		"input_time_zone_offset":             "InputTimeZoneOffset",
		"key":                                "Key",
		"kms_key_id":                         "KmsKeyId",
		"model_name":                         "ModelName",
		"prefix":                             "Prefix",
		"role_arn":                           "RoleArn",
		"s3_input_configuration":             "S3InputConfiguration",
		"s3_output_configuration":            "S3OutputConfiguration",
		"server_side_kms_key_id":             "ServerSideKmsKeyId",
		"tags":                               "Tags",
		"timestamp_format":                   "TimestampFormat",
		"value":                              "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
