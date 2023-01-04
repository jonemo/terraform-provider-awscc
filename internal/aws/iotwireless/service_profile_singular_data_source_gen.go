// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotwireless_service_profile", serviceProfileDataSource)
}

// serviceProfileDataSource returns the Terraform awscc_iotwireless_service_profile data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTWireless::ServiceProfile resource.
func serviceProfileDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Service profile Arn. Returned after successful create.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Service profile Arn. Returned after successful create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Service profile Id. Returned after successful create.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Service profile Id. Returned after successful create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoRaWAN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "LoRaWAN supports all LoRa specific attributes for service profile for CreateServiceProfile operation",
		//	  "properties": {
		//	    "AddGwMetadata": {
		//	      "type": "boolean"
		//	    },
		//	    "ChannelMask": {
		//	      "type": "string"
		//	    },
		//	    "DevStatusReqFreq": {
		//	      "type": "integer"
		//	    },
		//	    "DlBucketSize": {
		//	      "type": "integer"
		//	    },
		//	    "DlRate": {
		//	      "type": "integer"
		//	    },
		//	    "DlRatePolicy": {
		//	      "type": "string"
		//	    },
		//	    "DrMax": {
		//	      "type": "integer"
		//	    },
		//	    "DrMin": {
		//	      "type": "integer"
		//	    },
		//	    "HrAllowed": {
		//	      "type": "boolean"
		//	    },
		//	    "MinGwDiversity": {
		//	      "type": "integer"
		//	    },
		//	    "NwkGeoLoc": {
		//	      "type": "boolean"
		//	    },
		//	    "PrAllowed": {
		//	      "type": "boolean"
		//	    },
		//	    "RaAllowed": {
		//	      "type": "boolean"
		//	    },
		//	    "ReportDevStatusBattery": {
		//	      "type": "boolean"
		//	    },
		//	    "ReportDevStatusMargin": {
		//	      "type": "boolean"
		//	    },
		//	    "TargetPer": {
		//	      "type": "integer"
		//	    },
		//	    "UlBucketSize": {
		//	      "type": "integer"
		//	    },
		//	    "UlRate": {
		//	      "type": "integer"
		//	    },
		//	    "UlRatePolicy": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"lo_ra_wan": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AddGwMetadata
				"add_gw_metadata": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ChannelMask
				"channel_mask": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DevStatusReqFreq
				"dev_status_req_freq": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DlBucketSize
				"dl_bucket_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DlRate
				"dl_rate": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DlRatePolicy
				"dl_rate_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DrMax
				"dr_max": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DrMin
				"dr_min": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: HrAllowed
				"hr_allowed": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MinGwDiversity
				"min_gw_diversity": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: NwkGeoLoc
				"nwk_geo_loc": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PrAllowed
				"pr_allowed": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RaAllowed
				"ra_allowed": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ReportDevStatusBattery
				"report_dev_status_battery": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ReportDevStatusMargin
				"report_dev_status_margin": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TargetPer
				"target_per": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: UlBucketSize
				"ul_bucket_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: UlRate
				"ul_rate": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: UlRatePolicy
				"ul_rate_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "LoRaWAN supports all LoRa specific attributes for service profile for CreateServiceProfile operation",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of service profile",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of service profile",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the service profile.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
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
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the service profile.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTWireless::ServiceProfile",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::ServiceProfile").WithTerraformTypeName("awscc_iotwireless_service_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"add_gw_metadata":           "AddGwMetadata",
		"arn":                       "Arn",
		"channel_mask":              "ChannelMask",
		"dev_status_req_freq":       "DevStatusReqFreq",
		"dl_bucket_size":            "DlBucketSize",
		"dl_rate":                   "DlRate",
		"dl_rate_policy":            "DlRatePolicy",
		"dr_max":                    "DrMax",
		"dr_min":                    "DrMin",
		"hr_allowed":                "HrAllowed",
		"id":                        "Id",
		"key":                       "Key",
		"lo_ra_wan":                 "LoRaWAN",
		"min_gw_diversity":          "MinGwDiversity",
		"name":                      "Name",
		"nwk_geo_loc":               "NwkGeoLoc",
		"pr_allowed":                "PrAllowed",
		"ra_allowed":                "RaAllowed",
		"report_dev_status_battery": "ReportDevStatusBattery",
		"report_dev_status_margin":  "ReportDevStatusMargin",
		"tags":                      "Tags",
		"target_per":                "TargetPer",
		"ul_bucket_size":            "UlBucketSize",
		"ul_rate":                   "UlRate",
		"ul_rate_policy":            "UlRatePolicy",
		"value":                     "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
