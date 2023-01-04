// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudformation_module_default_version", moduleDefaultVersionDataSource)
}

// moduleDefaultVersionDataSource returns the Terraform awscc_cloudformation_module_default_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFormation::ModuleDefaultVersion resource.
func moduleDefaultVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the module version to set as the default version.",
		//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/module/.+/[0-9]{8}$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the module version to set as the default version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModuleName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of a module existing in the registry.",
		//	  "pattern": "^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::MODULE",
		//	  "type": "string"
		//	}
		"module_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of a module existing in the registry.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of an existing version of the named module to set as the default.",
		//	  "pattern": "^[0-9]{8}$",
		//	  "type": "string"
		//	}
		"version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of an existing version of the named module to set as the default.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFormation::ModuleDefaultVersion",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::ModuleDefaultVersion").WithTerraformTypeName("awscc_cloudformation_module_default_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":         "Arn",
		"module_name": "ModuleName",
		"version_id":  "VersionId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
