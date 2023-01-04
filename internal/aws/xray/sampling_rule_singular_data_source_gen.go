// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package xray

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_xray_sampling_rule", samplingRuleDataSource)
}

// samplingRuleDataSource returns the Terraform awscc_xray_sampling_rule data source.
// This Terraform data source corresponds to the CloudFormation AWS::XRay::SamplingRule resource.
func samplingRuleDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: RuleARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	  "type": "string"
		//	}
		"rule_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RuleName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	  "maxLength": 32,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SamplingRule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Attributes": {
		//	      "additionalProperties": false,
		//	      "$comment": "String to string map",
		//	      "description": "Matches attributes derived from the request.",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "FixedRate": {
		//	      "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
		//	      "maximum": 1,
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "HTTPMethod": {
		//	      "description": "Matches the HTTP method from a request URL.",
		//	      "maxLength": 10,
		//	      "type": "string"
		//	    },
		//	    "Host": {
		//	      "description": "Matches the hostname from a request URL.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "Priority": {
		//	      "description": "The priority of the sampling rule.",
		//	      "maximum": 9999,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "ReservoirSize": {
		//	      "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "ResourceARN": {
		//	      "description": "Matches the ARN of the AWS resource on which the service runs.",
		//	      "maxLength": 500,
		//	      "type": "string"
		//	    },
		//	    "RuleARN": {
		//	      "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	      "type": "string"
		//	    },
		//	    "RuleName": {
		//	      "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	      "maxLength": 32,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "ServiceName": {
		//	      "description": "Matches the name that the service uses to identify itself in segments.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "ServiceType": {
		//	      "description": "Matches the origin that the service uses to identify its type in segments.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "URLPath": {
		//	      "description": "Matches the path from a request URL.",
		//	      "maxLength": 128,
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "description": "The version of the sampling rule format (1)",
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sampling_rule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Attributes
				"attributes":        // Pattern: ""
				schema.MapAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Matches attributes derived from the request.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: FixedRate
				"fixed_rate": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The percentage of matching requests to instrument, after the reservoir is exhausted.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: HTTPMethod
				"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the HTTP method from a request URL.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Host
				"host": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the hostname from a request URL.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Priority
				"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The priority of the sampling rule.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ReservoirSize
				"reservoir_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ResourceARN
				"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the ARN of the AWS resource on which the service runs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RuleARN
				"rule_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RuleName
				"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ServiceName
				"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the name that the service uses to identify itself in segments.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ServiceType
				"service_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the origin that the service uses to identify its type in segments.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: URLPath
				"url_path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the path from a request URL.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The version of the sampling rule format (1)",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SamplingRuleRecord
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CreatedAt": {
		//	      "description": "When the rule was created, in Unix time seconds.",
		//	      "type": "string"
		//	    },
		//	    "ModifiedAt": {
		//	      "description": "When the rule was modified, in Unix time seconds.",
		//	      "type": "string"
		//	    },
		//	    "SamplingRule": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Attributes": {
		//	          "additionalProperties": false,
		//	          "$comment": "String to string map",
		//	          "description": "Matches attributes derived from the request.",
		//	          "patternProperties": {
		//	            "": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "FixedRate": {
		//	          "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
		//	          "maximum": 1,
		//	          "minimum": 0,
		//	          "type": "number"
		//	        },
		//	        "HTTPMethod": {
		//	          "description": "Matches the HTTP method from a request URL.",
		//	          "maxLength": 10,
		//	          "type": "string"
		//	        },
		//	        "Host": {
		//	          "description": "Matches the hostname from a request URL.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        },
		//	        "Priority": {
		//	          "description": "The priority of the sampling rule.",
		//	          "maximum": 9999,
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        },
		//	        "ReservoirSize": {
		//	          "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
		//	          "minimum": 0,
		//	          "type": "integer"
		//	        },
		//	        "ResourceARN": {
		//	          "description": "Matches the ARN of the AWS resource on which the service runs.",
		//	          "maxLength": 500,
		//	          "type": "string"
		//	        },
		//	        "RuleARN": {
		//	          "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	          "type": "string"
		//	        },
		//	        "RuleName": {
		//	          "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	          "maxLength": 32,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "ServiceName": {
		//	          "description": "Matches the name that the service uses to identify itself in segments.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        },
		//	        "ServiceType": {
		//	          "description": "Matches the origin that the service uses to identify its type in segments.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        },
		//	        "URLPath": {
		//	          "description": "Matches the path from a request URL.",
		//	          "maxLength": 128,
		//	          "type": "string"
		//	        },
		//	        "Version": {
		//	          "description": "The version of the sampling rule format (1)",
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sampling_rule_record": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CreatedAt
				"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "When the rule was created, in Unix time seconds.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ModifiedAt
				"modified_at": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "When the rule was modified, in Unix time seconds.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SamplingRule
				"sampling_rule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Attributes
						"attributes":        // Pattern: ""
						schema.MapAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Matches attributes derived from the request.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: FixedRate
						"fixed_rate": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Description: "The percentage of matching requests to instrument, after the reservoir is exhausted.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: HTTPMethod
						"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the HTTP method from a request URL.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Host
						"host": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the hostname from a request URL.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Priority
						"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The priority of the sampling rule.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ReservoirSize
						"reservoir_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ResourceARN
						"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the ARN of the AWS resource on which the service runs.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: RuleARN
						"rule_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: RuleName
						"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ServiceName
						"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the name that the service uses to identify itself in segments.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ServiceType
						"service_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the origin that the service uses to identify its type in segments.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: URLPath
						"url_path": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the path from a request URL.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Version
						"version": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The version of the sampling rule format (1)",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SamplingRuleUpdate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Attributes": {
		//	      "additionalProperties": false,
		//	      "$comment": "String to string map",
		//	      "description": "Matches attributes derived from the request.",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "FixedRate": {
		//	      "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
		//	      "maximum": 1,
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "HTTPMethod": {
		//	      "description": "Matches the HTTP method from a request URL.",
		//	      "maxLength": 10,
		//	      "type": "string"
		//	    },
		//	    "Host": {
		//	      "description": "Matches the hostname from a request URL.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "Priority": {
		//	      "description": "The priority of the sampling rule.",
		//	      "maximum": 9999,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "ReservoirSize": {
		//	      "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "ResourceARN": {
		//	      "description": "Matches the ARN of the AWS resource on which the service runs.",
		//	      "maxLength": 500,
		//	      "type": "string"
		//	    },
		//	    "RuleARN": {
		//	      "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	      "type": "string"
		//	    },
		//	    "RuleName": {
		//	      "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	      "maxLength": 32,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "ServiceName": {
		//	      "description": "Matches the name that the service uses to identify itself in segments.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "ServiceType": {
		//	      "description": "Matches the origin that the service uses to identify its type in segments.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "URLPath": {
		//	      "description": "Matches the path from a request URL.",
		//	      "maxLength": 128,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sampling_rule_update": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Attributes
				"attributes":        // Pattern: ""
				schema.MapAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Matches attributes derived from the request.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: FixedRate
				"fixed_rate": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The percentage of matching requests to instrument, after the reservoir is exhausted.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: HTTPMethod
				"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the HTTP method from a request URL.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Host
				"host": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the hostname from a request URL.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Priority
				"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The priority of the sampling rule.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ReservoirSize
				"reservoir_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ResourceARN
				"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the ARN of the AWS resource on which the service runs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RuleARN
				"rule_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RuleName
				"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ServiceName
				"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the name that the service uses to identify itself in segments.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ServiceType
				"service_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the origin that the service uses to identify its type in segments.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: URLPath
				"url_path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the path from a request URL.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
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
		//	      "Key",
		//	      "Value"
		//	    ],
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
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::XRay::SamplingRule",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::XRay::SamplingRule").WithTerraformTypeName("awscc_xray_sampling_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attributes":           "Attributes",
		"created_at":           "CreatedAt",
		"fixed_rate":           "FixedRate",
		"host":                 "Host",
		"http_method":          "HTTPMethod",
		"key":                  "Key",
		"modified_at":          "ModifiedAt",
		"priority":             "Priority",
		"reservoir_size":       "ReservoirSize",
		"resource_arn":         "ResourceARN",
		"rule_arn":             "RuleARN",
		"rule_name":            "RuleName",
		"sampling_rule":        "SamplingRule",
		"sampling_rule_record": "SamplingRuleRecord",
		"sampling_rule_update": "SamplingRuleUpdate",
		"service_name":         "ServiceName",
		"service_type":         "ServiceType",
		"tags":                 "Tags",
		"url_path":             "URLPath",
		"value":                "Value",
		"version":              "Version",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
