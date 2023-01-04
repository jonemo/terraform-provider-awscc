// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_mediapackage_asset", assetDataSource)
}

// assetDataSource returns the Terraform awscc_mediapackage_asset data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaPackage::Asset resource.
func assetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Asset.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Asset.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time the Asset was initially submitted for Ingest.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time the Asset was initially submitted for Ingest.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EgressEndpoints
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of egress endpoints available for the Asset.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The endpoint URL used to access an Asset using one PackagingConfiguration.",
		//	    "properties": {
		//	      "PackagingConfigurationId": {
		//	        "description": "The ID of the PackagingConfiguration being applied to the Asset.",
		//	        "type": "string"
		//	      },
		//	      "Url": {
		//	        "description": "The URL of the parent manifest for the repackaged Asset.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "PackagingConfigurationId",
		//	      "Url"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"egress_endpoints": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: PackagingConfigurationId
					"packaging_configuration_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the PackagingConfiguration being applied to the Asset.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Url
					"url": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The URL of the parent manifest for the repackaged Asset.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of egress endpoints available for the Asset.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier for the Asset.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier for the Asset.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PackagingGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the PackagingGroup for the Asset.",
		//	  "type": "string"
		//	}
		"packaging_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the PackagingGroup for the Asset.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The resource ID to include in SPEKE key requests.",
		//	  "type": "string"
		//	}
		"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The resource ID to include in SPEKE key requests.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of the source object in S3.",
		//	  "type": "string"
		//	}
		"source_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of the source object in S3.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IAM role_arn used to access the source S3 bucket.",
		//	  "type": "string"
		//	}
		"source_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IAM role_arn used to access the source S3 bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of tags associated with a resource",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
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
			Description: "A collection of tags associated with a resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MediaPackage::Asset",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaPackage::Asset").WithTerraformTypeName("awscc_mediapackage_asset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                        "Arn",
		"created_at":                 "CreatedAt",
		"egress_endpoints":           "EgressEndpoints",
		"id":                         "Id",
		"key":                        "Key",
		"packaging_configuration_id": "PackagingConfigurationId",
		"packaging_group_id":         "PackagingGroupId",
		"resource_id":                "ResourceId",
		"source_arn":                 "SourceArn",
		"source_role_arn":            "SourceRoleArn",
		"tags":                       "Tags",
		"url":                        "Url",
		"value":                      "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
