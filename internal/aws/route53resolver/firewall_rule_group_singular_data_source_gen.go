// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53resolver_firewall_rule_group", firewallRuleGroupDataSource)
}

// firewallRuleGroupDataSource returns the Terraform awscc_route53resolver_firewall_rule_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53Resolver::FirewallRuleGroup resource.
func firewallRuleGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn",
		//	  "maxLength": 600,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Rfc3339TimeString",
		//	  "maxLength": 40,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Rfc3339TimeString",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatorRequestId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the creator request.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"creator_request_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the creator request.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FirewallRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FirewallRules",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Firewall Rule associating the Rule Group to a Domain List",
		//	    "properties": {
		//	      "Action": {
		//	        "description": "Rule Action",
		//	        "enum": [
		//	          "ALLOW",
		//	          "BLOCK",
		//	          "ALERT"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "BlockOverrideDnsType": {
		//	        "description": "BlockOverrideDnsType",
		//	        "enum": [
		//	          "CNAME"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "BlockOverrideDomain": {
		//	        "description": "BlockOverrideDomain",
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "BlockOverrideTtl": {
		//	        "description": "BlockOverrideTtl",
		//	        "maximum": 604800,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      },
		//	      "BlockResponse": {
		//	        "description": "BlockResponse",
		//	        "enum": [
		//	          "NODATA",
		//	          "NXDOMAIN",
		//	          "OVERRIDE"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "FirewallDomainListId": {
		//	        "description": "ResourceId",
		//	        "maxLength": 64,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Priority": {
		//	        "description": "Rule Priority",
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "required": [
		//	      "FirewallDomainListId",
		//	      "Priority",
		//	      "Action"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"firewall_rules": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Action
					"action": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Rule Action",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: BlockOverrideDnsType
					"block_override_dns_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "BlockOverrideDnsType",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: BlockOverrideDomain
					"block_override_domain": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "BlockOverrideDomain",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: BlockOverrideTtl
					"block_override_ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "BlockOverrideTtl",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: BlockResponse
					"block_response": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "BlockResponse",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: FirewallDomainListId
					"firewall_domain_list_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "ResourceId",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Priority
					"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "Rule Priority",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "FirewallRules",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ResourceId",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ResourceId",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModificationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Rfc3339TimeString",
		//	  "maxLength": 40,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"modification_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Rfc3339TimeString",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FirewallRuleGroupName",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FirewallRuleGroupName",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OwnerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AccountId",
		//	  "maxLength": 32,
		//	  "minLength": 12,
		//	  "type": "string"
		//	}
		"owner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AccountId",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RuleCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Count",
		//	  "type": "integer"
		//	}
		"rule_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Count",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ShareStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
		//	  "enum": [
		//	    "NOT_SHARED",
		//	    "SHARED_WITH_ME",
		//	    "SHARED_BY_ME"
		//	  ],
		//	  "type": "string"
		//	}
		"share_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
		//	  "enum": [
		//	    "COMPLETE",
		//	    "DELETING",
		//	    "UPDATING",
		//	    "INACTIVE_OWNER_ACCOUNT_CLOSED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StatusMessage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FirewallRuleGroupStatus",
		//	  "type": "string"
		//	}
		"status_message": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FirewallRuleGroupStatus",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 255,
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53Resolver::FirewallRuleGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::FirewallRuleGroup").WithTerraformTypeName("awscc_route53resolver_firewall_rule_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                  "Action",
		"arn":                     "Arn",
		"block_override_dns_type": "BlockOverrideDnsType",
		"block_override_domain":   "BlockOverrideDomain",
		"block_override_ttl":      "BlockOverrideTtl",
		"block_response":          "BlockResponse",
		"creation_time":           "CreationTime",
		"creator_request_id":      "CreatorRequestId",
		"firewall_domain_list_id": "FirewallDomainListId",
		"firewall_rules":          "FirewallRules",
		"id":                      "Id",
		"key":                     "Key",
		"modification_time":       "ModificationTime",
		"name":                    "Name",
		"owner_id":                "OwnerId",
		"priority":                "Priority",
		"rule_count":              "RuleCount",
		"share_status":            "ShareStatus",
		"status":                  "Status",
		"status_message":          "StatusMessage",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
