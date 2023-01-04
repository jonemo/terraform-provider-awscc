// Code generated by generators/resource/main.go; DO NOT EDIT.

package ce

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ce_anomaly_monitor", anomalyMonitorResource)
}

// anomalyMonitorResource returns the Terraform awscc_ce_anomaly_monitor resource.
// This Terraform resource corresponds to the CloudFormation AWS::CE::AnomalyMonitor resource.
func anomalyMonitorResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date when the monitor was created. ",
		//	  "maxLength": 40,
		//	  "minLength": 0,
		//	  "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?",
		//	  "type": "string"
		//	}
		"creation_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date when the monitor was created. ",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DimensionalValueCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The value for evaluated dimensions.",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"dimensional_value_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The value for evaluated dimensions.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastEvaluatedDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date when the monitor last evaluated for anomalies.",
		//	  "maxLength": 40,
		//	  "minLength": 0,
		//	  "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?",
		//	  "type": "string"
		//	}
		"last_evaluated_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date when the monitor last evaluated for anomalies.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date when the monitor was last updated.",
		//	  "maxLength": 40,
		//	  "minLength": 0,
		//	  "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?",
		//	  "type": "string"
		//	}
		"last_updated_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date when the monitor was last updated.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MonitorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Monitor ARN",
		//	  "pattern": "^arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
		//	  "type": "string"
		//	}
		"monitor_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Monitor ARN",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MonitorDimension
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The dimensions to evaluate",
		//	  "enum": [
		//	    "SERVICE"
		//	  ],
		//	  "type": "string"
		//	}
		"monitor_dimension": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The dimensions to evaluate",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SERVICE",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MonitorName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the monitor.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "pattern": "[\\S\\s]*",
		//	  "type": "string"
		//	}
		"monitor_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the monitor.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 1024),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\S\\s]*"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: MonitorSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"monitor_specification": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MonitorType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "DIMENSIONAL",
		//	    "CUSTOM"
		//	  ],
		//	  "type": "string"
		//	}
		"monitor_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"DIMENSIONAL",
					"CUSTOM",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags to assign to monitor.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name for the tag.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag.",
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
		"resource_tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name for the tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags to assign to monitor.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. You can use Cost Anomaly Detection by creating monitor.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::AnomalyMonitor").WithTerraformTypeName("awscc_ce_anomaly_monitor")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_date":           "CreationDate",
		"dimensional_value_count": "DimensionalValueCount",
		"key":                     "Key",
		"last_evaluated_date":     "LastEvaluatedDate",
		"last_updated_date":       "LastUpdatedDate",
		"monitor_arn":             "MonitorArn",
		"monitor_dimension":       "MonitorDimension",
		"monitor_name":            "MonitorName",
		"monitor_specification":   "MonitorSpecification",
		"monitor_type":            "MonitorType",
		"resource_tags":           "ResourceTags",
		"value":                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
