// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_volume", volumeResource)
}

// volumeResource returns the Terraform awscc_ec2_volume resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::Volume resource.
func volumeResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"auto_enable_io": {
			// Property: AutoEnableIO
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "boolean"
			//	}
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"availability_zone": {
			// Property: AvailabilityZone
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Required: true,
		},
		"encrypted": {
			// Property: Encrypted
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "boolean"
			//	}
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
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
		"iops": {
			// Property: Iops
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "integer"
			//	}
			Type:     types.Int64Type,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"multi_attach_enabled": {
			// Property: MultiAttachEnabled
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "boolean"
			//	}
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"outpost_arn": {
			// Property: OutpostArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"size": {
			// Property: Size
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "integer"
			//	}
			Type:     types.Int64Type,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"snapshot_id": {
			// Property: SnapshotId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
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
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"throughput": {
			// Property: Throughput
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "integer"
			//	}
			Type:     types.Int64Type,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"volume_type": {
			// Property: VolumeType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::EC2::Volume",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::Volume").WithTerraformTypeName("awscc_ec2_volume")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_enable_io":       "AutoEnableIO",
		"availability_zone":    "AvailabilityZone",
		"encrypted":            "Encrypted",
		"id":                   "Id",
		"iops":                 "Iops",
		"key":                  "Key",
		"kms_key_id":           "KmsKeyId",
		"multi_attach_enabled": "MultiAttachEnabled",
		"outpost_arn":          "OutpostArn",
		"size":                 "Size",
		"snapshot_id":          "SnapshotId",
		"tags":                 "Tags",
		"throughput":           "Throughput",
		"value":                "Value",
		"volume_type":          "VolumeType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
