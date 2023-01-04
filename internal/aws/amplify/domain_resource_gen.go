// Code generated by generators/resource/main.go; DO NOT EDIT.

package amplify

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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
	registry.AddResourceFactory("awscc_amplify_domain", domainResource)
}

// domainResource returns the Terraform awscc_amplify_domain resource.
// This Terraform resource corresponds to the CloudFormation AWS::Amplify::Domain resource.
func domainResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AppId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 20,
		//	  "minLength": 1,
		//	  "pattern": "d[a-z0-9]+",
		//	  "type": "string"
		//	}
		"app_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 20),
				stringvalidator.RegexMatches(regexp.MustCompile("d[a-z0-9]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AutoSubDomainCreationPatterns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "maxLength": 2048,
		//	    "minLength": 1,
		//	    "pattern": "(?s).+",
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"auto_sub_domain_creation_patterns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 2048),
					stringvalidator.RegexMatches(regexp.MustCompile("(?s).+"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AutoSubDomainIAMRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "pattern": "^$|^arn:.+:iam::\\d{12}:role.+",
		//	  "type": "string"
		//	}
		"auto_sub_domain_iam_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1000),
				stringvalidator.RegexMatches(regexp.MustCompile("^$|^arn:.+:iam::\\d{12}:role.+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CertificateRecord
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "type": "string"
		//	}
		"certificate_record": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DomainName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DomainStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"domain_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnableAutoSubDomain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_auto_sub_domain": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StatusReason
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "type": "string"
		//	}
		"status_reason": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubDomainSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "BranchName": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "pattern": "(?s).+",
		//	        "type": "string"
		//	      },
		//	      "Prefix": {
		//	        "maxLength": 255,
		//	        "pattern": "(?s).*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Prefix",
		//	      "BranchName"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 255,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"sub_domain_settings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: BranchName
					"branch_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 255),
							stringvalidator.RegexMatches(regexp.MustCompile("(?s).+"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Prefix
					"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(255),
							stringvalidator.RegexMatches(regexp.MustCompile("(?s).*"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Required: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(255),
			}, /*END VALIDATORS*/
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
		Description: "The AWS::Amplify::Domain resource allows you to connect a custom domain to your app.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Amplify::Domain").WithTerraformTypeName("awscc_amplify_domain")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_id":                            "AppId",
		"arn":                               "Arn",
		"auto_sub_domain_creation_patterns": "AutoSubDomainCreationPatterns",
		"auto_sub_domain_iam_role":          "AutoSubDomainIAMRole",
		"branch_name":                       "BranchName",
		"certificate_record":                "CertificateRecord",
		"domain_name":                       "DomainName",
		"domain_status":                     "DomainStatus",
		"enable_auto_sub_domain":            "EnableAutoSubDomain",
		"prefix":                            "Prefix",
		"status_reason":                     "StatusReason",
		"sub_domain_settings":               "SubDomainSettings",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
