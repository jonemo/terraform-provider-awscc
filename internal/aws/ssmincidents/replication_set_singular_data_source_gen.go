// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ssmincidents

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ssmincidents_replication_set", replicationSetDataSource)
}

// replicationSetDataSource returns the Terraform awscc_ssmincidents_replication_set data source.
// This Terraform data source corresponds to the CloudFormation AWS::SSMIncidents::ReplicationSet resource.
func replicationSetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the ReplicationSet.",
		//	  "maxLength": 1000,
		//	  "pattern": "^arn:aws(-(cn|us-gov|iso(-b)?))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the ReplicationSet.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeletionProtected
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "Configures the ReplicationSet deletion protection.",
		//	  "type": "boolean"
		//	}
		"deletion_protected": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Configures the ReplicationSet deletion protection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Regions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ReplicationSet configuration.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The ReplicationSet regional configuration.",
		//	    "properties": {
		//	      "RegionConfiguration": {
		//	        "additionalProperties": false,
		//	        "description": "The ReplicationSet regional configuration.",
		//	        "properties": {
		//	          "SseKmsKeyId": {
		//	            "description": "The ARN of the ReplicationSet.",
		//	            "maxLength": 1000,
		//	            "pattern": "^arn:aws(-(cn|us-gov|iso(-b)?))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "SseKmsKeyId"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "RegionName": {
		//	        "description": "The AWS region name.",
		//	        "maxLength": 20,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 3,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"regions": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: RegionConfiguration
					"region_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: SseKmsKeyId
							"sse_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The ARN of the ReplicationSet.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The ReplicationSet regional configuration.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: RegionName
					"region_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The AWS region name.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The ReplicationSet configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": [],
		//	  "description": "The tags to apply to the replication set.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to tag a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags to apply to the replication set.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SSMIncidents::ReplicationSet",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMIncidents::ReplicationSet").WithTerraformTypeName("awscc_ssmincidents_replication_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                  "Arn",
		"deletion_protected":   "DeletionProtected",
		"key":                  "Key",
		"region_configuration": "RegionConfiguration",
		"region_name":          "RegionName",
		"regions":              "Regions",
		"sse_kms_key_id":       "SseKmsKeyId",
		"tags":                 "Tags",
		"value":                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
