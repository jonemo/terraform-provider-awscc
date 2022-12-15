// Code generated by generators/resource/main.go; DO NOT EDIT.

package groundstation

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_groundstation_dataflow_endpoint_group", dataflowEndpointGroupResource)
}

// dataflowEndpointGroupResource returns the Terraform awscc_groundstation_dataflow_endpoint_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::GroundStation::DataflowEndpointGroup resource.
func dataflowEndpointGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"endpoint_details": {
			// Property: EndpointDetails
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Endpoint": {
			//	        "additionalProperties": false,
			//	        "properties": {
			//	          "Address": {
			//	            "additionalProperties": false,
			//	            "properties": {
			//	              "Name": {
			//	                "type": "string"
			//	              },
			//	              "Port": {
			//	                "type": "integer"
			//	              }
			//	            },
			//	            "type": "object"
			//	          },
			//	          "Mtu": {
			//	            "type": "integer"
			//	          },
			//	          "Name": {
			//	            "pattern": "^[ a-zA-Z0-9_:-]{1,256}$",
			//	            "type": "string"
			//	          }
			//	        },
			//	        "type": "object"
			//	      },
			//	      "SecurityDetails": {
			//	        "additionalProperties": false,
			//	        "properties": {
			//	          "RoleArn": {
			//	            "type": "string"
			//	          },
			//	          "SecurityGroupIds": {
			//	            "items": {
			//	              "type": "string"
			//	            },
			//	            "type": "array"
			//	          },
			//	          "SubnetIds": {
			//	            "items": {
			//	              "type": "string"
			//	            },
			//	            "type": "array"
			//	          }
			//	        },
			//	        "type": "object"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "minItems": 1,
			//	  "type": "array"
			//	}
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"endpoint": {
						// Property: Endpoint
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"address": {
									// Property: Address
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"name": {
												// Property: Name
												Type:     types.StringType,
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"port": {
												// Property: Port
												Type:     types.Int64Type,
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"mtu": {
									// Property: Mtu
									Type:     types.Int64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("^[ a-zA-Z0-9_:-]{1,256}$"), ""),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"security_details": {
						// Property: SecurityDetails
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"role_arn": {
									// Property: RoleArn
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"security_group_ids": {
									// Property: SecurityGroupIds
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"subnet_ids": {
									// Property: SubnetIds
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(1),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,128}$",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,256}$",
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
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^[ a-zA-Z0-9\\+\\-=._:/@]{1,128}$"), ""),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^[ a-zA-Z0-9\\+\\-=._:/@]{1,256}$"), ""),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "AWS Ground Station DataflowEndpointGroup schema for CloudFormation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GroundStation::DataflowEndpointGroup").WithTerraformTypeName("awscc_groundstation_dataflow_endpoint_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":            "Address",
		"arn":                "Arn",
		"endpoint":           "Endpoint",
		"endpoint_details":   "EndpointDetails",
		"id":                 "Id",
		"key":                "Key",
		"mtu":                "Mtu",
		"name":               "Name",
		"port":               "Port",
		"role_arn":           "RoleArn",
		"security_details":   "SecurityDetails",
		"security_group_ids": "SecurityGroupIds",
		"subnet_ids":         "SubnetIds",
		"tags":               "Tags",
		"value":              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
