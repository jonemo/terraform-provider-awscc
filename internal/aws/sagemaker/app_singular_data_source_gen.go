// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sagemaker_app", appDataSource)
}

// appDataSource returns the Terraform awscc_sagemaker_app data source.
// This Terraform data source corresponds to the CloudFormation AWS::SageMaker::App resource.
func appDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AppArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the app.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "arn:aws[a-z\\-]*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:app/.*",
		//	  "type": "string"
		//	}
		"app_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the app.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AppName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the app.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}",
		//	  "type": "string"
		//	}
		"app_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the app.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AppType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of app.",
		//	  "enum": [
		//	    "JupyterServer",
		//	    "KernelGateway",
		//	    "RStudioServerPro",
		//	    "RSessionGateway",
		//	    "Canvas"
		//	  ],
		//	  "type": "string"
		//	}
		"app_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of app.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The domain ID.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"domain_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The domain ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceSpec
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.",
		//	  "properties": {
		//	    "InstanceType": {
		//	      "description": "The instance type that the image version runs on.",
		//	      "enum": [
		//	        "system",
		//	        "ml.t3.micro",
		//	        "ml.t3.small",
		//	        "ml.t3.medium",
		//	        "ml.t3.large",
		//	        "ml.t3.xlarge",
		//	        "ml.t3.2xlarge",
		//	        "ml.m5.large",
		//	        "ml.m5.xlarge",
		//	        "ml.m5.2xlarge",
		//	        "ml.m5.4xlarge",
		//	        "ml.m5.8xlarge",
		//	        "ml.m5.12xlarge",
		//	        "ml.m5.16xlarge",
		//	        "ml.m5.24xlarge",
		//	        "ml.c5.large",
		//	        "ml.c5.xlarge",
		//	        "ml.c5.2xlarge",
		//	        "ml.c5.4xlarge",
		//	        "ml.c5.9xlarge",
		//	        "ml.c5.12xlarge",
		//	        "ml.c5.18xlarge",
		//	        "ml.c5.24xlarge",
		//	        "ml.p3.2xlarge",
		//	        "ml.p3.8xlarge",
		//	        "ml.p3.16xlarge",
		//	        "ml.g4dn.xlarge",
		//	        "ml.g4dn.2xlarge",
		//	        "ml.g4dn.4xlarge",
		//	        "ml.g4dn.8xlarge",
		//	        "ml.g4dn.12xlarge",
		//	        "ml.g4dn.16xlarge",
		//	        "ml.r5.large",
		//	        "ml.r5.xlarge",
		//	        "ml.r5.2xlarge",
		//	        "ml.r5.4xlarge",
		//	        "ml.r5.8xlarge",
		//	        "ml.r5.12xlarge",
		//	        "ml.r5.16xlarge",
		//	        "ml.r5.24xlarge",
		//	        "ml.p3dn.24xlarge",
		//	        "ml.m5d.large",
		//	        "ml.m5d.xlarge",
		//	        "ml.m5d.2xlarge",
		//	        "ml.m5d.4xlarge",
		//	        "ml.m5d.8xlarge",
		//	        "ml.m5d.12xlarge",
		//	        "ml.m5d.16xlarge",
		//	        "ml.m5d.24xlarge",
		//	        "ml.g5.xlarge",
		//	        "ml.g5.2xlarge",
		//	        "ml.g5.4xlarge",
		//	        "ml.g5.8xlarge",
		//	        "ml.g5.12xlarge",
		//	        "ml.g5.16xlarge",
		//	        "ml.g5.24xlarge",
		//	        "ml.g5.48xlarge"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "SageMakerImageArn": {
		//	      "description": "The ARN of the SageMaker image that the image version belongs to.",
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "pattern": "^arn:aws(-[\\w]+)*:sagemaker:.+:[0-9]{12}:image/[a-z0-9]([-.]?[a-z0-9])*$",
		//	      "type": "string"
		//	    },
		//	    "SageMakerImageVersionArn": {
		//	      "description": "The ARN of the image version created on the instance.",
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "pattern": "^arn:aws(-[\\w]+)*:sagemaker:.+:[0-9]{12}:image-version/[a-z0-9]([-.]?[a-z0-9])*/[0-9]+$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"resource_spec": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: InstanceType
				"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The instance type that the image version runs on.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SageMakerImageArn
				"sage_maker_image_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the SageMaker image that the image version belongs to.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SageMakerImageVersionArn
				"sage_maker_image_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the image version created on the instance.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags to apply to the app.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of tags to apply to the app.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UserProfileName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user profile name.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}",
		//	  "type": "string"
		//	}
		"user_profile_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user profile name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SageMaker::App",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::App").WithTerraformTypeName("awscc_sagemaker_app")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_arn":                      "AppArn",
		"app_name":                     "AppName",
		"app_type":                     "AppType",
		"domain_id":                    "DomainId",
		"instance_type":                "InstanceType",
		"key":                          "Key",
		"resource_spec":                "ResourceSpec",
		"sage_maker_image_arn":         "SageMakerImageArn",
		"sage_maker_image_version_arn": "SageMakerImageVersionArn",
		"tags":                         "Tags",
		"user_profile_name":            "UserProfileName",
		"value":                        "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
