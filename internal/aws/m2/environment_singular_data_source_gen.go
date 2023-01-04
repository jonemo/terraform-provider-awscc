// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package m2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_m2_environment", environmentDataSource)
}

// environmentDataSource returns the Terraform awscc_m2_environment data source.
// This Terraform data source corresponds to the CloudFormation AWS::M2::Environment resource.
func environmentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the environment.",
		//	  "maxLength": 500,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EngineType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The target platform for the environment.",
		//	  "enum": [
		//	    "microfocus",
		//	    "bluage"
		//	  ],
		//	  "type": "string"
		//	}
		"engine_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The target platform for the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of the runtime engine for the environment.",
		//	  "pattern": "^\\S{1,10}$",
		//	  "type": "string"
		//	}
		"engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version of the runtime engine for the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the runtime environment.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"environment_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the runtime environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier of the environment.",
		//	  "pattern": "^\\S{1,80}$",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier of the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HighAvailabilityConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Defines the details of a high availability configuration.",
		//	  "properties": {
		//	    "DesiredCapacity": {
		//	      "maximum": 100,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "DesiredCapacity"
		//	  ],
		//	  "type": "object"
		//	}
		"high_availability_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DesiredCapacity
				"desired_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Defines the details of a high availability configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of instance underlying the environment.",
		//	  "pattern": "^\\S{1,20}$",
		//	  "type": "string"
		//	}
		"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of instance underlying the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the environment.",
		//	  "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PreferredMaintenanceWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.",
		//	  "pattern": "^\\S{1,50}$",
		//	  "type": "string"
		//	}
		"preferred_maintenance_window": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PubliclyAccessible
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the environment is publicly accessible.",
		//	  "type": "boolean"
		//	}
		"publicly_accessible": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the environment is publicly accessible.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of security groups for the VPC associated with this environment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^\\S{1,50}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The list of security groups for the VPC associated with this environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StorageConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The storage configurations defined for the runtime environment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "Defines the storage configuration for an environment.",
		//	    "properties": {
		//	      "Efs": {
		//	        "additionalProperties": false,
		//	        "description": "Defines the storage configuration for an Amazon EFS file system.",
		//	        "properties": {
		//	          "FileSystemId": {
		//	            "description": "The file system identifier.",
		//	            "pattern": "^\\S{1,200}$",
		//	            "type": "string"
		//	          },
		//	          "MountPoint": {
		//	            "description": "The mount point for the file system.",
		//	            "pattern": "^\\S{1,200}$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "FileSystemId",
		//	          "MountPoint"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Fsx": {
		//	        "additionalProperties": false,
		//	        "description": "Defines the storage configuration for an Amazon FSx file system.",
		//	        "properties": {
		//	          "FileSystemId": {
		//	            "description": "The file system identifier.",
		//	            "pattern": "^\\S{1,200}$",
		//	            "type": "string"
		//	          },
		//	          "MountPoint": {
		//	            "description": "The mount point for the file system.",
		//	            "pattern": "^\\S{1,200}$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "FileSystemId",
		//	          "MountPoint"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"storage_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Efs
					"efs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: FileSystemId
							"file_system_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The file system identifier.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: MountPoint
							"mount_point": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The mount point for the file system.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Defines the storage configuration for an Amazon EFS file system.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Fsx
					"fsx": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: FileSystemId
							"file_system_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The file system identifier.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: MountPoint
							"mount_point": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The mount point for the file system.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Defines the storage configuration for an Amazon FSx file system.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The storage configurations defined for the runtime environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifiers of the subnets assigned to this runtime environment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^\\S{1,50}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The unique identifiers of the subnets assigned to this runtime environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Tags associated to this environment.",
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "minLength": 0,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Tags associated to this environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::M2::Environment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::M2::Environment").WithTerraformTypeName("awscc_m2_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":                  "Description",
		"desired_capacity":             "DesiredCapacity",
		"efs":                          "Efs",
		"engine_type":                  "EngineType",
		"engine_version":               "EngineVersion",
		"environment_arn":              "EnvironmentArn",
		"environment_id":               "EnvironmentId",
		"file_system_id":               "FileSystemId",
		"fsx":                          "Fsx",
		"high_availability_config":     "HighAvailabilityConfig",
		"instance_type":                "InstanceType",
		"kms_key_id":                   "KmsKeyId",
		"mount_point":                  "MountPoint",
		"name":                         "Name",
		"preferred_maintenance_window": "PreferredMaintenanceWindow",
		"publicly_accessible":          "PubliclyAccessible",
		"security_group_ids":           "SecurityGroupIds",
		"storage_configurations":       "StorageConfigurations",
		"subnet_ids":                   "SubnetIds",
		"tags":                         "Tags",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
