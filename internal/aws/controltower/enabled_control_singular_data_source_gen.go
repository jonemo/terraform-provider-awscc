// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package controltower

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_controltower_enabled_control", enabledControlDataSource)
}

// enabledControlDataSource returns the Terraform awscc_controltower_enabled_control data source.
// This Terraform data source corresponds to the CloudFormation AWS::ControlTower::EnabledControl resource.
func enabledControlDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ControlIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn of the control.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:aws[0-9a-zA-Z_\\-:\\/]+$",
		//	  "type": "string"
		//	}
		"control_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn of the control.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn for Organizational unit to which the control needs to be applied",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:aws[0-9a-zA-Z_\\-:\\/]+$",
		//	  "type": "string"
		//	}
		"target_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn for Organizational unit to which the control needs to be applied",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ControlTower::EnabledControl",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ControlTower::EnabledControl").WithTerraformTypeName("awscc_controltower_enabled_control")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"control_identifier": "ControlIdentifier",
		"target_identifier":  "TargetIdentifier",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
