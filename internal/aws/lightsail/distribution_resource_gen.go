// Code generated by generators/resource/main.go; DO NOT EDIT.

package lightsail

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_lightsail_distribution", distributionResource)
}

// distributionResource returns the Terraform awscc_lightsail_distribution resource.
// This Terraform resource corresponds to the CloudFormation AWS::Lightsail::Distribution resource.
func distributionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AbleToUpdateBundle
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the bundle that is currently applied to your distribution, specified using the distributionName parameter, can be changed to another bundle.",
		//	  "type": "boolean"
		//	}
		"able_to_update_bundle": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the bundle that is currently applied to your distribution, specified using the distributionName parameter, can be changed to another bundle.",
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BundleId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The bundle ID to use for the distribution.",
		//	  "type": "string"
		//	}
		"bundle_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The bundle ID to use for the distribution.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: CacheBehaviorSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that describes the cache behavior settings for the distribution.",
		//	  "properties": {
		//	    "AllowedHTTPMethods": {
		//	      "description": "The HTTP methods that are processed and forwarded to the distribution's origin.",
		//	      "type": "string"
		//	    },
		//	    "CachedHTTPMethods": {
		//	      "description": "The HTTP method responses that are cached by your distribution.",
		//	      "type": "string"
		//	    },
		//	    "DefaultTTL": {
		//	      "description": "The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.",
		//	      "format": "int64",
		//	      "type": "integer"
		//	    },
		//	    "ForwardedCookies": {
		//	      "additionalProperties": false,
		//	      "description": "An object that describes the cookies that are forwarded to the origin. Your content is cached based on the cookies that are forwarded.",
		//	      "properties": {
		//	        "CookiesAllowList": {
		//	          "description": "The specific cookies to forward to your distribution's origin.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "Option": {
		//	          "description": "Specifies which cookies to forward to the distribution's origin for a cache behavior: all, none, or allow-list to forward only the cookies specified in the cookiesAllowList parameter.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ForwardedHeaders": {
		//	      "additionalProperties": false,
		//	      "description": "An object that describes the headers that are forwarded to the origin. Your content is cached based on the headers that are forwarded.",
		//	      "properties": {
		//	        "HeadersAllowList": {
		//	          "description": "The specific headers to forward to your distribution's origin.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "Option": {
		//	          "description": "The headers that you want your distribution to forward to your origin and base caching on.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ForwardedQueryStrings": {
		//	      "additionalProperties": false,
		//	      "description": "An object that describes the query strings that are forwarded to the origin. Your content is cached based on the query strings that are forwarded.",
		//	      "properties": {
		//	        "Option": {
		//	          "description": "Indicates whether the distribution forwards and caches based on query strings.",
		//	          "type": "boolean"
		//	        },
		//	        "QueryStringsAllowList": {
		//	          "description": "The specific query strings that the distribution forwards to the origin.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "MaximumTTL": {
		//	      "description": "The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
		//	      "format": "int64",
		//	      "type": "integer"
		//	    },
		//	    "MinimumTTL": {
		//	      "description": "The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
		//	      "format": "int64",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"cache_behavior_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllowedHTTPMethods
				"allowed_http_methods": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The HTTP methods that are processed and forwarded to the distribution's origin.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CachedHTTPMethods
				"cached_http_methods": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The HTTP method responses that are cached by your distribution.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DefaultTTL
				"default_ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ForwardedCookies
				"forwarded_cookies": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CookiesAllowList
						"cookies_allow_list": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The specific cookies to forward to your distribution's origin.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
								setplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Option
						"option": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies which cookies to forward to the distribution's origin for a cache behavior: all, none, or allow-list to forward only the cookies specified in the cookiesAllowList parameter.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "An object that describes the cookies that are forwarded to the origin. Your content is cached based on the cookies that are forwarded.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ForwardedHeaders
				"forwarded_headers": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: HeadersAllowList
						"headers_allow_list": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The specific headers to forward to your distribution's origin.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
								setplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Option
						"option": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The headers that you want your distribution to forward to your origin and base caching on.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "An object that describes the headers that are forwarded to the origin. Your content is cached based on the headers that are forwarded.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ForwardedQueryStrings
				"forwarded_query_strings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Option
						"option": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Indicates whether the distribution forwards and caches based on query strings.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: QueryStringsAllowList
						"query_strings_allow_list": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The specific query strings that the distribution forwards to the origin.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
								setplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "An object that describes the query strings that are forwarded to the origin. Your content is cached based on the query strings that are forwarded.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MaximumTTL
				"maximum_ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MinimumTTL
				"minimum_ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object that describes the cache behavior settings for the distribution.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CacheBehaviors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of objects that describe the per-path cache behavior for the distribution.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Describes the per-path cache behavior of an Amazon Lightsail content delivery network (CDN) distribution.",
		//	    "properties": {
		//	      "Behavior": {
		//	        "description": "The cache behavior for the specified path.",
		//	        "type": "string"
		//	      },
		//	      "Path": {
		//	        "description": "The path to a directory or file to cached, or not cache. Use an asterisk symbol to specify wildcard directories (path/to/assets/*), and file types (*.html, *jpg, *js). Directories and file paths are case-sensitive.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"cache_behaviors": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Behavior
					"behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The cache behavior for the specified path.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Path
					"path": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The path to a directory or file to cached, or not cache. Use an asterisk symbol to specify wildcard directories (path/to/assets/*), and file types (*.html, *jpg, *js). Directories and file paths are case-sensitive.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of objects that describe the per-path cache behavior for the distribution.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CertificateName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The certificate attached to the Distribution.",
		//	  "type": "string"
		//	}
		"certificate_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The certificate attached to the Distribution.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultCacheBehavior
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that describes the default cache behavior for the distribution.",
		//	  "properties": {
		//	    "Behavior": {
		//	      "description": "The cache behavior of the distribution.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"default_cache_behavior": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Behavior
				"behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The cache behavior of the distribution.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object that describes the default cache behavior for the distribution.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: DistributionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"distribution_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DistributionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the distribution.",
		//	  "pattern": "\\w[\\w\\-]*\\w",
		//	  "type": "string"
		//	}
		"distribution_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the distribution.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("\\w[\\w\\-]*\\w"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpAddressType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP address type for the distribution.",
		//	  "type": "string"
		//	}
		"ip_address_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IP address type for the distribution.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IsEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the distribution is enabled.",
		//	  "type": "boolean"
		//	}
		"is_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the distribution is enabled.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Origin
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that describes the origin resource for the distribution, such as a Lightsail instance or load balancer.",
		//	  "properties": {
		//	    "Name": {
		//	      "description": "The name of the origin resource.",
		//	      "type": "string"
		//	    },
		//	    "ProtocolPolicy": {
		//	      "description": "The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.",
		//	      "type": "string"
		//	    },
		//	    "RegionName": {
		//	      "description": "The AWS Region name of the origin resource.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"origin": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the origin resource.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ProtocolPolicy
				"protocol_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RegionName
				"region_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The AWS Region name of the origin resource.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object that describes the origin resource for the distribution, such as a Lightsail instance or load balancer.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the distribution.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the distribution.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::Lightsail::Distribution",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Distribution").WithTerraformTypeName("awscc_lightsail_distribution")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"able_to_update_bundle":    "AbleToUpdateBundle",
		"allowed_http_methods":     "AllowedHTTPMethods",
		"behavior":                 "Behavior",
		"bundle_id":                "BundleId",
		"cache_behavior_settings":  "CacheBehaviorSettings",
		"cache_behaviors":          "CacheBehaviors",
		"cached_http_methods":      "CachedHTTPMethods",
		"certificate_name":         "CertificateName",
		"cookies_allow_list":       "CookiesAllowList",
		"default_cache_behavior":   "DefaultCacheBehavior",
		"default_ttl":              "DefaultTTL",
		"distribution_arn":         "DistributionArn",
		"distribution_name":        "DistributionName",
		"forwarded_cookies":        "ForwardedCookies",
		"forwarded_headers":        "ForwardedHeaders",
		"forwarded_query_strings":  "ForwardedQueryStrings",
		"headers_allow_list":       "HeadersAllowList",
		"ip_address_type":          "IpAddressType",
		"is_enabled":               "IsEnabled",
		"key":                      "Key",
		"maximum_ttl":              "MaximumTTL",
		"minimum_ttl":              "MinimumTTL",
		"name":                     "Name",
		"option":                   "Option",
		"origin":                   "Origin",
		"path":                     "Path",
		"protocol_policy":          "ProtocolPolicy",
		"query_strings_allow_list": "QueryStringsAllowList",
		"region_name":              "RegionName",
		"status":                   "Status",
		"tags":                     "Tags",
		"value":                    "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
