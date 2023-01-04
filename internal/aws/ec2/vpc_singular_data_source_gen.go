// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_vpc", vPCDataSource)
}

// vPCDataSource returns the Terraform awscc_ec2_vpc data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::VPC resource.
func vPCDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The primary IPv4 CIDR block for the VPC.",
		//	  "type": "string"
		//	}
		"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The primary IPv4 CIDR block for the VPC.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CidrBlockAssociations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of IPv4 CIDR block association IDs for the VPC.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"cidr_block_associations": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of IPv4 CIDR block association IDs for the VPC.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultNetworkAcl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default network ACL ID that is associated with the VPC.",
		//	  "insertionOrder": false,
		//	  "type": "string"
		//	}
		"default_network_acl": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The default network ACL ID that is associated with the VPC.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultSecurityGroup
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default security group ID that is associated with the VPC.",
		//	  "insertionOrder": false,
		//	  "type": "string"
		//	}
		"default_security_group": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The default security group ID that is associated with the VPC.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnableDnsHostnames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs.",
		//	  "type": "boolean"
		//	}
		"enable_dns_hostnames": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnableDnsSupport
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range \"plus two\" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default.",
		//	  "type": "boolean"
		//	}
		"enable_dns_support": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range \"plus two\" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceTenancy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The allowed tenancy of instances launched into the VPC.\n\n\"default\": An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.\n\n\"dedicated\": An instance launched into the VPC is a Dedicated Instance by default, unless you explicitly specify a tenancy of host during instance launch. You cannot specify a tenancy of default during instance launch.\n\nUpdating InstanceTenancy requires no replacement only if you are updating its value from \"dedicated\" to \"default\". Updating InstanceTenancy from \"default\" to \"dedicated\" requires replacement.",
		//	  "type": "string"
		//	}
		"instance_tenancy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The allowed tenancy of instances launched into the VPC.\n\n\"default\": An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.\n\n\"dedicated\": An instance launched into the VPC is a Dedicated Instance by default, unless you explicitly specify a tenancy of host during instance launch. You cannot specify a tenancy of default during instance launch.\n\nUpdating InstanceTenancy requires no replacement only if you are updating its value from \"dedicated\" to \"default\". Updating InstanceTenancy from \"default\" to \"dedicated\" requires replacement.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv4IpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR",
		//	  "type": "string"
		//	}
		"ipv_4_ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv4NetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool",
		//	  "type": "integer"
		//	}
		"ipv_4_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6CidrBlocks
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of IPv6 CIDR blocks that are associated with the VPC.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"ipv_6_cidr_blocks": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of IPv6 CIDR blocks that are associated with the VPC.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags for the VPC.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
			Description: "The tags for the VPC.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Id for the model.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Id for the model.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::VPC",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VPC").WithTerraformTypeName("awscc_ec2_vpc")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cidr_block":              "CidrBlock",
		"cidr_block_associations": "CidrBlockAssociations",
		"default_network_acl":     "DefaultNetworkAcl",
		"default_security_group":  "DefaultSecurityGroup",
		"enable_dns_hostnames":    "EnableDnsHostnames",
		"enable_dns_support":      "EnableDnsSupport",
		"instance_tenancy":        "InstanceTenancy",
		"ipv_4_ipam_pool_id":      "Ipv4IpamPoolId",
		"ipv_4_netmask_length":    "Ipv4NetmaskLength",
		"ipv_6_cidr_blocks":       "Ipv6CidrBlocks",
		"key":                     "Key",
		"tags":                    "Tags",
		"value":                   "Value",
		"vpc_id":                  "VpcId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
