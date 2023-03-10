package gluecatalogtable


type GlueCatalogTableStorageDescriptorSkewedInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table#skewed_column_names GlueCatalogTable#skewed_column_names}.
	SkewedColumnNames *[]*string `field:"optional" json:"skewedColumnNames" yaml:"skewedColumnNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table#skewed_column_value_location_maps GlueCatalogTable#skewed_column_value_location_maps}.
	SkewedColumnValueLocationMaps *map[string]*string `field:"optional" json:"skewedColumnValueLocationMaps" yaml:"skewedColumnValueLocationMaps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table#skewed_column_values GlueCatalogTable#skewed_column_values}.
	SkewedColumnValues *[]*string `field:"optional" json:"skewedColumnValues" yaml:"skewedColumnValues"`
}

