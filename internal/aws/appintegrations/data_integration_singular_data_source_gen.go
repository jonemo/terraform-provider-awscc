// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appintegrations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_appintegrations_data_integration", dataIntegrationDataSource)
}

// dataIntegrationDataSource returns the Terraform awscc_appintegrations_data_integration data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppIntegrations::DataIntegration resource.
func dataIntegrationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DataIntegrationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the data integration.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"data_integration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the data integration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The data integration description.",
		//	  "maxLength": 1000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The data integration description.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifer of the data integration.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifer of the data integration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The KMS key of the data integration.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": ".*\\S.*",
		//	  "type": "string"
		//	}
		"kms_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The KMS key of the data integration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the data integration.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9/\\._\\-]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the data integration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScheduleConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The name of the data and how often it should be pulled from the source.",
		//	  "properties": {
		//	    "FirstExecutionFrom": {
		//	      "description": "The start date for objects to import in the first flow run. Epoch or ISO timestamp format is supported.",
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "pattern": ".*\\S.*",
		//	      "type": "string"
		//	    },
		//	    "Object": {
		//	      "description": "The name of the object to pull from the data source.",
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "pattern": "^[a-zA-Z0-9/\\._\\-]+$",
		//	      "type": "string"
		//	    },
		//	    "ScheduleExpression": {
		//	      "description": "How often the data should be pulled from data source.",
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "pattern": ".*\\S.*",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "FirstExecutionFrom",
		//	    "Object",
		//	    "ScheduleExpression"
		//	  ],
		//	  "type": "object"
		//	}
		"schedule_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FirstExecutionFrom
				"first_execution_from": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The start date for objects to import in the first flow run. Epoch or ISO timestamp format is supported.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Object
				"object": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the object to pull from the data source.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ScheduleExpression
				"schedule_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "How often the data should be pulled from data source.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The name of the data and how often it should be pulled from the source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceURI
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URI of the data source.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^\\w+\\:\\/\\/\\w+\\/[\\w/!@#+=.-]+$",
		//	  "type": "string"
		//	}
		"source_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URI of the data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags (keys and values) associated with the data integration.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A label for tagging DataIntegration resources",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A key to identify the tag.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "Corresponding tag value for the key.",
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
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A key to identify the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Corresponding tag value for the key.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags (keys and values) associated with the data integration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::AppIntegrations::DataIntegration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppIntegrations::DataIntegration").WithTerraformTypeName("awscc_appintegrations_data_integration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"data_integration_arn": "DataIntegrationArn",
		"description":          "Description",
		"first_execution_from": "FirstExecutionFrom",
		"id":                   "Id",
		"key":                  "Key",
		"kms_key":              "KmsKey",
		"name":                 "Name",
		"object":               "Object",
		"schedule_config":      "ScheduleConfig",
		"schedule_expression":  "ScheduleExpression",
		"source_uri":           "SourceURI",
		"tags":                 "Tags",
		"value":                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
