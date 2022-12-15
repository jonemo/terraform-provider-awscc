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
	registry.AddResourceFactory("awscc_opensearchserverless_access_policy", accessPolicyResource)
}

// accessPolicyResource returns the Terraform awscc_opensearchserverless_access_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::OpenSearchServerless::AccessPolicy resource.
func accessPolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The description of the policy",
			//	  "maxLength": 1000,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The description of the policy",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1000),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the policy",
			//	  "maxLength": 32,
			//	  "minLength": 3,
			//	  "pattern": "^[a-z][a-z0-9-]{2,31}$",
			//	  "type": "string"
			//	}
			Description: "The name of the policy",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(3, 32),
				validate.StringMatch(regexp.MustCompile("^[a-z][a-z0-9-]{2,31}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The JSON policy document that is the content for the policy",
			//	  "maxLength": 20480,
			//	  "minLength": 1,
			//	  "pattern": "",
			//	  "type": "string"
			//	}
			Description: "The JSON policy document that is the content for the policy",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 20480),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The possible types for the access policy",
			//	  "enum": [
			//	    "data"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The possible types for the access policy",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"data",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Amazon OpenSearchServerless access policy resource",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::OpenSearchServerless::AccessPolicy").WithTerraformTypeName("awscc_opensearchserverless_access_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description": "Description",
		"name":        "Name",
		"policy":      "Policy",
		"type":        "Type",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
