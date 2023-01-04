// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lakeformation_data_cells_filters", dataCellsFiltersDataSource)
}

// dataCellsFiltersDataSource returns the Terraform awscc_lakeformation_data_cells_filters data source.
// This Terraform data source corresponds to the CloudFormation AWS::LakeFormation::DataCellsFilter resource.
func dataCellsFiltersDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description: "Uniquely identifies the data source.",
			Computed:    true,
		},
		"ids": schema.SetAttribute{
			Description: "Set of Resource Identifiers.",
			ElementType: types.StringType,
			Computed:    true,
		},
	}

	schema := schema.Schema{
		Description: "Plural Data Source schema for AWS::LakeFormation::DataCellsFilter",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LakeFormation::DataCellsFilter").WithTerraformTypeName("awscc_lakeformation_data_cells_filters")
	opts = opts.WithTerraformSchema(schema)

	v, err := generic.NewPluralDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
