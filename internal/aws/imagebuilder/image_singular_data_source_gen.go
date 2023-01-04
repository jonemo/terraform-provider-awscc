// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_imagebuilder_image", imageDataSource)
}

// imageDataSource returns the Terraform awscc_imagebuilder_image data source.
// This Terraform data source corresponds to the CloudFormation AWS::ImageBuilder::Image resource.
func imageDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the image.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the image.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ContainerRecipeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
		//	  "type": "string"
		//	}
		"container_recipe_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DistributionConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the distribution configuration.",
		//	  "type": "string"
		//	}
		"distribution_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the distribution configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnhancedImageMetadataEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Collects additional information about the image being created, including the operating system (OS) version and package list.",
		//	  "type": "boolean"
		//	}
		"enhanced_image_metadata_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Collects additional information about the image being created, including the operating system (OS) version and package list.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AMI ID of the EC2 AMI in current region.",
		//	  "type": "string"
		//	}
		"image_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AMI ID of the EC2 AMI in current region.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageRecipeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
		//	  "type": "string"
		//	}
		"image_recipe_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageTestsConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The image tests configuration used when creating this image.",
		//	  "properties": {
		//	    "ImageTestsEnabled": {
		//	      "description": "ImageTestsEnabled",
		//	      "type": "boolean"
		//	    },
		//	    "TimeoutMinutes": {
		//	      "description": "TimeoutMinutes",
		//	      "maximum": 1440,
		//	      "minimum": 60,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"image_tests_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ImageTestsEnabled
				"image_tests_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "ImageTestsEnabled",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: TimeoutMinutes
				"timeout_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "TimeoutMinutes",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The image tests configuration used when creating this image.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "URI for containers created in current Region with default ECR image tag",
		//	  "type": "string"
		//	}
		"image_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "URI for containers created in current Region with default ECR image tag",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InfrastructureConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the infrastructure configuration.",
		//	  "type": "string"
		//	}
		"infrastructure_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the infrastructure configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the image.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the image.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The tags associated with the image.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The tags associated with the image.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ImageBuilder::Image",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::Image").WithTerraformTypeName("awscc_imagebuilder_image")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                              "Arn",
		"container_recipe_arn":             "ContainerRecipeArn",
		"distribution_configuration_arn":   "DistributionConfigurationArn",
		"enhanced_image_metadata_enabled":  "EnhancedImageMetadataEnabled",
		"image_id":                         "ImageId",
		"image_recipe_arn":                 "ImageRecipeArn",
		"image_tests_configuration":        "ImageTestsConfiguration",
		"image_tests_enabled":              "ImageTestsEnabled",
		"image_uri":                        "ImageUri",
		"infrastructure_configuration_arn": "InfrastructureConfigurationArn",
		"name":                             "Name",
		"tags":                             "Tags",
		"timeout_minutes":                  "TimeoutMinutes",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
