// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_globalaccelerator_endpoint_group", endpointGroupResource)
}

// endpointGroupResource returns the Terraform awscc_globalaccelerator_endpoint_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::GlobalAccelerator::EndpointGroup resource.
func endpointGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: EndpointConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of endpoint objects.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The configuration for a given endpoint",
		//	    "properties": {
		//	      "AttachmentArn": {
		//	        "description": "Attachment ARN that provides access control to the cross account endpoint. Not required for resources hosted in the same account as the endpoint group.",
		//	        "type": "string"
		//	      },
		//	      "ClientIPPreservationEnabled": {
		//	        "default": true,
		//	        "description": "true if client ip should be preserved",
		//	        "type": "boolean"
		//	      },
		//	      "EndpointId": {
		//	        "description": "Id of the endpoint. For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID",
		//	        "type": "string"
		//	      },
		//	      "Weight": {
		//	        "default": 100,
		//	        "description": "The weight for the endpoint.",
		//	        "maximum": 255,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "required": [
		//	      "EndpointId"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"endpoint_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AttachmentArn
					"attachment_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Attachment ARN that provides access control to the cross account endpoint. Not required for resources hosted in the same account as the endpoint group.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ClientIPPreservationEnabled
					"client_ip_preservation_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "true if client ip should be preserved",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
							generic.BoolDefaultValue(true),
							boolplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: EndpointId
					"endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Id of the endpoint. For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Weight
					"weight": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "The weight for the endpoint.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{ /*START VALIDATORS*/
							int64validator.Between(0, 255),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							generic.Int64DefaultValue(100),
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of endpoint objects.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EndpointGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the endpoint group",
		//	  "type": "string"
		//	}
		"endpoint_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the endpoint group",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EndpointGroupRegion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the AWS Region where the endpoint group is located",
		//	  "type": "string"
		//	}
		"endpoint_group_region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the AWS Region where the endpoint group is located",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckIntervalSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 30,
		//	  "description": "The time in seconds between each health check for an endpoint. Must be a value of 10 or 30",
		//	  "type": "integer"
		//	}
		"health_check_interval_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The time in seconds between each health check for an endpoint. Must be a value of 10 or 30",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				generic.Int64DefaultValue(30),
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckPath
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "/",
		//	  "description": "",
		//	  "type": "string"
		//	}
		"health_check_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("/"),
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckPort
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": -1,
		//	  "description": "The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
		//	  "maximum": 65535,
		//	  "minimum": -1,
		//	  "type": "integer"
		//	}
		"health_check_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(-1, 65535),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				generic.Int64DefaultValue(-1),
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckProtocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "TCP",
		//	  "description": "The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
		//	  "enum": [
		//	    "TCP",
		//	    "HTTP",
		//	    "HTTPS"
		//	  ],
		//	  "type": "string"
		//	}
		"health_check_protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"TCP",
					"HTTP",
					"HTTPS",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("TCP"),
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ListenerArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the listener",
		//	  "type": "string"
		//	}
		"listener_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the listener",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PortOverrides
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "listener to endpoint port mapping.",
		//	    "properties": {
		//	      "EndpointPort": {
		//	        "description": "A network port number",
		//	        "maximum": 65535,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      },
		//	      "ListenerPort": {
		//	        "description": "A network port number",
		//	        "maximum": 65535,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "required": [
		//	      "ListenerPort",
		//	      "EndpointPort"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"port_overrides": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: EndpointPort
					"endpoint_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "A network port number",
						Required:    true,
						Validators: []validator.Int64{ /*START VALIDATORS*/
							int64validator.Between(0, 65535),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: ListenerPort
					"listener_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "A network port number",
						Required:    true,
						Validators: []validator.Int64{ /*START VALIDATORS*/
							int64validator.Between(0, 65535),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ThresholdCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 3,
		//	  "description": "The number of consecutive health checks required to set the state of the endpoint to unhealthy.",
		//	  "type": "integer"
		//	}
		"threshold_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of consecutive health checks required to set the state of the endpoint to unhealthy.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				generic.Int64DefaultValue(3),
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TrafficDialPercentage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 100,
		//	  "description": "The percentage of traffic to sent to an AWS Region",
		//	  "maximum": 100,
		//	  "minimum": 0,
		//	  "type": "number"
		//	}
		"traffic_dial_percentage": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The percentage of traffic to sent to an AWS Region",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Float64{ /*START VALIDATORS*/
				float64validator.Between(0.000000, 100.000000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				generic.Float64DefaultValue(100.000000),
				float64planmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::GlobalAccelerator::EndpointGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GlobalAccelerator::EndpointGroup").WithTerraformTypeName("awscc_globalaccelerator_endpoint_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attachment_arn":                 "AttachmentArn",
		"client_ip_preservation_enabled": "ClientIPPreservationEnabled",
		"endpoint_configurations":        "EndpointConfigurations",
		"endpoint_group_arn":             "EndpointGroupArn",
		"endpoint_group_region":          "EndpointGroupRegion",
		"endpoint_id":                    "EndpointId",
		"endpoint_port":                  "EndpointPort",
		"health_check_interval_seconds":  "HealthCheckIntervalSeconds",
		"health_check_path":              "HealthCheckPath",
		"health_check_port":              "HealthCheckPort",
		"health_check_protocol":          "HealthCheckProtocol",
		"listener_arn":                   "ListenerArn",
		"listener_port":                  "ListenerPort",
		"port_overrides":                 "PortOverrides",
		"threshold_count":                "ThresholdCount",
		"traffic_dial_percentage":        "TrafficDialPercentage",
		"weight":                         "Weight",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/EndpointConfigurations/*/AttachmentArn",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
