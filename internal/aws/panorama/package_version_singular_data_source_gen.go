// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package panorama

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_panorama_package_version", packageVersionDataSource)
}

// packageVersionDataSource returns the Terraform awscc_panorama_package_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::Panorama::PackageVersion resource.
func packageVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: IsLatestPatch
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"is_latest_patch": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MarkLatest
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"mark_latest": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: OwnerAccount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 12,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9a-z\\_]+$",
		//	  "type": "string"
		//	}
		"owner_account": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PackageArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"package_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PackageId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9\\-\\_\\/]+$",
		//	  "type": "string"
		//	}
		"package_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PackageName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9\\-\\_]+$",
		//	  "type": "string"
		//	}
		"package_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PackageVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^([0-9]+)\\.([0-9]+)$",
		//	  "type": "string"
		//	}
		"package_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PatchVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9]+$",
		//	  "type": "string"
		//	}
		"patch_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RegisteredTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"registered_time": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "REGISTER_PENDING",
		//	    "REGISTER_COMPLETED",
		//	    "FAILED",
		//	    "DELETING"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: StatusDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"status_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedLatestPatchVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9]+$",
		//	  "type": "string"
		//	}
		"updated_latest_patch_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Panorama::PackageVersion",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Panorama::PackageVersion").WithTerraformTypeName("awscc_panorama_package_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"is_latest_patch":              "IsLatestPatch",
		"mark_latest":                  "MarkLatest",
		"owner_account":                "OwnerAccount",
		"package_arn":                  "PackageArn",
		"package_id":                   "PackageId",
		"package_name":                 "PackageName",
		"package_version":              "PackageVersion",
		"patch_version":                "PatchVersion",
		"registered_time":              "RegisteredTime",
		"status":                       "Status",
		"status_description":           "StatusDescription",
		"updated_latest_patch_version": "UpdatedLatestPatchVersion",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
