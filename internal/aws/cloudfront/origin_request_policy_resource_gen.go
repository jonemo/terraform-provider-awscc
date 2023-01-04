// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_cloudfront_origin_request_policy", originRequestPolicyResource)
}

// originRequestPolicyResource returns the Terraform awscc_cloudfront_origin_request_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::CloudFront::OriginRequestPolicy resource.
func originRequestPolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"last_modified_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OriginRequestPolicyConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Comment": {
		//	      "type": "string"
		//	    },
		//	    "CookiesConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CookieBehavior": {
		//	          "pattern": "^(none|whitelist|all)$",
		//	          "type": "string"
		//	        },
		//	        "Cookies": {
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "required": [
		//	        "CookieBehavior"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "HeadersConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "HeaderBehavior": {
		//	          "pattern": "^(none|whitelist|allViewer|allViewerAndWhitelistCloudFront)$",
		//	          "type": "string"
		//	        },
		//	        "Headers": {
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "required": [
		//	        "HeaderBehavior"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Name": {
		//	      "type": "string"
		//	    },
		//	    "QueryStringsConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "QueryStringBehavior": {
		//	          "pattern": "^(none|whitelist|all)$",
		//	          "type": "string"
		//	        },
		//	        "QueryStrings": {
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "required": [
		//	        "QueryStringBehavior"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "Name",
		//	    "HeadersConfig",
		//	    "CookiesConfig",
		//	    "QueryStringsConfig"
		//	  ],
		//	  "type": "object"
		//	}
		"origin_request_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Comment
				"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CookiesConfig
				"cookies_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CookieBehavior
						"cookie_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("^(none|whitelist|all)$"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Cookies
						"cookies": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: HeadersConfig
				"headers_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: HeaderBehavior
						"header_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("^(none|whitelist|allViewer|allViewerAndWhitelistCloudFront)$"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Headers
						"headers": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: QueryStringsConfig
				"query_strings_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: QueryStringBehavior
						"query_string_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("^(none|whitelist|all)$"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: QueryStrings
						"query_strings": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::CloudFront::OriginRequestPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::OriginRequestPolicy").WithTerraformTypeName("awscc_cloudfront_origin_request_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"comment":                      "Comment",
		"cookie_behavior":              "CookieBehavior",
		"cookies":                      "Cookies",
		"cookies_config":               "CookiesConfig",
		"header_behavior":              "HeaderBehavior",
		"headers":                      "Headers",
		"headers_config":               "HeadersConfig",
		"id":                           "Id",
		"last_modified_time":           "LastModifiedTime",
		"name":                         "Name",
		"origin_request_policy_config": "OriginRequestPolicyConfig",
		"query_string_behavior":        "QueryStringBehavior",
		"query_strings":                "QueryStrings",
		"query_strings_config":         "QueryStringsConfig",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
