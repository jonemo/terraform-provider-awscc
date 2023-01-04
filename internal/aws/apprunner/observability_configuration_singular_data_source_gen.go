// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apprunner

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apprunner_observability_configuration", observabilityConfigurationDataSource)
}

// observabilityConfigurationDataSource returns the Terraform awscc_apprunner_observability_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppRunner::ObservabilityConfiguration resource.
func observabilityConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Latest
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "It's set to true for the configuration with the highest Revision among all configurations that share the same Name. It's set to false otherwise.",
		//	  "type": "boolean"
		//	}
		"latest": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "It's set to true for the configuration with the highest Revision among all configurations that share the same Name. It's set to false otherwise.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ObservabilityConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of this ObservabilityConfiguration",
		//	  "maxLength": 1011,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"observability_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of this ObservabilityConfiguration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ObservabilityConfigurationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the observability configuration. When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.",
		//	  "maxLength": 32,
		//	  "minLength": 4,
		//	  "pattern": "[A-Za-z0-9][A-Za-z0-9\\-_]{3,31}",
		//	  "type": "string"
		//	}
		"observability_configuration_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the observability configuration. When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ObservabilityConfigurationRevision
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The revision of this observability configuration. It's unique among all the active configurations ('Status': 'ACTIVE') that share the same ObservabilityConfigurationName.",
		//	  "type": "integer"
		//	}
		"observability_configuration_revision": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The revision of this observability configuration. It's unique among all the active configurations ('Status': 'ACTIVE') that share the same ObservabilityConfigurationName.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of metadata items that you can associate with your observability configuration resource. A tag is a key-value pair.",
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
		//	    "type": "object"
		//	  },
		//	  "type": "array"
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
			Description: "A list of metadata items that you can associate with your observability configuration resource. A tag is a key-value pair.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TraceConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing.",
		//	  "properties": {
		//	    "Vendor": {
		//	      "description": "The implementation provider chosen for tracing App Runner services.",
		//	      "enum": [
		//	        "AWSXRAY"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Vendor"
		//	  ],
		//	  "type": "object"
		//	}
		"trace_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Vendor
				"vendor": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The implementation provider chosen for tracing App Runner services.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::AppRunner::ObservabilityConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppRunner::ObservabilityConfiguration").WithTerraformTypeName("awscc_apprunner_observability_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                                  "Key",
		"latest":                               "Latest",
		"observability_configuration_arn":      "ObservabilityConfigurationArn",
		"observability_configuration_name":     "ObservabilityConfigurationName",
		"observability_configuration_revision": "ObservabilityConfigurationRevision",
		"tags":                                 "Tags",
		"trace_configuration":                  "TraceConfiguration",
		"value":                                "Value",
		"vendor":                               "Vendor",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
