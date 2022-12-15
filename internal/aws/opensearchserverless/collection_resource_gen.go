// Code generated by generators/resource/main.go; DO NOT EDIT.

package opensearchserverless

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
	registry.AddResourceFactory("awscc_opensearchserverless_collection", collectionResource)
}

// collectionResource returns the Terraform awscc_opensearchserverless_collection resource.
// This Terraform resource corresponds to the CloudFormation AWS::OpenSearchServerless::Collection resource.
func collectionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the collection.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the collection.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"collection_endpoint": {
			// Property: CollectionEndpoint
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The endpoint for the collection.",
			//	  "type": "string"
			//	}
			Description: "The endpoint for the collection.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"dashboard_endpoint": {
			// Property: DashboardEndpoint
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The OpenSearch Dashboards endpoint for the collection.",
			//	  "type": "string"
			//	}
			Description: "The OpenSearch Dashboards endpoint for the collection.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The description of the collection",
			//	  "maxLength": 1000,
			//	  "type": "string"
			//	}
			Description: "The description of the collection",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(1000),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The identifier of the collection",
			//	  "maxLength": 40,
			//	  "minLength": 3,
			//	  "type": "string"
			//	}
			Description: "The identifier of the collection",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the collection.\n\nThe name must meet the following criteria:\nUnique to your account and AWS Region\nStarts with a lowercase letter\nContains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)\nContains between 3 and 32 characters\n",
			//	  "maxLength": 32,
			//	  "minLength": 3,
			//	  "pattern": "^[a-z][a-z0-9-]{2,31}$",
			//	  "type": "string"
			//	}
			Description: "The name of the collection.\n\nThe name must meet the following criteria:\nUnique to your account and AWS Region\nStarts with a lowercase letter\nContains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)\nContains between 3 and 32 characters\n",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(3, 32),
				validate.StringMatch(regexp.MustCompile("^[a-z][a-z0-9-]{2,31}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "List of tags to be added to the resource",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair metadata associated with resource",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key in the key-value pair",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value in the key-value pair",
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
			//	  "maxItems": 50,
			//	  "minItems": 0,
			//	  "type": "array"
			//	}
			Description: "List of tags to be added to the resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key in the key-value pair",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value in the key-value pair",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// Tags is a write-only property.
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The possible types for the collection",
			//	  "enum": [
			//	    "SEARCH",
			//	    "TIMESERIES"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The possible types for the collection",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"SEARCH",
					"TIMESERIES",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Amazon OpenSearchServerless collection resource",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::OpenSearchServerless::Collection").WithTerraformTypeName("awscc_opensearchserverless_collection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"collection_endpoint": "CollectionEndpoint",
		"dashboard_endpoint":  "DashboardEndpoint",
		"description":         "Description",
		"id":                  "Id",
		"key":                 "Key",
		"name":                "Name",
		"tags":                "Tags",
		"type":                "Type",
		"value":               "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
